package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/ethereum/go-ethereum/p2p/nat"
)

func main() {
	const targetAddr = "127.0.0.1:25565"

	mapper := nat.Any()
	if mapper == nil {
		mapper = nat.UPnP()
	}
	if mapper == nil {
		log.Println("No NAT-PMP/UPnP service found. Continuing without port mapping.")
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Pick a random port
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to listen on a random port: %v", err)
	}
	mappedPort := ln.Addr().(*net.TCPAddr).Port
	ln.Close()

	var extPort uint16
	if mapper != nil {
		extPort, err = mapper.AddMapping("TCP", mappedPort, mappedPort, "AutoPort", time.Hour)
		if err == nil {
			publicIP, _ := mapper.ExternalIP()
			if publicIP != nil {
				log.Printf("TCP port %d mapped: public %s:%d", mappedPort, publicIP.String(), extPort)
			} else {
				log.Printf("TCP port %d mapped (external port %d), but could not determine public IP", mappedPort, extPort)
			}
		} else {
			log.Printf("TCP port %d mapping failed: %v", mappedPort, err)
		}
	}

	// Start TCP listener/tunnel
	go func(p int) {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", p))
		if err != nil {
			log.Printf("TCP listen on port %d failed: %v", p, err)
			return
		}
		log.Printf("Listening TCP on port %d", p)
		for {
			clientConn, err := ln.Accept()
			if err != nil {
				log.Println("Accept error:", err)
				continue
			}
			go handleTCP(clientConn, targetAddr)
		}
	}(mappedPort)

	<-stop
	log.Println("Shutting down.")
}

// handleTCP tunnels a single TCP connection to the target server.
func handleTCP(clientConn net.Conn, targetAddr string) {
	defer func() {
		err := clientConn.Close()
		if err != nil {
			log.Printf("Error closing clientConn: %v", err)
		}
	}()
	serverConn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		log.Println("Dial to target failed:", err)
		return
	}
	defer serverConn.Close()
	// Bidirectional copy (tunnel)
	go io.Copy(serverConn, clientConn)
	io.Copy(clientConn, serverConn)
}
