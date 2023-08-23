package dnsmessage

import (
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/clwg/pdns-protobuf-logger/protos"
	"github.com/clwg/pdns-protobuf-logger/utils"
	"github.com/clwg/pdns-protobuf-logger/writer"
)

var PassiveDNSChannel = make(chan *pb.PBDNSMessage, 10)

func PassiveDNS(logger *writer.Logger) {
	log.Printf("Passive logging enabled")
	for message := range RawMessageChannel {
		// Client query responses only
		if message.GetType() == pb.PBDNSMessage_DNSResponseType {
			for _, rrs := range message.Response.GetRrs() {

				passiveRecord := PassiveDNSRecord{}

				var rdata string
				if rrs.GetType() == 1 || rrs.GetType() == 28 {
					rdata = net.IP(rrs.GetRdata()).String()
				} else {
					rdata = string(rrs.GetRdata())
				}

				qname := message.Question.GetQName()
				rname := rrs.GetName()
				rtype := rrs.GetType()
				keyParts := []string{qname, rname, fmt.Sprint(rtype), rdata}
				separator := ":"
				key := strings.Join(keyParts, separator)

				id, err := utils.GenerateUUIDv5(key)
				if err != nil {
					log.Printf("Error generating UUID: %v", err)
				}

				passiveRecord.Timestamp = utils.ConvertToTimestamp(message.GetTimeSec(), message.GetTimeUsec())
				passiveRecord.Id = id
				passiveRecord.Qname = qname
				passiveRecord.Rname = rname
				passiveRecord.Rtype = rtype
				passiveRecord.Rdata = rdata

				log.Printf("%+v", passiveRecord)

				logger.Log(passiveRecord)

			}

		}
	}
}
