package dnsmessage

import (
	"log"
	"net"

	pb "github.com/clwg/pdns-protobuf-logger/protos"
	"github.com/clwg/pdns-protobuf-logger/utils"
	"github.com/clwg/pdns-protobuf-logger/writer"
)

func QueryResponse(logger *writer.Logger) {

	log.Printf("Query Response logging enabled")

	for message := range RawMessageChannel {

		if message.GetType() == pb.PBDNSMessage_DNSResponseType {

			responseRecord := QueryResponseRecord{}

			responseRecord.SocketProtocol = message.GetSocketProtocol().String()
			responseRecord.Qname = message.Question.GetQName()
			responseRecord.Rcode = message.GetResponse().GetRcode()

			responseRecord.FromIp = net.IP(message.GetFrom()).String()
			responseRecord.FromPort = message.GetFromPort()
			responseRecord.ToIp = net.IP(message.GetTo()).String()
			responseRecord.ToPort = message.GetToPort()
			responseRecord.Timestamp = utils.ConvertToTimestamp(message.GetTimeSec(), message.GetTimeUsec())

			//log.Printf("%+v", responseRecord)

			logger.Log(responseRecord)

		}
	}
}
