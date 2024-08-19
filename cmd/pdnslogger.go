package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/clwg/pdns-protobuf-logger/dnsmessage"

	logwriter "github.com/clwg/go-rotating-logger"
	"github.com/clwg/pdns-protobuf-logger/connection"
)

func main() {

	var passiveLogging bool
	var detailedLogging bool
	var authoritativeLogging bool
	var queryresponseLogging bool

	flag.BoolVar(&passiveLogging, "passive", false, "Enable passive logging")
	flag.BoolVar(&detailedLogging, "detailed", false, "Enable detailed logging")
	flag.BoolVar(&authoritativeLogging, "authoritative", false, "Enable authoritative logging")
	flag.BoolVar(&queryresponseLogging, "queryresponse", false, "Enable client query response logging")

	flag.Parse()

	// Listen on TCP port 44353 on all interfaces.
	listener, err := net.Listen("tcp", ":44353")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	if detailedLogging {
		log.Printf("Detailed logging enabled")

		DetailedConfig := logwriter.LoggerConfig{
			FilenamePrefix: "detailed",
			LogDir:         "./logs",
			MaxLines:       100000,
			RotationTime:   600 * time.Second,
		}

		DetailedLogger, err := logwriter.NewLogger(DetailedConfig)
		if err != nil {
			panic(err)
		}

		go dnsmessage.Detailed(DetailedLogger) // Goroutine for detailed DNS messages
	}

	if passiveLogging {

		PassiveConfig := logwriter.LoggerConfig{
			FilenamePrefix: "passive",
			LogDir:         "./logs",
			MaxLines:       100000,
			RotationTime:   600 * time.Second,
		}

		PassiveLogger, err := logwriter.NewLogger(PassiveConfig)
		if err != nil {
			panic(err)
		}

		go dnsmessage.PassiveDNS(PassiveLogger)
	}

	if authoritativeLogging {

		AuthoritativeConfig := logwriter.LoggerConfig{
			FilenamePrefix: "authoritative",
			LogDir:         "./logs",
			MaxLines:       100000,
			RotationTime:   600 * time.Second,
		}

		AuthoritativeLogger, err := logwriter.NewLogger(AuthoritativeConfig)
		if err != nil {
			panic(err)
		}

		go dnsmessage.Authoritative(AuthoritativeLogger)
	}

	if queryresponseLogging {

		QueryResponseConfig := logwriter.LoggerConfig{
			FilenamePrefix: "queryresponse",
			LogDir:         "./logs",
			MaxLines:       100000,
			RotationTime:   600 * time.Second,
		}

		QueryResponseLogger, err := logwriter.NewLogger(QueryResponseConfig)
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
