package model

type TUICOutboundOptions struct {
	DialerOptions
	ServerOptions
	UUID              string `json:"uuid,omitempty"`
	Password          string `json:"password,omitempty"`
	CongestionControl string `json:"congestion_control,omitempty"`
	UDPRelayMode      string `json:"udp_relay_mode,omitempty"`
	UDPOverStream     bool   `json:"udp_over_stream,omitempty"`
	ZeroRTTHandshake  bool   `json:"zero_rtt_handshake,omitempty"`
	Heartbeat         string `json:"heartbeat,omitempty"`
	Network           string `json:"network,omitempty"`
	OutboundTLSOptionsContainer
}
