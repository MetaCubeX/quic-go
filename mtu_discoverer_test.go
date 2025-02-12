package quic

import (
	"fmt"
	"math/rand/v2"
	"testing"
	"time"

	"github.com/metacubex/quic-go/internal/protocol"
	"github.com/metacubex/quic-go/internal/utils"
	"github.com/metacubex/quic-go/logging"

	"github.com/stretchr/testify/require"
)

func TestMTUDiscovererTiming(t *testing.T) {
	const rtt = 100 * time.Millisecond
	var rttStats utils.RTTStats
	rttStats.UpdateRTT(rtt, 0)
	d := newMTUDiscoverer(&rttStats, 1000, 2000, func(s protocol.ByteCount) {}, nil)

	now := time.Now()
	require.False(t, d.ShouldSendProbe(now))
	d.Start(now)
	require.False(t, d.ShouldSendProbe(now))
	require.False(t, d.ShouldSendProbe(now.Add(rtt*9/2)))
	now = now.Add(5 * rtt)
	require.True(t, d.ShouldSendProbe(now))

	// only a single outstanding probe packet is permitted
	ping, _ := d.GetPing(now)
	require.False(t, d.ShouldSendProbe(now))
	now = now.Add(5 * rtt)
	require.False(t, d.ShouldSendProbe(now))
	ping.Handler.OnLost(ping.Frame)
	require.True(t, d.ShouldSendProbe(now))
}

func TestMTUDiscovererAckAndLoss(t *testing.T) {
	var mtu protocol.ByteCount
	d := newMTUDiscoverer(&utils.RTTStats{}, 1000, 2000, func(s protocol.ByteCount) { mtu = s }, nil)
	// we use an RTT of 0 here, so we don't have to advance the timer on every step
	now := time.Now()
	ping, size := d.GetPing(now)
	require.Equal(t, protocol.ByteCount(1500), size)
	// the MTU is reduced if the frame is lost
	ping.Handler.OnLost(ping.Frame)
	require.Zero(t, mtu) // no change to the MTU yet

	require.True(t, d.ShouldSendProbe(now))
	ping, size = d.GetPing(now)
	require.Equal(t, protocol.ByteCount(1250), size)
	ping.Handler.OnAcked(ping.Frame)
	require.Equal(t, protocol.ByteCount(1250), mtu) // the MTU is increased

	// Even though the 1500 byte MTU probe packet was lost, we try again with a higher MTU.
	// This protects against regular (non-MTU-related) packet loss.
	require.True(t, d.ShouldSendProbe(now))
	ping, size = d.GetPing(now)
	require.Greater(t, size, protocol.ByteCount(1500))
	ping.Handler.OnAcked(ping.Frame)
	require.Equal(t, size, mtu)

	// We continue probing until the MTU is close to the maximum.
	var steps int
	oldSize := size
	for d.ShouldSendProbe(now) {
		ping, size = d.GetPing(now)
		require.Greater(t, size, oldSize)
		oldSize = size
		ping.Handler.OnAcked(ping.Frame)
		steps++
		require.Less(t, steps, 10)
	}
	require.Less(t, 2000-maxMTUDiff, size)
}

func TestMTUDiscovererMTUDiscovery(t *testing.T) {
	for i := range 5 {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			testMTUDiscovererMTUDiscovery(t)
		})
	}
}

