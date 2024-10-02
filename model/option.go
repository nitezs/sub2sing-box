package model

import (
	"bytes"
	"encoding/json"

	"github.com/sagernet/sing-box/option"
)

type _Options struct {
	RawMessage   json.RawMessage             `json:"-"`
	Schema       string                      `json:"$schema,omitempty"`
	Log          *LogOptions                 `json:"log,omitempty"`
	DNS          *option.DNSOptions          `json:"dns,omitempty"`
	NTP          *option.NTPOptions          `json:"ntp,omitempty"`
	Inbounds     []option.Inbound            `json:"inbounds,omitempty"`
	Outbounds    []Outbound                  `json:"outbounds,omitempty"`
	Route        *option.RouteOptions        `json:"route,omitempty"`
	Experimental *option.ExperimentalOptions `json:"experimental,omitempty"`
}

type Options _Options

func (o *Options) UnmarshalJSON(content []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(content))
	decoder.DisallowUnknownFields()
	err := decoder.Decode((*_Options)(o))
	if err != nil {
		return err
	}
	o.RawMessage = content
	return nil
}

type LogOptions struct {
	Disabled     bool   `json:"disabled,omitempty"`
	Level        string `json:"level,omitempty"`
	Output       string `json:"output,omitempty"`
	Timestamp    bool   `json:"timestamp,omitempty"`
	DisableColor bool   `json:"-"`
}
