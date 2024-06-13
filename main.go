// main.go
package main

import (
	"time"

	"github.com/Vincent-Omondi/blockchain/network"
)

func main() {
	trLocal := network.NewLocalTransport(network.NetAddr("LOCAL"))
	trRemote := network.NewLocalTransport(network.NetAddr("REMOTE"))

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("hello world"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOps{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
