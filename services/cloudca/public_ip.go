package cloudca

type PublicIp struct {
	Id string `json:"id,omitempty"`
	IpAddress string `json:"ipaddress,omitempty"`
	Purposes []string `json:"purposes,omitempty"`
	Ports []string `json:"ports,omitempty"`
}