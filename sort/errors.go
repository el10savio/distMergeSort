package sort

import "errors"

var (
	ErrNoPeers               = errors.New("no peers available")
	ErrEmptyPeer             = errors.New("empty peer provided")
	ErrInvalidHTTPStatusCode = errors.New("received invalid status code")
)
