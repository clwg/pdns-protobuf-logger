# goProtobufPDNSLogger

[![Go](https://github.com/clwg/goProtobufPDNSLogger/actions/workflows/go.yml/badge.svg)](https://github.com/clwg/goProtobufPDNSLogger/actions/workflows/go.yml)

`goProtobufPDNSLogger` is a powerDNS(recursor, authoritative, dnsdist) that logs DNS protobuf messages to rotating compressed jsonl log files.

- Authoritative DNS message logging
- Query response logging
- Configurable logging parameters such as filename prefix, log directory, maximum lines per log file, and log rotation time

## Usage

The main functionality of the project is encapsulated in the `cmd/pdnslogger.go` file. This file sets up the logging configurations and handles incoming connections.

```go
// Initialize the logger configurations
AuthoritativeConfig := writer.LoggerConfig{
    FilenamePrefix: "authoritative",
    LogDir:         "./logs",
    MaxLines:       100000,
    RotationTime:   600 * time.Second,
}

QueryResponseConfig := writer.LoggerConfig{
    FilenamePrefix: "queryresponse",
    LogDir:         "./logs",
    MaxLines:       100000,
    RotationTime:   600 * time.Second,
}


## Installation

```bash
git clone https://github.com/clwg/goProtobufPDNSLogger.git
cd goProtobufPDNSLogger
go build
```

## Contributing

Contributions are welcome. Please submit a pull request or create an issue to discuss the changes you want to make.

## License

This project is licensed under the GNU Affero General Public License v3.0.