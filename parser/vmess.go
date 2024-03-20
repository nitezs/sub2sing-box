package parser

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"strings"
	model2 "sub2sing-box/model"
	"sub2sing-box/util"
)

func ParseVmess(proxy string) (model2.Proxy, error) {
	if !strings.HasPrefix(proxy, "vmess://") {
		return model2.Proxy{}, errors.New("invalid vmess url")
	}
	base64, err := util.DecodeBase64(strings.TrimPrefix(proxy, "vmess://"))
	if err != nil {
		return model2.Proxy{}, errors.New("invalid vmess url" + err.Error())
	}
	var vmess model2.VmessJson
	err = json.Unmarshal([]byte(base64), &vmess)
	if err != nil {
		return model2.Proxy{}, errors.New("invalid vmess url" + err.Error())
	}
	port := 0
	switch vmess.Port.(type) {
	case string:
		port, err = strconv.Atoi(vmess.Port.(string))
		if err != nil {
			return model2.Proxy{}, errors.New("invalid vmess url" + err.Error())
		}
	case float64:
		port = int(vmess.Port.(float64))
	}
	aid := 0
	switch vmess.Aid.(type) {
	case string:
		aid, err = strconv.Atoi(vmess.Aid.(string))
		if err != nil {
			return model2.Proxy{}, errors.New("invalid vmess url" + err.Error())
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

	result := model2.Proxy{
		Type: "vmess",
		Tag:  name,
		VMess: model2.VMess{
			Server:     vmess.Add,
			ServerPort: uint16(port),
			UUID:       vmess.Id,
			AlterId:    aid,
			Security:   vmess.Scy,
		},
	}

	if vmess.Tls == "tls" {
		var alpn []string
		if strings.Contains(vmess.Alpn, ",") {
			alpn = strings.Split(vmess.Alpn, ",")
		} else {
			alpn = nil
		}
		result.VMess.TLS = &model2.OutboundTLSOptions{
			Enabled: true,
			UTLS: &model2.OutboundUTLSOptions{
				Fingerprint: vmess.Fp,
			},
			ALPN: alpn,
		}
	}

	if vmess.Net == "ws" {
		if vmess.Path == "" {
			vmess.Path = "/"
		}
		if vmess.Host == "" {
			vmess.Host = vmess.Add
		}
		ws := model2.V2RayWebsocketOptions{
			Path: vmess.Path,
			Headers: map[string]string{
				"Host": vmess.Host,
			},
		}
		result.VMess.Transport = &model2.V2RayTransportOptions{
			Type:             "ws",
			WebsocketOptions: ws,
		}
	}

	if vmess.Net == "quic" {
		quic := model2.V2RayQUICOptions{}
		result.VMess.Transport = &model2.V2RayTransportOptions{
			Type:        "quic",
			QUICOptions: quic,
		}
		result.VMess.TLS = &model2.OutboundTLSOptions{
			Enabled: true,
		}
	}

	if vmess.Net == "grpc" {
		grpc := model2.V2RayGRPCOptions{
			ServiceName:         vmess.Path,
			PermitWithoutStream: true,
		}
		result.VMess.Transport = &model2.V2RayTransportOptions{
			Type:        "grpc",
			GRPCOptions: grpc,
		}
	}

	if vmess.Net == "h2" {
		httpOps := model2.V2RayHTTPOptions{
			Host: strings.Split(vmess.Host, ","),
			Path: vmess.Path,
		}
		result.VMess.Transport = &model2.V2RayTransportOptions{
			Type:        "http",
			HTTPOptions: httpOps,
		}
	}

	return result, nil
}
