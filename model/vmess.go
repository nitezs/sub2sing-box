package model

type VmessJson struct {
	V    string      `json:"v"`
	Ps   string      `json:"ps"`
	Add  string      `json:"add"`
	Port interface{} `json:"port"`
	Id   string      `json:"id"`
	Aid  interface{} `json:"aid"`
	Scy  string      `json:"scy"`
	Net  string      `json:"net"`
	Type string      `json:"type"`
	Host string      `json:"host"`
	Path string      `json:"path"`
	Tls  string      `json:"tls"`
	Sni  string      `json:"sni"`
	Alpn string      `json:"alpn"`
	Fp   string      `json:"fp"`
}

type VMessOutboundOptions struct {
	DialerOptions
	ServerOptions
	UUID                string `json:"uuid"`
	Security            string `json:"security"`
	AlterId             int    `json:"alter_id,omitempty"`
	GlobalPadding       bool   `json:"global_padding,omitempty"`
	AuthenticatedLength bool   `json:"authenticated_length,omitempty"`
	Network             string `json:"network,omitempty"`
	OutboundTLSOptionsContainer
	PacketEncoding string                    `json:"packet_encoding,omitempty"`
	Multiplex      *OutboundMultiplexOptions `json:"multiplex,omitempty"`
	Transport      *V2RayTransportOptions    `json:"transport,omitempty"`
}
