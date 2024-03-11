package model

type Trojan struct {
	Server     string                    `json:"server"`
	ServerPort uint16                    `json:"server_port"`
	Password   string                    `json:"password"`
	Network    string                    `json:"network,omitempty"`
	TLS        *OutboundTLSOptions       `json:"tls,omitempty"`
	Multiplex  *OutboundMultiplexOptions `json:"multiplex,omitempty"`
	Transport  *V2RayTransportOptions    `json:"transport,omitempty"`
}
