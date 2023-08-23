package dnsmessage

import (
	"github.com/clwg/pdns-protobuf-logger/utils"

	pb "github.com/clwg/pdns-protobuf-logger/protos"
)

var RawMessageChannel = make(chan *pb.PBDNSMessage, 10)

func HandleRawMessages() {
	for message := range RawMessageChannel {

		switch message.GetType() {

		case utils.TypeDNSQuery:
			QueryChannel <- message

		case utils.TypeDNSResponse:
			ResponseChannel <- message

		case utils.TypeDNSIncomingResponse:
			IncomingResponseChannel <- message

		}
	}
}
