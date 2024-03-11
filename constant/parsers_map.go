package constant

import (
	"sub2sing-box/model"
	"sub2sing-box/parser"
)

var ParserMap map[string]func(string) (model.Proxy, error) = map[string]func(string) (model.Proxy, error){
	"ss://":        parser.ParseShadowsocks,
	"vmess://":     parser.ParseVmess,
	"trojan://":    parser.ParseTrojan,
	"vless://":     parser.ParseVless,
	"hysteria://":  parser.ParseHysteria,
	"hy2://":       parser.ParseHysteria2,
	"hysteria2://": parser.ParseHysteria2,
}
