package main

import (
	"encoding/json"
	"flag"
	"log"
	"net"
	"os"
	"time"

	"github.com/clwg/pdns-protobuf-logger/dnsmessage"

	logwriter "github.com/clwg/go-rotating-logger"
	"github.com/clwg/pdns-protobuf-logger/connection"
)

type Config struct {
	PassiveLogging       bool   `json:"passive_logging"`
	DetailedLogging      bool   `json:"detailed_logging"`
	AuthoritativeLogging bool   `json:"authoritative_logging"`
	QueryResponseLogging bool   `json:"query_response_logging"`
	LogDir               string `json:"log_dir"`
	MaxLines             int    `json:"max_lines"`
	RotationTimeSeconds  int    `json:"rotation_time_seconds"`
	Port                 string `json:"port"`
}

func main() {
	configPath := flag.String("config", "config.json", "Path to the configuration file")
	flag.Parse()

	// Load configuration from the specified file
	config, err := loadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Listen on TCP port specified in the config
	listener, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	if config.DetailedLogging {
		log.Printf("Detailed logging enabled")

		DetailedConfig := logwriter.LoggerConfig{
			FilenamePrefix: "detailed",
			LogDir:         config.LogDir,
			MaxLines:       config.MaxLines,
			RotationTime:   time.Duration(config.RotationTimeSeconds) * time.Second,
		}

		DetailedLogger, err := logwriter.NewLogger(DetailedConfig)
		if err != nil {
			panic(err)
		}

		go dnsmessage.Detailed(DetailedLogger)
	}

	if config.PassiveLogging {

		PassiveConfig := logwriter.LoggerConfig{
			FilenamePrefix: "passive",
			LogDir:         config.LogDir,
			MaxLines:       config.MaxLines,
			RotationTime:   time.Duration(config.RotationTimeSeconds) * time.Second,
		}

		PassiveLogger, err := logwriter.NewLogger(PassiveConfig)
		if err != nil {
			panic(err)
		}

		go dnsmessage.PassiveDNS(PassiveLogger)
	}

	if config.AuthoritativeLogging {

		AuthoritativeConfig := logwriter.LoggerConfig{
			FilenamePrefix: "authoritative",
			LogDir:         config.LogDir,
			MaxLines:       config.MaxLines,
			RotationTime:   time.Duration(config.RotationTimeSeconds) * time.Second,
		}

		AuthoritativeLogger, err := logwriter.NewLogger(AuthoritativeConfig)
		if err != nil {
			panic(err)
		}

		go dnsmessage.Authoritative(AuthoritativeLogger)
	}

	if config.QueryResponseLogging {

		QueryResponseConfig := logwriter.LoggerConfig{
			FilenamePrefix: "queryresponse",
			LogDir:         config.LogDir,
			MaxLines:       config.MaxLines,
			RotationTime:   time.Duration(config.RotationTimeSeconds) * time.Second,
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
