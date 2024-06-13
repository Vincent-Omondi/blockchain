// network/server.go
package network

import (
	"fmt"
	"time"
)

type ServerOps struct {
	Transports []Transport
}

type Server struct {
	ServerOps
	rpcCh  chan RPC
	quitCh chan struct{}
}

func NewServer(ops ServerOps) *Server {
	return &Server{
		ServerOps: ops,
		rpcCh:     make(chan RPC, 1024),
		quitCh:    make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(5 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%v\n", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Println("do stuff every 5 seconds")
		}
	}
	fmt.Println("Server shutdown")
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
        go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
            }
		} (tr)
    }
}