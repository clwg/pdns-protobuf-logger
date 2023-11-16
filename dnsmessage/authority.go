package dnsmessage

import (
	"log"
	"net"
	"strings"

	pb "github.com/clwg/goProtobufPDNSLogger/protos"
	"github.com/clwg/goProtobufPDNSLogger/utils"
	"github.com/clwg/goProtobufPDNSLogger/writer"
)

func Authoritative(logger *writer.Logger) {

	log.Printf("Authoritative logging enabled")

	for message := range RawMessageChannel {
		// Client queryresponse only
		if message.GetType() == pb.PBDNSMessage_DNSIncomingResponseType {
			for _, rrs := range message.Response.GetRrs() {

				authorityRecord := AuthoritativeDNSRecord{}

				var rdata string

				switch rrs.GetType() {
				case 1, 28:
					rdata = net.IP(rrs.GetRdata()).String()
				default:
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
