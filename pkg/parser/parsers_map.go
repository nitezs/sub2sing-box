package parser

import (
	"sub2sing-box/internal/model"
)

var ParserMap map[string]func(string) (model.Proxy, error) = map[string]func(string) (model.Proxy, error){
	"ss://":        ParseShadowsocks,
	"vmess://":     ParseVmess,
	"trojan://":    ParseTrojan,
	"vless://":     ParseVless,
	"hysteria://":  ParseHysteria,
	"hy2://":       ParseHysteria2,
	"hysteria2://": ParseHysteria2,
}
