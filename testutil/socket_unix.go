//go:build !windows

package testutil

import (
	"net"
	"testing"
)

// DummyUnixSocket represents a dummy Unix socket for testing socket connections.
type DummyUnixSocket struct {
	Dir      string // Directory containing the socket
	Path     string // Full path to the socket file
	listener net.Listener
}

// StartDummyUnixSocket creates a dummy Unix socket that accepts connections
// and immediately responds with test data. This is used to verify that database
// drivers correctly use socket connections.
//
// The socketName parameter is the name of the socket file to create within
// the temporary directory. For MySQL, this is typically "mysql.sock".
// For PostgreSQL, this is typically ".s.PGSQL.<port>".
//
// Returns a DummyUnixSocket which must be closed by calling Close().
func StartDummyUnixSocket(t *testing.T, dirPrefix, socketName string) *DummyUnixSocket {
	_ = "STUB: not implemented"
	return nil
}

func (s *DummyUnixSocket) acceptLoop() { _ = "STUB: not implemented"; return }

// Respond with garbage data to trigger a protocol error (not "connection refused")

// Close shuts down the socket and cleans up temporary files.
func (s *DummyUnixSocket) Close() { _ = "STUB: not implemented"; return }
