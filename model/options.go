package model

import (
	"bytes"
	"context"

	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common/json"
	"github.com/sagernet/sing/common/json/badjson"
)

type _Options struct {
	option.Options
	Endpoints []Endpoint `json:"endpoints,omitempty"`
	Inbounds  []Inbound  `json:"inbounds,omitempty"`
	Outbounds []Outbound `json:"outbounds,omitempty"`
}

type Options _Options

func (o *Options) UnmarshalJSONContext(ctx context.Context, content []byte) error {
	decoder := json.NewDecoderContext(ctx, bytes.NewReader(content))
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

type StubOptions struct{}

type Endpoint option.Endpoint

func (e *Endpoint) MarshalJSON() ([]byte, error) {
	return badjson.MarshallObjects((*option.Endpoint)(e), e.Options)
}

type Inbound option.Inbound

func (i *Inbound) MarshalJSON() ([]byte, error) {
	return badjson.MarshallObjects((*option.Inbound)(i), i.Options)
}

type Outbound option.Outbound

func (o *Outbound) MarshalJSON() ([]byte, error) {
	return badjson.MarshallObjects((*option.Outbound)(o), o.Options)
}
