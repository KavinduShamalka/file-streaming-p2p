package p2p

// Peer is the interface that represent the remote node
type Peer interface {
	Close() error
}

// Transport is anything that handles the communication
// between nodes in the network. This can be of the
// form (TCP,UDP, Websockets,.....)
type Transport interface {
	ListenAndAccept() error //GRPC
	Consume() <-chan RPC
}
