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
	// Command-line switches for enabling passive and detailed logging
	var passiveLogging bool
	var detailedLogging bool
	flag.BoolVar(&passiveLogging, "passive", true, "Enable passive logging")
	flag.BoolVar(&detailedLogging, "detailed", false, "Enable detailed logging")
	flag.Parse()

	// Listen on TCP port 6666 on all interfaces.
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	if detailedLogging {
		log.Printf("Detailed logging enabled")

		// Create and start the logger for detailed DNS messages
		DetailedConfig := writer.LoggerConfig{
			FilenamePrefix: "logs/detailed/dns",
			MaxLines:       100000,
			RotationTime:   6000 * time.Second,
		}
		DetailedLogger, err := writer.NewLogger(DetailedConfig)
		if err != nil {
			panic(err)
		}
		go dnsmessage.HandleRawMessages(DetailedLogger) // Goroutine for detailed DNS messages
	}

	if passiveLogging {
		log.Printf("Passive logging enabled")
		// Create and start the logger for passive DNS messages
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
