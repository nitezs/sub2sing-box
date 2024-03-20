package model

type DirectOutboundOptions struct {
	DialerOptions
	OverrideAddress string `json:"override_address,omitempty"`
	OverridePort    uint16 `json:"override_port,omitempty"`
	ProxyProtocol   uint8  `json:"proxy_protocol,omitempty"`
}
