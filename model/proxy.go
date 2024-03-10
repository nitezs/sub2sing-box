package model

import (
	"encoding/json"
)

type Proxy struct {
	Type        string `json:"type"`
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
			Shadowsocks
		}{
			Type:        p.Type,
			Shadowsocks: p.Shadowsocks,
		})
	case "vmess":
		return json.Marshal(&struct {
			Type string `json:"type"`
			VMess
		}{
			Type:  p.Type,
			VMess: p.VMess,
		})
	case "vless":
		return json.Marshal(&struct {
			Type string `json:"type"`
			VLESS
		}{
			Type:  p.Type,
			VLESS: p.VLESS,
		})
	case "trojan":
		return json.Marshal(&struct {
			Type string `json:"type"`
			Trojan
		}{
			Type:   p.Type,
			Trojan: p.Trojan,
		})
	case "tuic":
		return json.Marshal(&struct {
			Type string `json:"type"`
			TUIC
		}{
			Type: p.Type,
			TUIC: p.TUIC,
		})
	case "hysteria":
		return json.Marshal(&struct {
			Type string `json:"type"`
			Hysteria
		}{
			Type:     p.Type,
			Hysteria: p.Hysteria,
		})
	case "hysteria2":
		return json.Marshal(&struct {
			Type string `json:"type"`
			Hysteria2
		}{
			Type:      p.Type,
			Hysteria2: p.Hysteria2,
		})
	default:
		return json.Marshal(p)
	}
}
