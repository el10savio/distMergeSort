package sort

import "errors"

var (
	ErrEmptyURL              = errors.New("empty url provided")
	ErrEmptyPeer             = errors.New("empty peer provided")
	ErrNoPeersAvailable      = errors.New("no peers available")
	ErrInvalidHTTPStatusCode = errors.New("received invalid status code")
)
