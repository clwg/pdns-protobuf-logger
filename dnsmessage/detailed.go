package dnsmessage

import (
	"net"

	logwriter "github.com/clwg/go-rotating-logger"
	pb "github.com/clwg/pdns-protobuf-logger/protos"
)

var RawMessageChannel = make(chan *pb.PBDNSMessage, 10)

func Detailed(logger *logwriter.Logger) {

	for message := range RawMessageChannel {
		dnsMsg := DNSMessage{
			Type:                    message.GetType().String(),
			MessageId:               message.GetMessageId(),
			ServerIdentity:          string(message.GetServerIdentity()),
			SocketFamily:            message.GetSocketFamily().String(),
			SocketProtocol:          message.GetSocketProtocol().String(),
			From:                    net.IP(message.GetFrom()).String(),
			To:                      net.IP(message.GetTo()).String(),
			InBytes:                 message.GetInBytes(),
			TimeSec:                 message.GetTimeSec(),
			TimeUsec:                message.GetTimeUsec(),
			Id:                      message.GetId(),
			OriginalRequestorSubnet: message.GetOriginalRequestorSubnet(),
			RequestorId:             message.GetRequestorId(),
			InitialRequestId:        message.GetInitialRequestId(),
			DeviceId:                message.GetDeviceId(),
			NewlyObservedDomain:     message.GetNewlyObservedDomain(),
			DeviceName:              message.GetDeviceName(),
			FromPort:                message.GetFromPort(),
			ToPort:                  message.GetToPort(),
		}

		question := message.GetQuestion()
		dnsMsg.Question = DNSQuestion{
			QName:  question.GetQName(),
			QType:  question.GetQType(),
			QClass: question.GetQClass(),
		}

		response := message.GetResponse()
		dnsMsg.Response = DNSResponse{
			Rcode:                response.GetRcode(),
			AppliedPolicy:        response.GetAppliedPolicy(),
			Tags:                 response.GetTags(),
			QueryTimeSec:         response.GetQueryTimeSec(),
			QueryTimeUsec:        response.GetQueryTimeUsec(),
			AppliedPolicyTrigger: response.GetAppliedPolicyTrigger(),
			AppliedPolicyHit:     response.GetAppliedPolicyHit(),
			Rrs:                  make([]DNSResponse_DNSRR, 0, len(response.GetRrs())),
		}

		for _, rrs := range response.GetRrs() {
			dnsRR := DNSResponse_DNSRR{
				Name:  rrs.GetName(),
				Type:  rrs.GetType(),
				Class: rrs.GetClass(),
				Ttl:   rrs.GetTtl(),
				Udr:   rrs.GetUdr(),
			}

			switch rrs.GetType() {
			case 1, 28:
				dnsRR.Rdata = net.IP(rrs.GetRdata()).String()
			default:
				dnsRR.Rdata = string(rrs.GetRdata())
			}

			dnsMsg.Response.Rrs = append(dnsMsg.Response.Rrs, dnsRR)
		}

		logger.Log(dnsMsg)
	}
}
