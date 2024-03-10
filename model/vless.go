package model

type VLESS struct {
	Type           string                    `json:"type"`
	Tag            string                    `json:"tag,omitempty"`
	Server         string                    `json:"server"`
	ServerPort     uint16                    `json:"server_port"`
	UUID           string                    `json:"uuid"`
	Flow           string                    `json:"flow,omitempty"`
	Network        string                    `json:"network,omitempty"`
	TLS            *OutboundTLSOptions       `json:"tls,omitempty"`
	Multiplex      *OutboundMultiplexOptions `json:"multiplex,omitempty"`
	Transport      *V2RayTransportOptions    `json:"transport,omitempty"`
	PacketEncoding *string                   `json:"packet_encoding,omitempty"`
}
