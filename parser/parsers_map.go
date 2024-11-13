package parser

import (
	"github.com/nitezs/sub2sing-box/constant"
	"github.com/nitezs/sub2sing-box/model"
)

var ParserMap map[string]func(string) (model.Outbound, error) = map[string]func(string) (model.Outbound, error){
	constant.ShadowsocksPrefix: ParseShadowsocks,
	constant.VMessPrefix:       ParseVmess,
	constant.TrojanPrefix:      ParseTrojan,
	constant.VLESSPrefix:       ParseVless,
	constant.HysteriaPrefix:    ParseHysteria,
	constant.Hysteria2Prefix1:  ParseHysteria2,
	constant.Hysteria2Prefix2:  ParseHysteria2,
	constant.SocksPrefix:       ParseSocks,
}
