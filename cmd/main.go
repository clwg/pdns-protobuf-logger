package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/clwg/pdns-protobuf-logger/dnsmessage"

	"github.com/clwg/pdns-protobuf-logger/connection"
	"github.com/clwg/pdns-protobuf-logger/writer"
)

func main() {
	// Command-line switches
	var passiveLogging bool
	var detailedLogging bool
	var authoritativeLogging bool
	var queryresponseLogging bool

	flag.BoolVar(&passiveLogging, "passive", true, "Enable passive logging")
	flag.BoolVar(&detailedLogging, "detailed", false, "Enable detailed logging")
	flag.BoolVar(&authoritativeLogging, "authoritative", false, "Enable authoritative logging")
	flag.BoolVar(&queryresponseLogging, "queryresponse", false, "Enable client query response logging")

	flag.Parse()

	// Listen on TCP port 6666 on all interfaces.
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	if detailedLogging {
		log.Printf("Detailed logging enabled")

		DetailedConfig := writer.LoggerConfig{
			FilenamePrefix: "logs/detailed/dns",
			MaxLines:       100000,
			RotationTime:   6000 * time.Second,
		}

		DetailedLogger, err := writer.NewLogger(DetailedConfig)
		if err != nil {
			panic(err)
		}

		go dnsmessage.Detailed(DetailedLogger) // Goroutine for detailed DNS messages
	}

	if passiveLogging {

		PassiveConfig := writer.LoggerConfig{
			FilenamePrefix: "logs/passive/dns",
			MaxLines:       100000,
			RotationTime:   6000 * time.Second,
		}

		PassiveLogger, err := writer.NewLogger(PassiveConfig)
		if err != nil {
			panic(err)
		}

		go dnsmessage.PassiveDNS(PassiveLogger)
	}

	if authoritativeLogging {

		AuthoritativeConfig := writer.LoggerConfig{
			FilenamePrefix: "logs/authoritative/dns",
			MaxLines:       100000,
			RotationTime:   6000 * time.Second,
		}

		AuthoritativeLogger, err := writer.NewLogger(AuthoritativeConfig)
		if err != nil {
			panic(err)
		}

		go dnsmessage.Authoritative(AuthoritativeLogger)
	}

	if queryresponseLogging {

		QueryResponseConfig := writer.LoggerConfig{
			FilenamePrefix: "logs/queryresponse/dns",
			MaxLines:       100000,
			RotationTime:   6000 * time.Second,
		}

		QueryResponseLogger, err := writer.NewLogger(QueryResponseConfig)
		if err != nil {
			panic(err)
		}

		go dnsmessage.QueryResponse(QueryResponseLogger)
	}

	// Handle incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept: %v", err)
			continue
		}
		go connection.HandleConnection(conn)
	}
}
