package model

import (
	"encoding/json"
)

type Proxy struct {
	Type        string `json:"type"`
	Tag         string `json:"tag,omitempty"`
	Shadowsocks `json:"-"`
	VMess       `json:"-"`
	VLESS       `json:"-"`
	Trojan      `json:"-"`
	TUIC        `json:"-"`
	Hysteria    `json:"-"`
	Hysteria2   `json:"-"`
}

func (p *Proxy) MarshalJSON() ([]byte, error) {
	switch p.Type {
	case "shadowsocks":
		return json.Marshal(&struct {
			Type string `json:"type"`
			Tag  string `json:"tag,omitempty"`
			Shadowsocks
		}{
			Type:        p.Type,
			Tag:         p.Tag,
			Shadowsocks: p.Shadowsocks,
		})
	case "vmess":
		return json.Marshal(&struct {
			Type string `json:"type"`
			Tag  string `json:"tag,omitempty"`
			VMess
		}{
			Type:  p.Type,
			Tag:   p.Tag,
			VMess: p.VMess,
		})
	case "vless":
		return json.Marshal(&struct {
			Type string `json:"type"`
			Tag  string `json:"tag,omitempty"`
			VLESS
		}{
			Type:  p.Type,
			Tag:   p.Tag,
			VLESS: p.VLESS,
		})
	case "trojan":
		return json.Marshal(&struct {
			Type string `json:"type"`
			Tag  string `json:"tag,omitempty"`
			Trojan
		}{
			Type:   p.Type,
			Tag:    p.Tag,
			Trojan: p.Trojan,
		})
	case "tuic":
		return json.Marshal(&struct {
			Type string `json:"type"`
			Tag  string `json:"tag,omitempty"`
			TUIC
		}{
			Type: p.Type,
			Tag:  p.Tag,
			TUIC: p.TUIC,
		})
	case "hysteria":
		return json.Marshal(&struct {
			Type string `json:"type"`
			Tag  string `json:"tag,omitempty"`
			Hysteria
		}{
			Type:     p.Type,
			Tag:      p.Tag,
			Hysteria: p.Hysteria,
		})
	case "hysteria2":
		return json.Marshal(&struct {
			Type string `json:"type"`
			Tag  string `json:"tag,omitempty"`
			Hysteria2
		}{
			Type:      p.Type,
			Tag:       p.Tag,
			Hysteria2: p.Hysteria2,
		})
	default:
		return json.Marshal(p)
	}
}
