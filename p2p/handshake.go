package p2p

// ErrInvalidHandshake is returned if the handshake between
// the local and remote node could not be established
// var ErrInvalidHandshake = errors.New("Invalid Handshake")

// HandshakeFunc is
type HandShakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
