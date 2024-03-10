package model

import (
	"encoding/json"
)

type V2RayTransportOptions struct {
	Type               string                  `json:"type"`
	HTTPOptions        V2RayHTTPOptions        `json:"-"`
	WebsocketOptions   V2RayWebsocketOptions   `json:"-"`
	QUICOptions        V2RayQUICOptions        `json:"-"`
	GRPCOptions        V2RayGRPCOptions        `json:"-"`
	HTTPUpgradeOptions V2RayHTTPUpgradeOptions `json:"-"`
}

func (o *V2RayTransportOptions) MarshalJSON() ([]byte, error) {
	switch o.Type {
	case "ws":
		return json.Marshal(&struct {
			Type string `json:"type"`
			*V2RayWebsocketOptions
		}{
			Type:                  o.Type,
			V2RayWebsocketOptions: &o.WebsocketOptions,
		})
	case "quic":
		return json.Marshal(&struct {
			Type string `json:"type"`
			*V2RayQUICOptions
		}{
			Type:             o.Type,
			V2RayQUICOptions: &o.QUICOptions,
		})
	case "grpc":
		return json.Marshal(&struct {
			Type string `json:"type"`
			*V2RayGRPCOptions
		}{
			Type:             o.Type,
			V2RayGRPCOptions: &o.GRPCOptions,
		})
	case "http":
		return json.Marshal(&struct {
			Type string `json:"type"`
			*V2RayHTTPOptions
		}{
			Type:             o.Type,
			V2RayHTTPOptions: &o.HTTPOptions,
		})
	default:
		return json.Marshal(&struct{}{})
	}
}

type V2RayHTTPOptions struct {
	Host        []string          `json:"host,omitempty"`
	Path        string            `json:"path,omitempty"`
	Method      string            `json:"method,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	IdleTimeout Duration          `json:"idle_timeout,omitempty"`
	PingTimeout Duration          `json:"ping_timeout,omitempty"`
}

type V2RayWebsocketOptions struct {
	Path                string            `json:"path,omitempty"`
	Headers             map[string]string `json:"headers,omitempty"`
	MaxEarlyData        uint32            `json:"max_early_data,omitempty"`
	EarlyDataHeaderName string            `json:"early_data_header_name,omitempty"`
}

type V2RayQUICOptions struct{}

type V2RayGRPCOptions struct {
	ServiceName         string   `json:"service_name,omitempty"`
	IdleTimeout         Duration `json:"idle_timeout,omitempty"`
	PingTimeout         Duration `json:"ping_timeout,omitempty"`
	PermitWithoutStream bool     `json:"permit_without_stream,omitempty"`
	ForceLite           bool     `json:"-"` // for test
}

type V2RayHTTPUpgradeOptions struct {
	Host    string            `json:"host,omitempty"`
	Path    string            `json:"path,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}
