package model

type ShadowTLSUser struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type ShadowTLSHandshakeOptions struct {
	ServerOptions
	DialerOptions
}

type ShadowTLSOutboundOptions struct {
	DialerOptions
	ServerOptions
	Version  int    `json:"version,omitempty"`
	Password string `json:"password,omitempty"`
	OutboundTLSOptionsContainer
}
