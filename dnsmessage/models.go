package dnsmessage

import "time"

type DNSMessage struct {
	Question                DNSQuestion // Struct
	Response                DNSResponse // Struct
	InBytes                 uint64      // 8 bytes
	TimeSec                 uint32      // 4 bytes
	TimeUsec                uint32      // 4 bytes
	Id                      uint32      // 4 bytes
	FromPort                uint32      // 4 bytes
	ToPort                  uint32      // 4 bytes
	Type                    string      // String (pointer)
	ServerIdentity          string      // String (pointer)
	SocketFamily            string      // String (pointer)
	SocketProtocol          string      // String (pointer)
	From                    string      // String (pointer)
	To                      string      // String (pointer)
	RequestorId             string      // String (pointer)
	DeviceName              string      // String (pointer)
	NewlyObservedDomain     bool        // 1 byte
	OriginalRequestorSubnet []byte      // Slice (pointer)
	InitialRequestId        []byte      // Slice (pointer)
	DeviceId                []byte      // Slice (pointer)
	MessageId               []byte      // Slice (pointer)

}

type DNSQuestion struct {
	QName  string // String (pointer)
	QType  uint32 // 4 bytes
	QClass uint32 // 4 bytes
}

type DNSResponse struct {
	AppliedPolicy        string              // String (pointer)
	AppliedPolicyTrigger string              // String (pointer)
	AppliedPolicyHit     string              // String (pointer)
	ValidationState      string              // String (pointer)
	Rrs                  []DNSResponse_DNSRR // Slice (pointer)
	Tags                 []string            // Slice (pointer)
	Rcode                uint32              // 4 bytes
	QueryTimeSec         uint32              // 4 bytes
	QueryTimeUsec        uint32              // 4 bytes
}

type DNSResponse_DNSRR struct {
	Name  string // String (pointer)
	Rdata string // String (pointer)
	Type  uint32 // 4 bytes
	Class uint32 // 4 bytes
	Ttl   uint32 // 4 bytes
	Udr   bool   // 1 byte
}

type PassiveDNSRecord struct {
	Timestamp time.Time // 24 bytes (on 64-bit architecture)
	Id        string    // String (pointer)
	Qname     string    // String (pointer)
	Rname     string    // String (pointer)
	Rdata     string    // String (pointer)
	Rtype     uint32    // 4 bytes
}

type AuthoritativeDNSRecord struct {
	Timestamp time.Time // 24 bytes (on 64-bit architecture)
	Id        string    // String (pointer)
	Qname     string    // String (pointer)
	ServerIp  string    // String (pointer)
	Rdata     string    // String (pointer)
}

type QueryResponseRecord struct {
	Timestamp      time.Time // 24 bytes (on 64-bit architecture)
	SocketProtocol string    // String (pointer)
	Qname          string    // String (pointer)
	FromIp         string    // String (pointer)
	ToIp           string    // String (pointer)
	Rcode          uint32    // 4 bytes
	FromPort       uint32    // 4 bytes
	ToPort         uint32    // 4 bytes
}
