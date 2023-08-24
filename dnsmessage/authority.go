package dnsmessage

import (
	"log"
	"net"
	"strings"

	pb "github.com/clwg/pdns-protobuf-logger/protos"
	"github.com/clwg/pdns-protobuf-logger/utils"
	"github.com/clwg/pdns-protobuf-logger/writer"
)

func Authoritative(logger *writer.Logger) {

	log.Printf("Authoritative logging enabled")

	for message := range RawMessageChannel {
		// Client queryresponse only
		if message.GetType() == pb.PBDNSMessage_DNSIncomingResponseType {
			for _, rrs := range message.Response.GetRrs() {

				authorityRecord := AuthoritativeDNSRecord{}

				var rdata string
				if rrs.GetType() == 1 || rrs.GetType() == 28 {
					rdata = net.IP(rrs.GetRdata()).String()
				} else {
					rdata = string(rrs.GetRdata())
				}

				qname := message.Question.GetQName()
				serverIp := net.IP(message.GetTo()).String()

				keyParts := []string{qname, serverIp, rdata}
				separator := ":"
				key := strings.Join(keyParts, separator)

				id, err := utils.GenerateUUIDv5(key)
				if err != nil {
					log.Printf("Error generating UUID: %v", err)
				}

				authorityRecord.Timestamp = utils.ConvertToTimestamp(message.GetTimeSec(), message.GetTimeUsec())
				authorityRecord.Id = id
				authorityRecord.Qname = qname
				authorityRecord.ServerIp = serverIp
				authorityRecord.Rdata = rdata

				//log.Printf("%+v", authorityRecord)

				logger.Log(authorityRecord)

			}
		}
	}
}
