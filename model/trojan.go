package model

type TrojanOutboundOptions struct {
	DialerOptions
	ServerOptions
	Password string `json:"password"`
	Network  string `json:"network,omitempty"`
	OutboundTLSOptionsContainer
	Multiplex *OutboundMultiplexOptions `json:"multiplex,omitempty"`
	Transport *V2RayTransportOptions    `json:"transport,omitempty"`
}
