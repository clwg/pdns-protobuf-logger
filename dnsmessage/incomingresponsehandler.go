package dnsmessage

import (
	"log"
	"net"

	pb "github.com/clwg/pdns-protobuf-logger/protos"
)

type IncomingDNSRecord struct {
	MsgType string
	Proto   string
	TimeSec int64
	QName   string
	RName   string
	RType   uint32
	RData   string
}

var IncomingResponseChannel = make(chan *pb.PBDNSMessage, 10)

func HandleDnsIncomingResponse() {
	for message := range ResponseChannel {
		ts := int64(message.GetTimeSec())
		qname := message.Question.GetQName()

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

			// You can now use the dnsResponse object as needed. Here's a log statement as an example:
			//log.Printf("RR: \t %v, %v, %v, %v, %v, %v, %v", passiveDNS.MsgType, passiveDNS.Proto, passiveDNS.TimeSec, passiveDNS.QName, passiveDNS.RName, passiveDNS.RType, passiveDNS.RData)
			log.Printf("DNS INCOMING: \t %+v", passiveDNS)

		}
	}
}
