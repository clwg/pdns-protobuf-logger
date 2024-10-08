package dnsmessage

import (
	"fmt"
	"log"
	"net"
	"strings"

	logwriter "github.com/clwg/go-rotating-logger"
	"github.com/clwg/pdns-protobuf-logger/utils"
)

func PassiveDNS(logger *logwriter.Logger) {
	log.Printf("Passive logging enabled")
	for message := range RawMessageChannel {
		//if message.GetType() != pb.PBDNSMessage_DNSResponseType {
		//	continue
		//}

		for _, rrs := range message.Response.GetRrs() {
			passiveRecord := PassiveDNSRecord{}

			var rdata string

			switch rrs.GetType() {
			case 1, 28:
				rdata = net.IP(rrs.GetRdata()).String()
			default:
				rdata = string(rrs.GetRdata())
			}

			qname := message.Question.GetQName()
			rname := rrs.GetName()
			rtype := rrs.GetType()
			keySeparator := ":"
			key := strings.Join([]string{qname, rname, fmt.Sprint(rtype), rdata}, keySeparator)

			id, err := utils.GenerateUUIDv5(key)
			if err != nil {
				log.Printf("Error generating UUID: %v", err)
				continue
			}

			passiveRecord.Timestamp = utils.ConvertToTimestamp(message.GetTimeSec(), message.GetTimeUsec())
			passiveRecord.Id = id
			passiveRecord.Qname = qname
			passiveRecord.Rname = rname
			passiveRecord.Rtype = rtype
			passiveRecord.Rdata = rdata

			logger.Log(passiveRecord)
		}
	}
}
