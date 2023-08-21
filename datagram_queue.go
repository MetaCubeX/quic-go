package quic

import (
	"context"
	"sync"

	"github.com/metacubex/quic-go/internal/protocol"
	"github.com/metacubex/quic-go/internal/utils"
	"github.com/metacubex/quic-go/internal/wire"
)

const DatagramFrameMaxPeekTimes = 10

type datagramQueue struct {
	sendQueue chan *wire.DatagramFrame
	nextFrame *wire.DatagramFrame
	peekTimes int

	rcvMx    sync.Mutex
	rcvQueue [][]byte
	rcvd     chan struct{} // used to notify Receive that a new datagram was received

	closeErr error
	closed   chan struct{}

	hasData func()

	dequeued chan error

	logger utils.Logger
}

func newDatagramQueue(hasData func(), logger utils.Logger) *datagramQueue {
	return &datagramQueue{
		hasData:   hasData,
		sendQueue: make(chan *wire.DatagramFrame, 1),
		rcvd:      make(chan struct{}, 1),
		dequeued:  make(chan error),
		closed:    make(chan struct{}),
		logger:    logger,
	}
}

// AddAndWait queues a new DATAGRAM frame for sending.
// It blocks until the frame has been dequeued.
func (h *datagramQueue) AddAndWait(f *wire.DatagramFrame) error {
	select {
	case h.sendQueue <- f:
		h.hasData()
	case <-h.closed:
		return h.closeErr
	}

	select {
	case err := <-h.dequeued:
		return err
	case <-h.closed:
		return h.closeErr
	}
}

// Peek gets the next DATAGRAM frame for sending.
// If actually sent out, Pop needs to be called before the next call to Peek.
func (h *datagramQueue) Peek() *wire.DatagramFrame {
	if h.nextFrame != nil {
		h.peekTimes++
		if h.peekTimes > DatagramFrameMaxPeekTimes {
			if h.logger.Debug() {
				h.logger.Debugf("Discarded DATAGRAM frame (%d bytes payload)", len(h.nextFrame.Data))
			}
			h.nextFrame = nil
		} else {
			return h.nextFrame
		}
	}
	select {
	case h.nextFrame = <-h.sendQueue:
	default:
		return nil
	}
	h.peekTimes = 0
	return h.nextFrame
}

func (h *datagramQueue) Pop() {
	if h.nextFrame == nil {
		panic("datagramQueue BUG: Pop called for nil frame")
	}
	h.nextFrame = nil
}

// HandleDatagramFrame handles a received DATAGRAM frame.
func (h *datagramQueue) HandleDatagramFrame(f *wire.DatagramFrame) {
	data := make([]byte, len(f.Data))
	copy(data, f.Data)
	var queued bool
	h.rcvMx.Lock()
	if len(h.rcvQueue) < protocol.DatagramRcvQueueLen {
		h.rcvQueue = append(h.rcvQueue, data)
		queued = true
		select {
		case h.rcvd <- struct{}{}:
		default:
		}
	}
	h.rcvMx.Unlock()
	if !queued && h.logger.Debug() {
		h.logger.Debugf("Discarding DATAGRAM frame (%d bytes payload)", len(f.Data))
	}
}

// Receive gets a received DATAGRAM frame.
func (h *datagramQueue) Receive(ctx context.Context) ([]byte, error) {
	for {
		h.rcvMx.Lock()
		if len(h.rcvQueue) > 0 {
			data := h.rcvQueue[0]
			h.rcvQueue = h.rcvQueue[1:]
			h.rcvMx.Unlock()
			return data, nil
		}
		h.rcvMx.Unlock()
		select {
		case <-h.rcvd:
			continue
		case <-h.closed:
			return nil, h.closeErr
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func (h *datagramQueue) CloseWithError(e error) {
	h.closeErr = e
	close(h.closed)
}
