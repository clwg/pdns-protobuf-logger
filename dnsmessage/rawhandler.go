package dnsmessage

import (
	"log"

	"github.com/clwg/pdns-protobuf-logger/utils"

	pb "github.com/clwg/pdns-protobuf-logger/protos"
)

var RawMessageChannel = make(chan *pb.PBDNSMessage, 10)

func HandleRawMessages() {
	for message := range RawMessageChannel {
		log.Printf("%v", message.String())
		log.Printf("\nType: %v %v", message.GetType(), message.Type)
		switch message.GetType() {
		case utils.TypeDNSQuery:
			QueryChannel <- message
		case utils.TypeDNSResponse:
			ResponseChannel <- message
		}
	}
}
