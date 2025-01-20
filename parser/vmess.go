package parser

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/nitezs/sub2sing-box/constant"
	"github.com/nitezs/sub2sing-box/model"
	"github.com/nitezs/sub2sing-box/util"
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common/json/badoption"
)

func ParseVmess(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, constant.VMessPrefix) {
		return model.Outbound{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy}
	}

	proxy = strings.TrimPrefix(proxy, constant.VMessPrefix)
	base64, err := util.DecodeBase64(proxy)
	if err != nil {
		return model.Outbound{}, &ParseError{Type: ErrInvalidStruct, Raw: proxy, Message: err.Error()}
	}

	var vmess model.VmessJson
	err = json.Unmarshal([]byte(base64), &vmess)
	if err != nil {
		return model.Outbound{}, &ParseError{Type: ErrInvalidStruct, Raw: proxy, Message: err.Error()}
	}

	var port uint16
	switch vmess.Port.(type) {
	case string:
		port, err = ParsePort(vmess.Port.(string))
		if err != nil {
			return model.Outbound{}, &ParseError{
				Type:    ErrInvalidPort,
				Message: err.Error(),
				Raw:     proxy,
			}
		}
	case float64:
		port = uint16(vmess.Port.(float64))
	}

	aid := 0
	switch vmess.Aid.(type) {
	case string:
		aid, err = strconv.Atoi(vmess.Aid.(string))
		if err != nil {
			return model.Outbound{}, &ParseError{Type: ErrInvalidStruct, Raw: proxy, Message: err.Error()}
		}
	case float64:
		aid = int(vmess.Aid.(float64))
	}

	if vmess.Scy == "" {
		vmess.Scy = "auto"
	}

	name, err := url.QueryUnescape(vmess.Ps)
	if err != nil {
		name = vmess.Ps
	}

	outboundOptions := option.VMessOutboundOptions{
		ServerOptions: option.ServerOptions{
			Server:     vmess.Add,
			ServerPort: port,
		},
		UUID:     vmess.Id,
		AlterId:  aid,
		Security: vmess.Scy,
	}

	if vmess.Tls == "tls" {
		var alpn []string
		if strings.Contains(vmess.Alpn, ",") {
			alpn = strings.Split(vmess.Alpn, ",")
		} else {
			alpn = nil
		}
		outboundOptions.OutboundTLSOptionsContainer = option.OutboundTLSOptionsContainer{
			TLS: &option.OutboundTLSOptions{
				Enabled: true,
				UTLS: &option.OutboundUTLSOptions{
					Fingerprint: vmess.Fp,
				},
				ALPN:       alpn,
				ServerName: vmess.Sni,
			},
		}
		if vmess.Fp != "" {
			outboundOptions.OutboundTLSOptionsContainer.TLS.UTLS = &option.OutboundUTLSOptions{
				Enabled:     true,
				Fingerprint: vmess.Fp,
			}
		}
	}

	if vmess.Net == "ws" {
		if vmess.Path == "" {
			vmess.Path = "/"
		}
		if vmess.Host == "" {
			vmess.Host = vmess.Add
		}
		outboundOptions.Transport = &option.V2RayTransportOptions{
			Type: "ws",
			WebsocketOptions: option.V2RayWebsocketOptions{
				Path: vmess.Path,
				Headers: badoption.HTTPHeader{
					"Host": {vmess.Host},
				},
			},
		}
	}

	if vmess.Net == "quic" {
		quic := option.V2RayQUICOptions{}
		outboundOptions.Transport = &option.V2RayTransportOptions{
			Type:        "quic",
			QUICOptions: quic,
		}
	}

	if vmess.Net == "grpc" {
		grpc := option.V2RayGRPCOptions{
			ServiceName:         vmess.Path,
			PermitWithoutStream: true,
		}
		outboundOptions.Transport = &option.V2RayTransportOptions{
			Type:        "grpc",
			GRPCOptions: grpc,
		}
	}

	if vmess.Net == "h2" {
		httpOps := option.V2RayHTTPOptions{
			Host: strings.Split(vmess.Host, ","),
			Path: vmess.Path,
		}
		outboundOptions.Transport = &option.V2RayTransportOptions{
			Type:        "http",
			HTTPOptions: httpOps,
		}
	}

	result := model.Outbound{
		Type:    "vmess",
		Tag:     name,
		Options: outboundOptions,
	}

	return result, nil
}
