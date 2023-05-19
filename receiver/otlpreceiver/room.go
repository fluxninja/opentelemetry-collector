package otlpreceiver // import "go.opentelemetry.io/collector/receiver/otlpreceiver"

import (
	"sync"
	"sync/atomic"
)

// Room is a helper for gracefully shutting things down while letting
// already-started operations to finish (like a graceful shutdown of a server).
//
// This is similar to sync.WaitGroup, but WaitGroup disallows concurrent Add
// and Wait while Room allows concurrent Enter and Close.
type Room struct {
	noCopy noCopy
	// count of things inside a room + (optionally) roomClosing flag
	state atomic.Int64
	// emptyChan is closed when the room empties
	emptyChan chan struct{}
	once      sync.Once // to ensure emptyChan is closed only once
}

const roomClosing int64 = 1 << 62

// TryEnter attempts to enter the room and returns if it succeeded.
//
// TryEnter will fail when room is closing or closed.
// If TryEnter succeeds, Leave needs to be called.
func (r *Room) TryEnter() bool {
	if r.state.Load() >= roomClosing {
		return false
	}

	newState := r.state.Add(1)
	if newState >= roomClosing {
		// Race with Close: We got very unlucky because room closed between
		// initial check (Load) and this Add. Undo this "illegal entering".
		r.Leave()
		return false
	}

	return true
}

// Leave leaves the room.
func (r *Room) Leave() {
	newState := r.state.Add(-1)
	if newState == roomClosing {
		// While multiple closing attempts shouldn't happen, they may occur in
		// some edge cases (see TryEnter).  Therefore, once.Do to avoid
		// double-close.
		r.once.Do(func() { close(r.emptyChan) })
	}
}

// Close wait for everybody to leave and closes the room (preventing future
// TryStarts to succeed).
//
// This can be called exactly once.
func (r *Room) Close() {
	// sanity check
	if r.emptyChan != nil {
		panic("double close")
	}
	r.emptyChan = make(chan struct{})
	newState := r.state.Add(roomClosing)
	if newState == roomClosing {
		// Nobody was there
		return
	}
	// Wait for last Leave
	<-r.emptyChan
}

// Helper for go vet to prevent copying.
// https://stackoverflow.com/a/52495303
type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}
