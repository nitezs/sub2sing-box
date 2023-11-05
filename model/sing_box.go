package model

type Outbound struct {
	Type       string         `json:"type"`
	Tag        string         `json:"tag"`
	Server     string         `json:"server,omitempty"`
	ServerPort int            `json:"server_port,omitempty"`
	Password   string         `json:"password,omitempty"`
	Method     string         `json:"method,omitempty"`
	Network    string         `json:"network,omitempty"`
	Plugin     string         `json:"plugin,omitempty"`
	PluginOpts map[string]any `json:"plugin_opts,omitempty"`
	UDPOverTCP bool           `json:"udp_over_tcp,omitempty"`
	Outbounds  []string       `json:"outbounds,omitempty"`
	Default    string         `json:"default,omitempty"`
}

type SingBox struct {
	Outbounds []Outbound `json:"outbounds"`
}
