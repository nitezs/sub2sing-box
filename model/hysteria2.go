package model

type Hysteria2Obfs struct {
	Type     string `json:"type,omitempty"`
	Password string `json:"password,omitempty"`
}

type Hysteria2OutboundOptions struct {
	DialerOptions
	ServerOptions
	UpMbps   int            `json:"up_mbps,omitempty"`
	DownMbps int            `json:"down_mbps,omitempty"`
	Obfs     *Hysteria2Obfs `json:"obfs,omitempty"`
	Password string         `json:"password,omitempty"`
	Network  string         `json:"network,omitempty"`
	OutboundTLSOptionsContainer
	BrutalDebug bool `json:"brutal_debug,omitempty"`
}
