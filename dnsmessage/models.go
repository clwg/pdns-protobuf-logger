package dnsmessage

import "time"

type DNSMessage struct {
	Question                DNSQuestion `json:"question"`
	Response                DNSResponse `json:"response"`
	InBytes                 uint64      `json:"in_bytes"`
	TimeSec                 uint32      `json:"time_sec"`
	TimeUsec                uint32      `json:"time_usec"`
	Id                      uint32      `json:"id"`
	FromPort                uint32      `json:"from_port"`
	ToPort                  uint32      `json:"to_port"`
	Type                    string      `json:"type"`
	ServerIdentity          string      `json:"server_identity"`
	SocketFamily            string      `json:"socket_family"`
	SocketProtocol          string      `json:"socket_protocol"`
	From                    string      `json:"from"`
	To                      string      `json:"to"`
	RequestorId             string      `json:"requestor_id"`
	DeviceName              string      `json:"device_name"`
	NewlyObservedDomain     bool        `json:"newly_observed_domain"`
	OriginalRequestorSubnet []byte      `json:"original_requestor_subnet"`
	InitialRequestId        []byte      `json:"initial_request_id"`
	DeviceId                []byte      `json:"device_id"`
	MessageId               []byte      `json:"message_id"`
}

type DNSQuestion struct {
	QName  string `json:"qname"`
	QType  uint32 `json:"qtype"`
	QClass uint32 `json:"qclass"`
}

type DNSResponse struct {
	AppliedPolicy        string              `json:"applied_policy"`
	AppliedPolicyTrigger string              `json:"applied_policy_trigger"`
	AppliedPolicyHit     string              `json:"applied_policy_hit"`
	ValidationState      string              `json:"validation_state"`
	Rrs                  []DNSResponse_DNSRR `json:"rrs"`
	Tags                 []string            `json:"tags"`
	Rcode                uint32              `json:"rcode"`
	QueryTimeSec         uint32              `json:"query_time_sec"`
	QueryTimeUsec        uint32              `json:"query_time_usec"`
}

type DNSResponse_DNSRR struct {
	Name  string `json:"name"`
	Rdata string `json:"rdata"`
	Type  uint32 `json:"type"`
	Class uint32 `json:"class"`
	Ttl   uint32 `json:"ttl"`
	Udr   bool   `json:"udr"`
}

type PassiveDNSRecord struct {
	Timestamp time.Time `json:"timestamp"`
	Id        string    `json:"id"`
	Qname     string    `json:"qname"`
	Rname     string    `json:"rname"`
	Rdata     string    `json:"rdata"`
	Rtype     uint32    `json:"rtype"`
}

type AuthoritativeDNSRecord struct {
	Timestamp time.Time `json:"timestamp"`
	Id        string    `json:"id"`
	Qname     string    `json:"qname"`
	ServerIp  string    `json:"server_ip"`
	Rdata     string    `json:"rdata"`
}

type QueryResponseRecord struct {
	Timestamp      time.Time `json:"timestamp"`
	SocketProtocol string    `json:"socket_protocol"`
	Qname          string    `json:"qname"`
	FromIp         string    `json:"from_ip"`
	ToIp           string    `json:"to_ip"`
	Rcode          uint32    `json:"rcode"`
	FromPort       uint32    `json:"from_port"`
	ToPort         uint32    `json:"to_port"`
}
