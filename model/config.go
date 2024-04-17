package model

import "encoding/json"

type Listable[T any] []T

func (l *Listable[T]) UnmarshalJSON(data []byte) error {
	var arr []T
	if err := json.Unmarshal(data, &arr); err == nil {
		*l = arr
		return nil
	}
	var v T
	if err := json.Unmarshal(data, &v); err == nil {
		*l = []T{v}
		return nil
	}
	return nil
}

type Config struct {
	Log          any           `json:"log,omitempty"`
	DNS          any           `json:"dns,omitempty"`
	NTP          any           `json:"ntp,omitempty"`
	Inbounds     any           `json:"inbounds,omitempty"`
	Outbounds    []Outbound    `json:"outbounds,omitempty"`
	Route        *RouteOptions `json:"route,omitempty"`
	Experimental any           `json:"experimental,omitempty"`
}

type RouteOptions struct {
	GeoIP               *GeoIPOptions     `json:"geoip,omitempty"`
	Geosite             *GeositeOptions   `json:"geosite,omitempty"`
	Rules               Listable[Rule]    `json:"rules,omitempty"`
	RuleSet             Listable[RuleSet] `json:"rule_set,omitempty"`
	Final               string            `json:"final,omitempty"`
	FindProcess         bool              `json:"find_process,omitempty"`
	AutoDetectInterface bool              `json:"auto_detect_interface,omitempty"`
	OverrideAndroidVPN  bool              `json:"override_android_vpn,omitempty"`
	DefaultInterface    string            `json:"default_interface,omitempty"`
	DefaultMark         int               `json:"default_mark,omitempty"`
}

type Rule struct {
	Type                     string           `json:"type,omitempty"`
	Inbound                  Listable[string] `json:"inbound,omitempty"`
	IPVersion                int              `json:"ip_version,omitempty"`
	Network                  Listable[string] `json:"network,omitempty"`
	AuthUser                 Listable[string] `json:"auth_user,omitempty"`
	Protocol                 string           `json:"protocol,omitempty"`
	Domain                   Listable[string] `json:"domain,omitempty"`
	DomainSuffix             Listable[string] `json:"domain_suffix,omitempty"`
	DomainKeyword            Listable[string] `json:"domain_keyword,omitempty"`
	DomainRegex              Listable[string] `json:"domain_regex,omitempty"`
	Geosite                  Listable[string] `json:"geosite,omitempty"`
	SourceGeoIP              Listable[string] `json:"source_geoip,omitempty"`
	GeoIP                    Listable[string] `json:"geoip,omitempty"`
	SourceIPCIDR             Listable[string] `json:"source_ip_cidr,omitempty"`
	SourceIPIsPrivate        bool             `json:"source_ip_is_private,omitempty"`
	IPCIDR                   Listable[string] `json:"ip_cidr,omitempty"`
	IPIsPrivate              bool             `json:"ip_is_private,omitempty"`
	SourcePort               Listable[uint16] `json:"source_port,omitempty"`
	SourcePortRange          Listable[string] `json:"source_port_range,omitempty"`
	Port                     Listable[uint16] `json:"port,omitempty"`
	PortRange                Listable[string] `json:"port_range,omitempty"`
	ProcessName              Listable[string] `json:"process_name,omitempty"`
	ProcessPath              Listable[string] `json:"process_path,omitempty"`
	PackageName              Listable[string] `json:"package_name,omitempty"`
	User                     Listable[string] `json:"user,omitempty"`
	UserID                   Listable[int32]  `json:"user_id,omitempty"`
	ClashMode                string           `json:"clash_mode,omitempty"`
	WIFISSID                 Listable[string] `json:"wifi_ssid,omitempty"`
	WIFIBSSID                Listable[string] `json:"wifi_bssid,omitempty"`
	RuleSet                  Listable[string] `json:"rule_set,omitempty"`
	RuleSetIPCIDRMatchSource bool             `json:"rule_set_ipcidr_match_source,omitempty"`
	Invert                   bool             `json:"invert,omitempty"`
	Outbound                 string           `json:"outbound,omitempty"`
	Mode                     string           `json:"mode,omitempty"`
	Rules                    Listable[Rule]   `json:"rules,omitempty"`
}

type GeoIPOptions struct {
	Path           string `json:"path,omitempty"`
	DownloadURL    string `json:"download_url,omitempty"`
	DownloadDetour string `json:"download_detour,omitempty"`
}

type GeositeOptions struct {
	Path           string `json:"path,omitempty"`
	DownloadURL    string `json:"download_url,omitempty"`
	DownloadDetour string `json:"download_detour,omitempty"`
}

type RuleSet struct {
	Type           string `json:"type"`
	Tag            string `json:"tag"`
	Format         string `json:"format"`
	Path           string `json:"path,omitempty"`
	URL            string `json:"url"`
	DownloadDetour string `json:"download_detour,omitempty"`
	UpdateInterval string `json:"update_interval,omitempty"`
}
