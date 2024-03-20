package model

type SocksOutboundOptions struct {
	DialerOptions
	ServerOptions
	Version    string             `json:"version,omitempty"`
	Username   string             `json:"username,omitempty"`
	Password   string             `json:"password,omitempty"`
	Network    Listable[string]   `json:"network,omitempty"`
	UDPOverTCP *UDPOverTCPOptions `json:"udp_over_tcp,omitempty"`
}

type HTTPOutboundOptions struct {
	DialerOptions
	ServerOptions
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	OutboundTLSOptionsContainer
	Path    string            `json:"path,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}
