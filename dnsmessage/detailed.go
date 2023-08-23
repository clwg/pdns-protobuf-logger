package dnsmessage

import (
	"net"

	pb "github.com/clwg/pdns-protobuf-logger/protos"
	"github.com/clwg/pdns-protobuf-logger/writer"
)

var RawMessageChannel = make(chan *pb.PBDNSMessage, 10)

func HandleRawMessages(logger *writer.Logger) {
	for message := range RawMessageChannel {
		dnsMsg := DNSMessage{}
		dnsMsg.Type = message.GetType().String()
		dnsMsg.MessageId = message.GetMessageId()
		dnsMsg.ServerIdentity = string(message.GetServerIdentity())
		dnsMsg.SocketFamily = message.GetSocketFamily().String()
		dnsMsg.SocketProtocol = message.GetSocketProtocol().String()
		dnsMsg.From = net.IP(message.GetFrom()).String()
		dnsMsg.To = net.IP(message.GetTo()).String()
		dnsMsg.InBytes = message.GetInBytes()
		dnsMsg.TimeSec = message.GetTimeSec()
		dnsMsg.TimeUsec = message.GetTimeUsec()
		dnsMsg.Id = message.GetId()
		dnsMsg.OriginalRequestorSubnet = message.GetOriginalRequestorSubnet()
		dnsMsg.RequestorId = message.GetRequestorId()
		dnsMsg.InitialRequestId = message.GetInitialRequestId()
		dnsMsg.DeviceId = message.GetDeviceId()
		dnsMsg.NewlyObservedDomain = message.GetNewlyObservedDomain()
		dnsMsg.DeviceName = message.GetDeviceName()
		dnsMsg.FromPort = message.GetFromPort()
		dnsMsg.ToPort = message.GetToPort()

		dnsMsg.Question.QName = message.GetQuestion().GetQName()
		dnsMsg.Question.QType = message.GetQuestion().GetQType()
		dnsMsg.Question.QClass = message.GetQuestion().GetQClass()

		dnsMsg.Response.Rcode = message.GetResponse().GetRcode()
		dnsMsg.Response.AppliedPolicy = message.GetResponse().GetAppliedPolicy()
		dnsMsg.Response.Tags = message.GetResponse().GetTags()
		dnsMsg.Response.QueryTimeSec = message.GetResponse().GetQueryTimeSec()
		dnsMsg.Response.QueryTimeUsec = message.GetResponse().GetQueryTimeUsec()
		dnsMsg.Response.AppliedPolicyTrigger = message.GetResponse().GetAppliedPolicyTrigger()
		dnsMsg.Response.AppliedPolicyHit = message.GetResponse().GetAppliedPolicyHit()

		for _, rrs := range message.GetResponse().GetRrs() {
			dnsRR := DNSResponse_DNSRR{}
			dnsRR.Name = rrs.GetName()
			dnsRR.Type = rrs.GetType()
			dnsRR.Class = rrs.GetClass()
			dnsRR.Ttl = rrs.GetTtl()
			dnsRR.Udr = rrs.GetUdr()
			if rrs.GetType() == 1 || rrs.GetType() == 28 {
				dnsRR.Rdata = net.IP(rrs.GetRdata()).String()
			} else {
				dnsRR.Rdata = string(rrs.GetRdata())
			}
			dnsMsg.Response.Rrs = append(dnsMsg.Response.Rrs, dnsRR)

		}

		logger.Log(dnsMsg)

	}
}
