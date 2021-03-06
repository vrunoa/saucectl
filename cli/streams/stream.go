package streams

import (
	"io"

	"github.com/docker/docker/pkg/term"
)

// Streams is an interface which exposes the standard input and output streams
type Streams interface {
	In() *In
	Out() *Out
	Err() io.Writer
}

// commonStream is an input stream used by the DockerCli to read user input
type commonStream struct {
	fd         uintptr
	isTerminal bool
	state      *term.State
}

// FD returns the file descriptor number for this stream
func (s *commonStream) FD() uintptr {
	return s.fd
}

// IsTerminal returns true if this stream is connected to a terminal
func (s *commonStream) IsTerminal() bool {
	return s.isTerminal
}
