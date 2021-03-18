package main

import (
	"fmt"
	"time"

	"../config"
	"../network/bcast"
	"../network/peers"
)

func main() {

	// We make a channel for receiving updates on the id's of the peers that are
	//  alive on the network
	peerUpdateCh := make(chan peers.PeerUpdate)
	go peers.Receiver(config.PeerPort, peerUpdateCh)

	counterTx := make(chan config.Counter)
	go bcast.Transmitter(config.BcastPort, counterTx)

	// The example message. We just send one of these every second.
	go func() {
		counter := 1
		Counter := config.Counter{counter, time.Now()}
		for {
			fmt.Println("Time: ", Counter.Timestamp, "\tCounter: ", Counter.Counter)
			counterTx <- Counter
			Counter.Counter++
			Counter.Timestamp = time.Now()
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case p := <-peerUpdateCh:
			fmt.Printf("Peer update:\n")
			fmt.Printf("  Peers:    %q\n", p.Peers)
			fmt.Printf("  New:      %q\n", p.New)
			fmt.Printf("  Lost:     %q\n", p.Lost)
		}
	}
}