func testMTUDiscovererMTUDiscovery(t *testing.T) {
	const rtt = 100 * time.Millisecond
	const startMTU protocol.ByteCount = 1000

	var rttStats utils.RTTStats
	rttStats.UpdateRTT(rtt, 0)

	maxMTU := protocol.ByteCount(rand.IntN(int(3000-startMTU))) + startMTU + 1
	currentMTU := startMTU
	var tracedMTU protocol.ByteCount
	var tracerDone bool
	d := newMTUDiscoverer(
		&rttStats,
		startMTU, maxMTU,
		func(s protocol.ByteCount) { currentMTU = s },
		&logging.ConnectionTracer{
			UpdatedMTU: func(mtu logging.ByteCount, done bool) {
				tracedMTU = mtu
				tracerDone = done
			},
		},
	)
	now := time.Now()
	d.Start(now)
	realMTU := protocol.ByteCount(rand.IntN(int(maxMTU-startMTU))) + startMTU
	t.Logf("MTU: %d, max: %d", realMTU, maxMTU)
	now = now.Add(mtuProbeDelay * rtt)
	var probes []protocol.ByteCount
	for d.ShouldSendProbe(now) {
		require.Less(t, len(probes), 25, fmt.Sprintf("too many iterations: %v", probes))
		ping, size := d.GetPing(now)
		probes = append(probes, size)
		if size <= realMTU {
			ping.Handler.OnAcked(ping.Frame)
		} else {
			ping.Handler.OnLost(ping.Frame)
		}
		now = now.Add(mtuProbeDelay * rtt)
	}
	diff := realMTU - currentMTU
	require.GreaterOrEqual(t, diff, protocol.ByteCount(0))
	if maxMTU > currentMTU+maxMTU {
		require.Equal(t, currentMTU, tracedMTU)
		require.True(t, tracerDone)
	}
	t.Logf("MTU discovered: %d (diff: %d)", currentMTU, diff)
	t.Logf("probes sent (%d): %v", len(probes), probes)
	require.LessOrEqual(t, diff, maxMTUDiff)
}

func TestMTUDiscovererWithRandomLoss(t *testing.T) {
	for i := range 5 {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			testMTUDiscovererWithRandomLoss(t)
		})
	}
}

func testMTUDiscovererWithRandomLoss(t *testing.T) {
	const rtt = 100 * time.Millisecond
	const startMTU protocol.ByteCount = 1000
	const maxRandomLoss = maxLostMTUProbes - 1

	rttStats := &utils.RTTStats{}
	rttStats.SetInitialRTT(rtt)
	require.Equal(t, rtt, rttStats.SmoothedRTT())

	maxMTU := protocol.ByteCount(rand.IntN(int(3000-startMTU))) + startMTU + 1
	currentMTU := startMTU
	var tracedMTU protocol.ByteCount
	var tracerDone bool

	d := newMTUDiscoverer(
		rttStats,
		startMTU,
		maxMTU,
		func(s protocol.ByteCount) { currentMTU = s },
		&logging.ConnectionTracer{
			UpdatedMTU: func(mtu logging.ByteCount, done bool) {
				tracedMTU = mtu
				tracerDone = done
			},
		},
	)
	d.Start(time.Now())
	now := time.Now()
	realMTU := protocol.ByteCount(rand.IntN(int(maxMTU-startMTU))) + startMTU
	t.Logf("MTU: %d, max: %d", realMTU, maxMTU)
	now = now.Add(mtuProbeDelay * rtt)
	var probes, randomLosses []protocol.ByteCount

	for d.ShouldSendProbe(now) {
		require.Less(t, len(probes), 32, fmt.Sprintf("too many iterations: %v", probes))
		ping, size := d.GetPing(now)
		probes = append(probes, size)
		packetFits := size <= realMTU
		var acked bool
		if packetFits {
			randomLoss := rand.IntN(maxLostMTUProbes) == 0 && len(randomLosses) < maxRandomLoss
			if randomLoss {
				randomLosses = append(randomLosses, size)
			} else {
				ping.Handler.OnAcked(ping.Frame)
				acked = true
			}
		}
		if !acked {
			ping.Handler.OnLost(ping.Frame)
		}
		now = now.Add(mtuProbeDelay * rtt)
	}

	diff := realMTU - currentMTU
	require.GreaterOrEqual(t, diff, protocol.ByteCount(0))
	if maxMTU > currentMTU+maxMTU {
		require.Equal(t, currentMTU, tracedMTU)
		require.True(t, tracerDone)
	}
	t.Logf("MTU discovered with random losses %v: %d (diff: %d)", randomLosses, currentMTU, diff)
	t.Logf("probes sent (%d): %v", len(probes), probes)
	require.LessOrEqual(t, diff, maxMTUDiff)
}
