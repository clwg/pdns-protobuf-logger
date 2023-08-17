package main

import (
	"log"
	"net"

	"github.com/clwg/pdns-protobuf-logger/dnsmessage"

	"github.com/clwg/pdns-protobuf-logger/connection"
)

func main() {
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	go dnsmessage.HandleRawMessages()
	go dnsmessage.HandleDnsResponse()
	go dnsmessage.HandleDnsQuery()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept: %v", err)
			continue
		}
		go connection.HandleConnection(conn)
	}
}
