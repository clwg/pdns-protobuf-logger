package dnsmessage

import (
	"log"
	"net"

	pb "github.com/clwg/pdns-protobuf-logger/protos"
)

// PassiveDNSRecord represents a passive DNS record with the extracted details
type PassiveDNSRecord struct {
	MsgType string // message type
	Proto   string // socket protocol
	TimeSec int64  // time in seconds
	QName   string // query name
	RName   string // response name
	RType   uint32 // response type
	RData   string // response data
}

// DNSErrorResponse represents a DNS error response with the extracted details
type DNSErrorResponse struct {
	MsgType         string // message type
	Proto           string // socket protocol
	TimeSec         int64  // time in seconds
	SourceIP        string // source IP address
	SourcePort      uint32 // source port
	DestinationIP   string // destination IP address
	DestinationPort uint32 // destination port
	QName           string // query name
	RCode           uint32 // response code
}

// ResponseChannel is a channel for receiving DNS messages
var ResponseChannel = make(chan *pb.PBDNSMessage, 10)

// HandleDnsResponse handles DNS responses received on the ResponseChannel
func HandleDnsResponse() {
	for message := range ResponseChannel {
		ts := int64(message.GetTimeSec())
		qname := message.Question.GetQName()

		// if more than one item in the RRs array, iterate over them
		if len(message.Response.GetRrs()) > 0 {

			for _, rrs := range message.Response.GetRrs() {
				var rdata string
				rname := rrs.GetName()

				// Check the type and construct rdata accordingly
				if rrs.GetType() == 1 || rrs.GetType() == 28 {
					ip := net.IP(rrs.GetRdata())
					rdata = ip.String()
				} else {
					rdata = string(rrs.GetRdata())
				}

				// Create a PassiveDNSRecord object with the extracted details
				passiveDNS := PassiveDNSRecord{
					MsgType: message.GetType().String(),
					Proto:   message.GetSocketProtocol().String(),
					TimeSec: ts,
					QName:   qname,
					RName:   rname,
					RType:   rrs.GetType(),
					RData:   rdata,
				}

				log.Printf("DNS PASSIVE: \t %+v", passiveDNS)

			}

		} else {
			// Create a DNSErrorResponse object with the extracted details
			dnsError := DNSErrorResponse{
				MsgType:         message.GetType().String(),
				Proto:           message.GetSocketProtocol().String(),
				TimeSec:         ts,
				SourceIP:        net.IP(message.GetFrom()).String(),
				SourcePort:      message.GetFromPort(),
				DestinationIP:   net.IP(message.GetTo()).String(),
				DestinationPort: message.GetToPort(),
				QName:           qname,
				RCode:           message.Response.GetRcode(),
			}

			log.Printf("DNS ERROR: \t %+v", dnsError)
		}

	}
}
