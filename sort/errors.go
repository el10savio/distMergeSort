package sort

import "errors"

var (
	errEmptyURL              = errors.New("empty url provided")
	errEmptyPeer             = errors.New("empty peer provided")
	errNoPeersAvailable      = errors.New("no peers available")
	errInvalidHTTPStatusCode = errors.New("received invalid status code")
)
