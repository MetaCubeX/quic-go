package http3

import (
	"context"
	"testing"

	"github.com/metacubex/http"
	"github.com/metacubex/quic-go"
	"github.com/metacubex/tls"
)

// TestTransportDialNilConn verifies that Transport.dial returns an error
// instead of panicking when the Dial hook returns a nil *quic.Conn together
// with a nil error.
//
// Such a hook violates the "nil error implies a non-nil connection" contract
// (an inner QUIC dial racing connection teardown can yield (nil, nil)).
// Without the guard the nil conn was passed to newClientConn and dereferenced
// in conn.QlogTrace(), crashing the whole process from the goroutine spawned
// by getClient.
func TestTransportDialNilConn(t *testing.T) {
	tr := &Transport{
		Dial: func(ctx context.Context, addr string, tlsCfg *tls.Config, cfg *quic.Config) (*quic.Conn, error) {
			return nil, nil // contract violation: nil conn, nil error
		},
	}
	defer tr.Close()

	req, err := http.NewRequest(http.MethodGet, "https://example.com/", nil)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}

	resp, err := tr.RoundTrip(req)
	if err == nil {
		t.Fatal("RoundTrip with a nil-returning Dial hook: got nil error, want a non-nil error")
	}
	if resp != nil {
		t.Fatalf("RoundTrip with a nil-returning Dial hook: got non-nil response %v, want nil", resp)
	}
}
