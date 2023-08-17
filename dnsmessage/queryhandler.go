package dnsmessage

import (
	"log"
	"net"

	pb "github.com/clwg/pdns-protobuf-logger/protos"
)

// DNSQuery struct will hold the details of a DNS query
type DNSQuery struct {
	SocketProtocol  string
	TimeSec         uint32
	QName           string
	QType           uint32
	QClass          uint32
	SourceIP        string
	SourcePort      uint32
	DestinationIP   string
	DestinationPort uint32
}

var QueryChannel = make(chan *pb.PBDNSMessage, 10)

func HandleDnsQuery() {
	for message := range QueryChannel {
		srcIP := net.IP(message.GetFrom())
		dstIP := net.IP(message.GetTo())

		// Construct the DNSQuery object
		query := DNSQuery{
			SocketProtocol:  message.GetSocketProtocol().String(),
			TimeSec:         message.GetTimeSec(),
			QName:           message.Question.GetQName(),
			QType:           message.Question.GetQType(),
			QClass:          message.Question.GetQClass(),
			SourceIP:        srcIP.String(),
			SourcePort:      message.GetFromPort(),
			DestinationIP:   dstIP.String(),
			DestinationPort: message.GetToPort(),
		}

		// Log the query details
		log.Printf("DNS QUERY: \t %+v", query)

	}
}
