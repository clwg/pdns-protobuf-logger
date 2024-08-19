package connection

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"time"

	"github.com/clwg/pdns-protobuf-logger/dnsmessage"
	pb "github.com/clwg/pdns-protobuf-logger/protos"
	"google.golang.org/protobuf/proto"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	timeoutDuration := 5 * time.Minute

	for {
		// Reset the read deadline after each successful read
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))

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
