package parser

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"strings"
	"sub2sing-box/model"
	. "sub2sing-box/util"
)

func ParseVmess(proxy string) (model.Proxy, error) {
	if !strings.HasPrefix(proxy, "vmess://") {
		return model.Proxy{}, errors.New("invalid vmess url")
	}
	base64, err := DecodeBase64(strings.TrimPrefix(proxy, "vmess://"))
	if err != nil {
		return model.Proxy{}, errors.New("invalid vmess url" + err.Error())
	}
	var vmess model.VmessJson
	err = json.Unmarshal([]byte(base64), &vmess)
	if err != nil {
		return model.Proxy{}, errors.New("invalid vmess url" + err.Error())
	}
	port := 0
	switch vmess.Port.(type) {
	case string:
		port, err = strconv.Atoi(vmess.Port.(string))
		if err != nil {
			return model.Proxy{}, errors.New("invalid vmess url" + err.Error())
		}
	case float64:
		port = int(vmess.Port.(float64))
	}
	aid := 0
	switch vmess.Aid.(type) {
	case string:
		aid, err = strconv.Atoi(vmess.Aid.(string))
		if err != nil {
			return model.Proxy{}, errors.New("invalid vmess url" + err.Error())
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

	result := model.Proxy{
		Type: "vmess",
		Tag:  name,
		VMess: model.VMess{
			Server:     vmess.Add,
			ServerPort: uint16(port),
			UUID:       vmess.Id,
			AlterId:    aid,
			Security:   vmess.Scy,
		},
	}

	if vmess.Tls == "tls" {
		tls := model.OutboundTLSOptions{
			Enabled: true,
			UTLS: &model.OutboundUTLSOptions{
				Fingerprint: vmess.Fp,
			},
			ALPN: strings.Split(vmess.Alpn, ","),
		}
		result.VMess.TLS = &tls
	}

	if vmess.Net == "ws" {
		if vmess.Path == "" {
			vmess.Path = "/"
		}
		if vmess.Host == "" {
			vmess.Host = vmess.Add
		}
		ws := model.V2RayWebsocketOptions{
			Path: vmess.Path,
			Headers: map[string]string{
				"Host": vmess.Host,
			},
		}
		transport := model.V2RayTransportOptions{
			Type:             "ws",
			WebsocketOptions: ws,
		}
		result.VMess.Transport = &transport
	}

	if vmess.Net == "quic" {
		quic := model.V2RayQUICOptions{}
		transport := model.V2RayTransportOptions{
			Type:        "quic",
			QUICOptions: quic,
		}
		result.VMess.Transport = &transport
	}

	if vmess.Net == "grpc" {
		grpc := model.V2RayGRPCOptions{
			ServiceName:         vmess.Path,
			PermitWithoutStream: true,
		}
		transport := model.V2RayTransportOptions{
			Type:        "grpc",
			GRPCOptions: grpc,
		}
		result.VMess.Transport = &transport
	}

	if vmess.Net == "h2" {
		httpOps := model.V2RayHTTPOptions{
			Host: strings.Split(vmess.Host, ","),
			Path: vmess.Path,
		}
		transport := model.V2RayTransportOptions{
			Type:        "http",
			HTTPOptions: httpOps,
		}
		result.VMess.Transport = &transport
	}

	return result, nil
}
