package model

type ShadowsocksOutboundOptions struct {
	DialerOptions
	ServerOptions
	Method        string                    `json:"method"`
	Password      string                    `json:"password"`
	Plugin        string                    `json:"plugin,omitempty"`
	PluginOptions string                    `json:"plugin_opts,omitempty"`
	Network       string                    `json:"network,omitempty"`
	UDPOverTCP    *UDPOverTCPOptions        `json:"udp_over_tcp,omitempty"`
	Multiplex     *OutboundMultiplexOptions `json:"multiplex,omitempty"`
}
