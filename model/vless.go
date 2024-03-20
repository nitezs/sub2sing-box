package model

type VLESSOutboundOptions struct {
	DialerOptions
	ServerOptions
	UUID    string `json:"uuid"`
	Flow    string `json:"flow,omitempty"`
	Network string `json:"network,omitempty"`
	OutboundTLSOptionsContainer
	Multiplex      *OutboundMultiplexOptions `json:"multiplex,omitempty"`
	Transport      *V2RayTransportOptions    `json:"transport,omitempty"`
	PacketEncoding *string                   `json:"packet_encoding,omitempty"`
}
