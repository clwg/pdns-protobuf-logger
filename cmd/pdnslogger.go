package main

import (
	"encoding/json"
	"flag"
	"log"
	"net"
	"os"
	"time"

	logwriter "github.com/clwg/go-rotating-logger"
	"github.com/clwg/pdns-protobuf-logger/connection"
	"github.com/clwg/pdns-protobuf-logger/dnsmessage"
)

type Config struct {
	LogType             string `json:"log_type"`
	LogDir              string `json:"log_dir"`
	MaxLines            int    `json:"max_lines"`
	RotationTimeSeconds int    `json:"rotation_time_seconds"`
	Port                string `json:"port"`
}

func main() {
	configPath := flag.String("config", "config.json", "Path to the configuration file")
	flag.Parse()

	// Load configuration from the specified file
	config, err := loadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Validate the logging type
	validLogTypes := map[string]bool{
		"detailed":       true,
		"passive":        true,
		"authoritative":  true,
		"query_response": true,
	}
	if !validLogTypes[config.LogType] {
		log.Fatalf("invalid log_type specified in config: %s", config.LogType)
	}

	// Listen on TCP port specified in the config
	listener, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	// Configure logging based on the specified log type
	loggerConfig := logwriter.LoggerConfig{
		FilenamePrefix: config.LogType,
		LogDir:         config.LogDir,
		MaxLines:       config.MaxLines,
		RotationTime:   time.Duration(config.RotationTimeSeconds) * time.Second,
	}

	logger, err := logwriter.NewLogger(loggerConfig)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	switch config.LogType {
	case "detailed":
		log.Printf("Detailed logging enabled")
		go dnsmessage.Detailed(logger)
	case "passive":
		log.Printf("Passive logging enabled")
		go dnsmessage.PassiveDNS(logger)
	case "authoritative":
		log.Printf("Authoritative logging enabled")
		go dnsmessage.Authoritative(logger)
	case "query_response":
		log.Printf("Query response logging enabled")
		go dnsmessage.QueryResponse(logger)
	default:
		log.Fatalf("unknown logging type: %s", config.LogType)
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

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
