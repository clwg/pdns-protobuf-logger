package dnsmessage

import "time"

type DNSMessage struct {
	Type                    string
	MessageId               []byte
	ServerIdentity          string
	SocketFamily            string
	SocketProtocol          string
	From                    string
	To                      string
	InBytes                 uint64
	TimeSec                 uint32
	TimeUsec                uint32
	Id                      uint32
	Question                DNSQuestion
	Response                DNSResponse
	OriginalRequestorSubnet []byte
	RequestorId             string
	InitialRequestId        []byte
	DeviceId                []byte
	NewlyObservedDomain     bool
	DeviceName              string
	FromPort                uint32
	ToPort                  uint32
}

type DNSQuestion struct {
	QName  string
	QType  uint32
	QClass uint32
}

type DNSResponse struct {
	Rcode                uint32
	Rrs                  []DNSResponse_DNSRR
	AppliedPolicy        string
	Tags                 []string
	QueryTimeSec         uint32
	QueryTimeUsec        uint32
	AppliedPolicyTrigger string
	AppliedPolicyHit     string
	ValidationState      string
}

type DNSResponse_DNSRR struct {
	Name  string
	Type  uint32
	Class uint32
	Ttl   uint32
	Rdata string
	Udr   bool
}

type PassiveDNSRecord struct {
	Id        string
	Qname     string
	Rname     string
	Rtype     uint32
	Rdata     string
	Timestamp time.Time
}

type AuthoritativeDNSRecord struct {
	Id        string
	Qname     string
	ServerIp  string
	Rdata     string
	Timestamp time.Time
}

type QueryResponseRecord struct {
	SocketProtocol string
	Qname          string
	Rcode          uint32
	FromIp         string
	FromPort       uint32
	ToIp           string
	ToPort         uint32
	Timestamp      time.Time
}
