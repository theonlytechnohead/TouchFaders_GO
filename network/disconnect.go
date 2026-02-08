package network

import (
	"io"
	"time"

	"github.com/theonlytechnohead/TouchFaders_GO/state"
)

func Disconnect() {
	state.Connected = false

	if state.Connection == nil {
		return
	}

	// to figure out if the connection is already closed, do a 1-byte read
	one := make([]byte, 1)
	state.Connection.SetReadDeadline(time.Now())
	if _, err := state.Connection.Read(one); err == io.EOF {
		state.Connection.Close()
		state.Connection = nil
	}
}
