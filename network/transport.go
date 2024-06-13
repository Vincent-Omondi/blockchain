// network/transport.go
package network

import "net"

type NetAddr string

func (n NetAddr) Network() string {
	return "local"
}

func (n NetAddr) String() string {
	return string(n)
}

type RPC struct {
	From    net.Addr
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(net.Addr, []byte) error
	Broadcast([]byte) error
	Addr() net.Addr
}
