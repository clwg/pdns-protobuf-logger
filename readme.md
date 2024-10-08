# pdns-protobuf-logger

[![Go](https://github.com/clwg/pdns-protobuf-logger/actions/workflows/go.yml/badge.svg)](https://github.com/clwg/pdns-protobuf-logger/actions/workflows/go.yml)

`pdns-protobuf-logger` is a powerDNS(recursor, authoritative, dnsdist) reciever that logs DNS protobuf messages to rotating compressed jsonl log files.


## Features

- Detailed logging of DNS queries and responses
- Passive logging focusing on response data.
- Authoritative logging that logs the iterative response server (must be set to root hints)
- Query Response logging that provides the full query and the ip address of the client.
- Automatic log rotation and compression.

## Usage

Modify the config.json to suit your needs and run the application.

{
    "log_type": "detailed",
    "log_dir": "./logs",
    "max_lines": 1000000,
    "rotation_time_seconds": 6000,
    "port": ":44353"
}

There are 4 log types that can be set in the configuration file:
 - detailed
 - passive
 - authoritative
 - query_response

```bash
go run cmd/pdnslogger.go -config config.json
```

Log files are rotated per the configuration file and are compressed using gzip and moved into the archive directory.


## Contributing

Contributions are welcome. Please submit a pull request or create an issue to discuss the changes you want to make.

## License

This project is licensed under the GNU Affero General Public License v3.0.