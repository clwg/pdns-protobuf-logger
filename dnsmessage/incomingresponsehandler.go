package dnsmessage

import (
	"log"
	"net"

	pb "github.com/clwg/pdns-protobuf-logger/protos"
)

type IncomingResponseDNSRecord struct {
	MsgType        string
	SocketProtocol string
	TimeSec        uint32
	QName          string
	QType          uint32
	QClass         uint32
	ServerIp       string
	ServerPort     uint32
	Rdata          string
	RType          uint32
}

var IncomingResponseChannel = make(chan *pb.PBDNSMessage, 10)

func HandleDnsIncomingResponse() {

	for message := range IncomingResponseChannel {
		//log.Printf("%v", message.String())
		if len(message.Response.GetRrs()) > 0 {

			for _, rrs := range message.Response.GetRrs() {

				var rdata string
				if rrs.GetType() == 1 || rrs.GetType() == 28 {
					ip := net.IP(rrs.GetRdata())
					rdata = ip.String()
				} else {
					rdata = string(rrs.GetRdata())
				}

				ServerIpAdderess := net.IP(message.GetTo())

				// Create a DnsResponse object with the extracted details
				incomingResponseDNS := IncomingResponseDNSRecord{
					SocketProtocol: message.GetSocketProtocol().String(),
					TimeSec:        message.GetTimeSec(),
					QName:          message.Question.GetQName(),
					QType:          message.Question.GetQType(),
					QClass:         message.Question.GetQClass(),
					ServerIp:       ServerIpAdderess.String(),
					ServerPort:     message.GetToPort(),
					Rdata:          rdata,
					RType:          rrs.GetType(),
				}

				log.Printf("Authoritative Answer: \t %+v", incomingResponseDNS)

			}
		}

	}

}
