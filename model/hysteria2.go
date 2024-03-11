package model

type Hysteria2Obfs struct {
	Type     string `json:"type,omitempty"`
	Password string `json:"password,omitempty"`
}

type Hysteria2 struct {
	Type        string              `json:"type"`
	Tag         string              `json:"tag,omitempty"`
	Server      string              `json:"server"`
	ServerPort  uint16              `json:"server_port"`
	UpMbps      int                 `json:"up_mbps,omitempty"`
	DownMbps    int                 `json:"down_mbps,omitempty"`
	Obfs        *Hysteria2Obfs      `json:"obfs,omitempty"`
	Password    string              `json:"password,omitempty"`
	Network     string              `json:"network,omitempty"`
	TLS         *OutboundTLSOptions `json:"tls,omitempty"`
	BrutalDebug bool                `json:"brutal_debug,omitempty"`
}
