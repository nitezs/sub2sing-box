package parser

import (
	"sub2sing-box/model"
)

var ParserMap map[string]func(string) (model.Outbound, error) = map[string]func(string) (model.Outbound, error){
	"ss://":        ParseShadowsocks,
	"vmess://":     ParseVmess,
	"trojan://":    ParseTrojan,
	"vless://":     ParseVless,
	"hysteria://":  ParseHysteria,
	"hy2://":       ParseHysteria2,
	"hysteria2://": ParseHysteria2,
}
