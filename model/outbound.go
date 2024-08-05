package model

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	C "sub2sing-box/constant"
)

type _Outbound struct {
	Type                string                      `json:"type"`
	Tag                 string                      `json:"tag,omitempty"`
	DirectOptions       DirectOutboundOptions       `json:"-"`
	SocksOptions        SocksOutboundOptions        `json:"-"`
	HTTPOptions         HTTPOutboundOptions         `json:"-"`
	ShadowsocksOptions  ShadowsocksOutboundOptions  `json:"-"`
	VMessOptions        VMessOutboundOptions        `json:"-"`
	TrojanOptions       TrojanOutboundOptions       `json:"-"`
	WireGuardOptions    WireGuardOutboundOptions    `json:"-"`
	HysteriaOptions     HysteriaOutboundOptions     `json:"-"`
	TorOptions          TorOutboundOptions          `json:"-"`
	SSHOptions          SSHOutboundOptions          `json:"-"`
	ShadowTLSOptions    ShadowTLSOutboundOptions    `json:"-"`
	ShadowsocksROptions ShadowsocksROutboundOptions `json:"-"`
	VLESSOptions        VLESSOutboundOptions        `json:"-"`
	TUICOptions         TUICOutboundOptions         `json:"-"`
	Hysteria2Options    Hysteria2OutboundOptions    `json:"-"`
	SelectorOptions     SelectorOutboundOptions     `json:"-"`
	URLTestOptions      URLTestOutboundOptions      `json:"-"`
}
type Outbound _Outbound

func (h *Outbound) RawOptions() (any, error) {
	var rawOptionsPtr any
	switch h.Type {
	case C.TypeDirect:
		rawOptionsPtr = &h.DirectOptions
	case C.TypeBlock, C.TypeDNS:
		rawOptionsPtr = nil
	case C.TypeSOCKS:
		rawOptionsPtr = &h.SocksOptions
	case C.TypeHTTP:
		rawOptionsPtr = &h.HTTPOptions
	case C.TypeShadowsocks:
		rawOptionsPtr = &h.ShadowsocksOptions
	case C.TypeVMess:
		rawOptionsPtr = &h.VMessOptions
	case C.TypeTrojan:
		rawOptionsPtr = &h.TrojanOptions
	case C.TypeWireGuard:
		rawOptionsPtr = &h.WireGuardOptions
	case C.TypeHysteria:
		rawOptionsPtr = &h.HysteriaOptions
	case C.TypeTor:
		rawOptionsPtr = &h.TorOptions
	case C.TypeSSH:
		rawOptionsPtr = &h.SSHOptions
	case C.TypeShadowTLS:
		rawOptionsPtr = &h.ShadowTLSOptions
	case C.TypeShadowsocksR:
		rawOptionsPtr = &h.ShadowsocksROptions
	case C.TypeVLESS:
		rawOptionsPtr = &h.VLESSOptions
	case C.TypeTUIC:
		rawOptionsPtr = &h.TUICOptions
	case C.TypeHysteria2:
		rawOptionsPtr = &h.Hysteria2Options
	case C.TypeSelector:
		rawOptionsPtr = &h.SelectorOptions
	case C.TypeURLTest:
		rawOptionsPtr = &h.URLTestOptions
	case "":
		return nil, errors.New("missing outbound type")
	default:
		return nil, errors.New("unknown outbound type: " + h.Type)
	}
	return rawOptionsPtr, nil
}

func (h *Outbound) MarshalJSON() ([]byte, error) {
	rawOptions, err := h.RawOptions()
	if err != nil {
		return nil, err
	}
	result := make(map[string]any)
	result["type"] = h.Type
	result["tag"] = h.Tag
	optsValue := reflect.ValueOf(rawOptions).Elem()
	optsType := optsValue.Type()
	for i := 0; i < optsType.NumField(); i++ {
		field := optsValue.Field(i)
		fieldType := optsType.Field(i)
		if fieldType.Anonymous {
			embeddedFields := reflect.ValueOf(field.Interface())
			if field.Kind() == reflect.Struct {
				for j := 0; j < embeddedFields.NumField(); j++ {
					embeddedField := embeddedFields.Field(j)
					embeddedFieldType := embeddedFields.Type().Field(j)
					processField(embeddedField, embeddedFieldType, result)
				}
			}
		} else {
			processField(field, fieldType, result)
		}
	}
	return json.Marshal(result)
}

func processField(field reflect.Value, fieldType reflect.StructField, result map[string]any) {
	jsonTag := fieldType.Tag.Get("json")
	if jsonTag == "-" {
		return
	}
	tagParts := strings.Split(jsonTag, ",")
	tagName := tagParts[0]
	if len(tagParts) > 1 && tagParts[1] == "omitempty" && field.IsZero() {
		return
	}
	result[tagName] = field.Interface()
}

func (h *Outbound) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, (*_Outbound)(h))
	if err != nil {
		return err
	}
	rawOptions, err := h.RawOptions()
	if err != nil {
		return err
	}
	if rawOptions == nil {
		return nil
	}
	err = json.Unmarshal(bytes, rawOptions)
	if err != nil {
		return err
	}
	rawOptionsType := reflect.TypeOf(rawOptions).Elem()
	hValue := reflect.ValueOf(h).Elem()
	for i := 0; i < hValue.NumField(); i++ {
		fieldType := hValue.Field(i).Type()
		if fieldType == rawOptionsType {
			hValue.Field(i).Set(reflect.ValueOf(rawOptions).Elem())
			return nil
		}
	}
	return errors.New("unknown outbound type: " + h.Type)
}

func (h *Outbound) GetOutbounds() []string {
	if h.Type == C.TypeSelector {
		return h.SelectorOptions.Outbounds
	}
	if h.Type == C.TypeURLTest {
		return h.URLTestOptions.Outbounds
	}
	return nil
}

func (h *Outbound) SetOutbounds(outbounds []string) {
	if h.Type == C.TypeSelector {
		h.SelectorOptions.Outbounds = outbounds
	}
	if h.Type == C.TypeURLTest {
		h.URLTestOptions.Outbounds = outbounds
	}
}

type DialerOptions struct {
	Detour              string `json:"detour,omitempty"`
	BindInterface       string `json:"bind_interface,omitempty"`
	Inet4BindAddress    string `json:"inet4_bind_address,omitempty"`
	Inet6BindAddress    string `json:"inet6_bind_address,omitempty"`
	ProtectPath         string `json:"protect_path,omitempty"`
	RoutingMark         int    `json:"routing_mark,omitempty"`
	ReuseAddr           bool   `json:"reuse_addr,omitempty"`
	ConnectTimeout      string `json:"connect_timeout,omitempty"`
	TCPFastOpen         bool   `json:"tcp_fast_open,omitempty"`
	TCPMultiPath        bool   `json:"tcp_multi_path,omitempty"`
	UDPFragment         *bool  `json:"udp_fragment,omitempty"`
	UDPFragmentDefault  bool   `json:"-"`
	DomainStrategy      string `json:"domain_strategy,omitempty"`
	FallbackDelay       string `json:"fallback_delay,omitempty"`
	IsWireGuardListener bool   `json:"-"`
}

type ServerOptions struct {
	Server     string `json:"server"`
	ServerPort uint16 `json:"server_port"`
}
