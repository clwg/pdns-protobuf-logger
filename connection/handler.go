package connection

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"time"

	"github.com/clwg/goProtobufPDNSLogger/dnsmessage"
	pb "github.com/clwg/goProtobufPDNSLogger/protos"
	"google.golang.org/protobuf/proto"
)

// HandleConnection reads incoming messages from a net.Conn and unmarshals them into a pb.PBDNSMessage.
// It then sends the unmarshalled message to the dnsmessage.RawMessageChannel.
// If there is an error reading or unmarshalling the message, the function logs the error and returns.
// The function sets a read deadline of 5 minutes on the connection.
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(5 * time.Minute))

	for {
		lenBuf := make([]byte, 2)
		_, err := io.ReadFull(conn, lenBuf)
		if err != nil {
			if err != io.EOF {
				log.Printf("failed to read length: %v", err)
			}
			return
		}

		messageLength := binary.BigEndian.Uint16(lenBuf)
		messageBuf := make([]byte, messageLength)
		_, err = io.ReadFull(conn, messageBuf)
		if err != nil {
			log.Printf("failed to read message: %v", err)
			return
		}

		message := &pb.PBDNSMessage{}
		if err := proto.Unmarshal(messageBuf, message); err != nil {
			log.Printf("failed to unmarshal: %v", err)
			return
		}

		dnsmessage.RawMessageChannel <- message // Send unmarshalled message to the channel
	}
}
