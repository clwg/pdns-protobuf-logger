package dnsmessage

import (
	"log"
	"net"

	pb "github.com/clwg/pdns-protobuf-logger/protos"
)

type PassiveDNSRecord struct {
	MsgType string
	Proto   string
	TimeSec int64
	QName   string
	RName   string
	RType   uint32
	RData   string
}

type DNSErrorResponse struct {
	MsgType         string
	Proto           string
	TimeSec         int64
	SourceIP        string
	SourcePort      uint32
	DestinationIP   string
	DestinationPort uint32
	QName           string
	RCode           uint32
}

var ResponseChannel = make(chan *pb.PBDNSMessage, 10)

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

				// Create a DnsResponse object with the extracted details
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
