package model

type Config struct {
	Log          *LogOptions          `json:"log,omitempty"`
	DNS          *DNSOptions          `json:"dns,omitempty"`
	NTP          *NTPOptions          `json:"ntp,omitempty"`
	Inbounds     []Inbound            `json:"inbounds,omitempty"`
	Outbounds    []Outbound           `json:"outbounds,omitempty"`
	Route        *RouteOptions        `json:"route,omitempty"`
	Experimental *ExperimentalOptions `json:"experimental,omitempty"`
}

type LogOptions struct {
	Disabled  bool   `json:"disabled,omitempty"`
	Level     string `json:"level,omitempty"`
	Output    string `json:"output,omitempty"`
	Timestamp bool   `json:"timestamp,omitempty"`
}

type DNSOptions struct {
	Servers          []DNSServerOptions `json:"servers,omitempty"`
	Rules            []DNSRule          `json:"rules,omitempty"`
	Final            string             `json:"final,omitempty"`
	ReverseMapping   bool               `json:"reverse_mapping,omitempty"`
	FakeIP           *DNSFakeIPOptions  `json:"fakeip,omitempty"`
	Strategy         string             `json:"strategy,omitempty"`
	DisableCache     bool               `json:"disable_cache,omitempty"`
	DisableExpire    bool               `json:"disable_expire,omitempty"`
	IndependentCache bool               `json:"independent_cache,omitempty"`
	ClientSubnet     string             `json:"client_subnet,omitempty"`
}

type DNSServerOptions struct {
	Tag             string `json:"tag,omitempty"`
	Address         string `json:"address"`
	AddressResolver string `json:"address_resolver,omitempty"`
	AddressStrategy string `json:"address_strategy,omitempty"`
	Strategy        string `json:"strategy,omitempty"`
	Detour          string `json:"detour,omitempty"`
	ClientSubnet    string `json:"client_subnet,omitempty"`
}

type DNSRule struct {
	Type                     string    `json:"type,omitempty"`
	Inbound                  []string  `json:"inbound,omitempty"`
	IPVersion                int       `json:"ip_version,omitempty"`
	QueryType                []string  `json:"query_type,omitempty"`
	Network                  []string  `json:"network,omitempty"`
	AuthUser                 []string  `json:"auth_user,omitempty"`
	Protocol                 []string  `json:"protocol,omitempty"`
	Domain                   []string  `json:"domain,omitempty"`
	DomainSuffix             []string  `json:"domain_suffix,omitempty"`
	DomainKeyword            []string  `json:"domain_keyword,omitempty"`
	DomainRegex              []string  `json:"domain_regex,omitempty"`
	Geosite                  []string  `json:"geosite,omitempty"`
	SourceGeoIP              []string  `json:"source_geoip,omitempty"`
	GeoIP                    []string  `json:"geoip,omitempty"`
	IPCIDR                   []string  `json:"ip_cidr,omitempty"`
	IPIsPrivate              bool      `json:"ip_is_private,omitempty"`
	SourceIPCIDR             []string  `json:"source_ip_cidr,omitempty"`
	SourceIPIsPrivate        bool      `json:"source_ip_is_private,omitempty"`
	SourcePort               []uint16  `json:"source_port,omitempty"`
	SourcePortRange          []string  `json:"source_port_range,omitempty"`
	Port                     []uint16  `json:"port,omitempty"`
	PortRange                []string  `json:"port_range,omitempty"`
	ProcessName              []string  `json:"process_name,omitempty"`
	ProcessPath              []string  `json:"process_path,omitempty"`
	PackageName              []string  `json:"package_name,omitempty"`
	User                     []string  `json:"user,omitempty"`
	UserID                   []int32   `json:"user_id,omitempty"`
	Outbound                 []string  `json:"outbound,omitempty"`
	ClashMode                string    `json:"clash_mode,omitempty"`
	WIFISSID                 []string  `json:"wifi_ssid,omitempty"`
	WIFIBSSID                []string  `json:"wifi_bssid,omitempty"`
	RuleSet                  []string  `json:"rule_set,omitempty"`
	RuleSetIPCIDRMatchSource bool      `json:"rule_set_ipcidr_match_source,omitempty"`
	Invert                   bool      `json:"invert,omitempty"`
	Server                   string    `json:"server,omitempty"`
	DisableCache             bool      `json:"disable_cache,omitempty"`
	RewriteTTL               *uint32   `json:"rewrite_ttl,omitempty"`
	ClientSubnet             string    `json:"client_subnet,omitempty"`
	Mode                     string    `json:"mode"`
	Rules                    []DNSRule `json:"rules,omitempty"`
}

type DNSFakeIPOptions struct {
	Enabled    bool   `json:"enabled,omitempty"`
	Inet4Range string `json:"inet4_range,omitempty"`
	Inet6Range string `json:"inet6_range,omitempty"`
}

type NTPOptions struct {
	Enabled          bool   `json:"enabled,omitempty"`
	Server           string `json:"server,omitempty"`
	ServerPort       uint16 `json:"server_port,omitempty"`
	Interval         string `json:"interval,omitempty"`
	WriteToSystem    bool   `json:"write_to_system,omitempty"`
	Detour           string `json:"detour,omitempty"`
	BindInterface    string `json:"bind_interface,omitempty"`
	Inet4BindAddress string `json:"inet4_bind_address,omitempty"`
	Inet6BindAddress string `json:"inet6_bind_address,omitempty"`
	ProtectPath      string `json:"protect_path,omitempty"`
	RoutingMark      int    `json:"routing_mark,omitempty"`
	ReuseAddr        bool   `json:"reuse_addr,omitempty"`
	ConnectTimeout   string `json:"connect_timeout,omitempty"`
	TCPFastOpen      bool   `json:"tcp_fast_open,omitempty"`
	TCPMultiPath     bool   `json:"tcp_multi_path,omitempty"`
	UDPFragment      bool   `json:"udp_fragment,omitempty"`
	DomainStrategy   string `json:"domain_strategy,omitempty"`
	FallbackDelay    string `json:"fallback_delay,omitempty"`
}

type Inbound struct {
	Type                      string              `json:"type"`
	Tag                       string              `json:"tag,omitempty"`
	InterfaceName             string              `json:"interface_name,omitempty"`
	MTU                       uint32              `json:"mtu,omitempty"`
	GSO                       bool                `json:"gso,omitempty"`
	Inet4Address              []string            `json:"inet4_address,omitempty"`
	Inet6Address              []string            `json:"inet6_address,omitempty"`
	AutoRoute                 bool                `json:"auto_route,omitempty"`
	StrictRoute               bool                `json:"strict_route,omitempty"`
	Inet4RouteAddress         []string            `json:"inet4_route_address,omitempty"`
	Inet6RouteAddress         []string            `json:"inet6_route_address,omitempty"`
	Inet4RouteExcludeAddress  []string            `json:"inet4_route_exclude_address,omitempty"`
	Inet6RouteExcludeAddress  []string            `json:"inet6_route_exclude_address,omitempty"`
	IncludeInterface          []string            `json:"include_interface,omitempty"`
	ExcludeInterface          []string            `json:"exclude_interface,omitempty"`
	IncludeUID                []uint32            `json:"include_uid,omitempty"`
	IncludeUIDRange           []string            `json:"include_uid_range,omitempty"`
	ExcludeUID                []uint32            `json:"exclude_uid,omitempty"`
	ExcludeUIDRange           []string            `json:"exclude_uid_range,omitempty"`
	IncludeAndroidUser        []int               `json:"include_android_user,omitempty"`
	IncludePackage            []string            `json:"include_package,omitempty"`
	ExcludePackage            []string            `json:"exclude_package,omitempty"`
	EndpointIndependentNat    bool                `json:"endpoint_independent_nat,omitempty"`
	UDPTimeout                string              `json:"udp_timeout,omitempty"`
	Stack                     string              `json:"stack,omitempty"`
	Platform                  *TunPlatformOptions `json:"platform,omitempty"`
	SniffEnabled              bool                `json:"sniff,omitempty"`
	SniffOverrideDestination  bool                `json:"sniff_override_destination,omitempty"`
	SniffTimeout              string              `json:"sniff_timeout,omitempty"`
	DomainStrategy            string              `json:"domain_strategy,omitempty"`
	UDPDisableDomainUnmapping bool                `json:"udp_disable_domain_unmapping,omitempty"`
}

type TunPlatformOptions struct {
	HTTPProxy *HTTPProxyOptions `json:"http_proxy,omitempty"`
}

type HTTPProxyOptions struct {
	Enabled      bool     `json:"enabled,omitempty"`
	Server       string   `json:"server"`
	ServerPort   uint16   `json:"server_port"`
	BypassDomain []string `json:"bypass_domain,omitempty"`
	MatchDomain  []string `json:"match_domain,omitempty"`
}

type Outbound struct {
	Type                string                    `json:"type"`
	Tag                 string                    `json:"tag,omitempty"`
	Detour              string                    `json:"detour,omitempty"`
	BindInterface       string                    `json:"bind_interface,omitempty"`
	Inet4BindAddress    string                    `json:"inet4_bind_address,omitempty"`
	Inet6BindAddress    string                    `json:"inet6_bind_address,omitempty"`
	ProtectPath         string                    `json:"protect_path,omitempty"`
	RoutingMark         int                       `json:"routing_mark,omitempty"`
	ReuseAddr           bool                      `json:"reuse_addr,omitempty"`
	ConnectTimeout      string                    `json:"connect_timeout,omitempty"`
	TCPFastOpen         bool                      `json:"tcp_fast_open,omitempty"`
	TCPMultiPath        bool                      `json:"tcp_multi_path,omitempty"`
	UDPFragment         *bool                     `json:"udp_fragment,omitempty"`
	DomainStrategy      string                    `json:"domain_strategy,omitempty"`
	FallbackDelay       string                    `json:"fallback_delay,omitempty"`
	OverrideAddress     string                    `json:"override_address,omitempty"`
	OverridePort        uint16                    `json:"override_port,omitempty"`
	ProxyProtocol       uint8                     `json:"proxy_protocol,omitempty"`
	Server              string                    `json:"server"`
	ServerPort          uint16                    `json:"server_port"`
	Version             string                    `json:"version,omitempty"`
	Username            string                    `json:"username,omitempty"`
	Password            string                    `json:"password,omitempty"`
	Network             string                    `json:"network,omitempty"`
	UDPOverTCP          *UDPOverTCPOptions        `json:"udp_over_tcp,omitempty"`
	TLS                 *OutboundTLSOptions       `json:"tls,omitempty"`
	Path                string                    `json:"path,omitempty"`
	Headers             map[string][]string       `json:"headers,omitempty"`
	Method              string                    `json:"method"`
	Plugin              string                    `json:"plugin,omitempty"`
	PluginOptions       string                    `json:"plugin_opts,omitempty"`
	Multiplex           *OutboundMultiplexOptions `json:"multiplex,omitempty"`
	UUID                string                    `json:"uuid"`
	Security            string                    `json:"security"`
	AlterId             int                       `json:"alter_id,omitempty"`
	GlobalPadding       bool                      `json:"global_padding,omitempty"`
	AuthenticatedLength bool                      `json:"authenticated_length,omitempty"`
	PacketEncoding      string                    `json:"packet_encoding,omitempty"`
	Transport           *V2RayTransportOptions    `json:"transport,omitempty"`
	SystemInterface     bool                      `json:"system_interface,omitempty"`
	GSO                 bool                      `json:"gso,omitempty"`
	InterfaceName       string                    `json:"interface_name,omitempty"`
	LocalAddress        []string                  `json:"local_address"`
	PrivateKey          string                    `json:"private_key"`
	Peers               []WireGuardPeer           `json:"peers,omitempty"`
	PeerPublicKey       string                    `json:"peer_public_key"`
	PreSharedKey        string                    `json:"pre_shared_key,omitempty"`
	Reserved            []uint8                   `json:"reserved,omitempty"`
	Workers             int                       `json:"workers,omitempty"`
	MTU                 uint32                    `json:"mtu,omitempty"`
	Up                  string                    `json:"up,omitempty"`
	UpMbps              int                       `json:"up_mbps,omitempty"`
	Down                string                    `json:"down,omitempty"`
	DownMbps            int                       `json:"down_mbps,omitempty"`
	Obfs                struct {
		string
		*Hysteria2Obfs
	} `json:"obfs,omitempty"`
	Auth                      []byte            `json:"auth,omitempty"`
	AuthString                string            `json:"auth_str,omitempty"`
	ReceiveWindowConn         uint64            `json:"recv_window_conn,omitempty"`
	ReceiveWindow             uint64            `json:"recv_window,omitempty"`
	DisableMTUDiscovery       bool              `json:"disable_mtu_discovery,omitempty"`
	ExecutablePath            string            `json:"executable_path,omitempty"`
	ExtraArgs                 []string          `json:"extra_args,omitempty"`
	DataDirectory             string            `json:"data_directory,omitempty"`
	Options                   map[string]string `json:"torrc,omitempty"`
	User                      string            `json:"user,omitempty"`
	PrivateKeyPath            string            `json:"private_key_path,omitempty"`
	PrivateKeyPassphrase      string            `json:"private_key_passphrase,omitempty"`
	HostKey                   []string          `json:"host_key,omitempty"`
	HostKeyAlgorithms         []string          `json:"host_key_algorithms,omitempty"`
	ClientVersion             string            `json:"client_version,omitempty"`
	ObfsParam                 string            `json:"obfs_param,omitempty"`
	Protocol                  string            `json:"protocol,omitempty"`
	ProtocolParam             string            `json:"protocol_param,omitempty"`
	Flow                      string            `json:"flow,omitempty"`
	CongestionControl         string            `json:"congestion_control,omitempty"`
	UDPRelayMode              string            `json:"udp_relay_mode,omitempty"`
	UDPOverStream             bool              `json:"udp_over_stream,omitempty"`
	ZeroRTTHandshake          bool              `json:"zero_rtt_handshake,omitempty"`
	Heartbeat                 string            `json:"heartbeat,omitempty"`
	BrutalDebug               bool              `json:"brutal_debug,omitempty"`
	Default                   string            `json:"default,omitempty"`
	Outbounds                 []string          `json:"outbounds"`
	URL                       string            `json:"url,omitempty"`
	Interval                  string            `json:"interval,omitempty"`
	Tolerance                 uint16            `json:"tolerance,omitempty"`
	IdleTimeout               string            `json:"idle_timeout,omitempty"`
	InterruptExistConnections bool              `json:"interrupt_exist_connections,omitempty"`
}

type WireGuardPeer struct {
	Server       string   `json:"server"`
	ServerPort   uint16   `json:"server_port"`
	PublicKey    string   `json:"public_key,omitempty"`
	PreSharedKey string   `json:"pre_shared_key,omitempty"`
	AllowedIPs   []string `json:"allowed_ips,omitempty"`
	Reserved     []uint8  `json:"reserved,omitempty"`
}

type RouteOptions struct {
	GeoIP               *GeoIPOptions   `json:"geoip,omitempty"`
	Geosite             *GeositeOptions `json:"geosite,omitempty"`
	Rules               []Rule          `json:"rules,omitempty"`
	RuleSet             []RuleSet       `json:"rule_set,omitempty"`
	Final               string          `json:"final,omitempty"`
	FindProcess         bool            `json:"find_process,omitempty"`
	AutoDetectInterface bool            `json:"auto_detect_interface,omitempty"`
	OverrideAndroidVPN  bool            `json:"override_android_vpn,omitempty"`
	DefaultInterface    string          `json:"default_interface,omitempty"`
	DefaultMark         int             `json:"default_mark,omitempty"`
}

type Rule struct {
	Type                     string   `json:"type,omitempty"`
	Inbound                  []string `json:"inbound,omitempty"`
	IPVersion                int      `json:"ip_version,omitempty"`
	Network                  []string `json:"network,omitempty"`
	AuthUser                 []string `json:"auth_user,omitempty"`
	Protocol                 []string `json:"protocol,omitempty"`
	Domain                   []string `json:"domain,omitempty"`
	DomainSuffix             []string `json:"domain_suffix,omitempty"`
	DomainKeyword            []string `json:"domain_keyword,omitempty"`
	DomainRegex              []string `json:"domain_regex,omitempty"`
	Geosite                  []string `json:"geosite,omitempty"`
	SourceGeoIP              []string `json:"source_geoip,omitempty"`
	GeoIP                    []string `json:"geoip,omitempty"`
	SourceIPCIDR             []string `json:"source_ip_cidr,omitempty"`
	SourceIPIsPrivate        bool     `json:"source_ip_is_private,omitempty"`
	IPCIDR                   []string `json:"ip_cidr,omitempty"`
	IPIsPrivate              bool     `json:"ip_is_private,omitempty"`
	SourcePort               []uint16 `json:"source_port,omitempty"`
	SourcePortRange          []string `json:"source_port_range,omitempty"`
	Port                     []uint16 `json:"port,omitempty"`
	PortRange                []string `json:"port_range,omitempty"`
	ProcessName              []string `json:"process_name,omitempty"`
	ProcessPath              []string `json:"process_path,omitempty"`
	PackageName              []string `json:"package_name,omitempty"`
	User                     []string `json:"user,omitempty"`
	UserID                   []int32  `json:"user_id,omitempty"`
	ClashMode                string   `json:"clash_mode,omitempty"`
	WIFISSID                 []string `json:"wifi_ssid,omitempty"`
	WIFIBSSID                []string `json:"wifi_bssid,omitempty"`
	RuleSet                  []string `json:"rule_set,omitempty"`
	RuleSetIPCIDRMatchSource bool     `json:"rule_set_ipcidr_match_source,omitempty"`
	Invert                   bool     `json:"invert,omitempty"`
	Outbound                 string   `json:"outbound,omitempty"`
	Mode                     string   `json:"mode"`
	Rules                    []Rule   `json:"rules,omitempty"`
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

type ExperimentalOptions struct {
	CacheFile *CacheFileOptions `json:"cache_file,omitempty"`
	ClashAPI  *ClashAPIOptions  `json:"clash_api,omitempty"`
	V2RayAPI  *V2RayAPIOptions  `json:"v2ray_api,omitempty"`
}

type CacheFileOptions struct {
	Enabled     bool   `json:"enabled,omitempty"`
	Path        string `json:"path,omitempty"`
	CacheID     string `json:"cache_id,omitempty"`
	StoreFakeIP bool   `json:"store_fakeip,omitempty"`
	StoreRDRC   bool   `json:"store_rdrc,omitempty"`
	RDRCTimeout string `json:"rdrc_timeout,omitempty"`
}

type ClashAPIOptions struct {
	ExternalController       string `json:"external_controller,omitempty"`
	ExternalUI               string `json:"external_ui,omitempty"`
	ExternalUIDownloadURL    string `json:"external_ui_download_url,omitempty"`
	ExternalUIDownloadDetour string `json:"external_ui_download_detour,omitempty"`
	Secret                   string `json:"secret,omitempty"`
	DefaultMode              string `json:"default_mode,omitempty"`

	// Deprecated: migrated to global cache file
	CacheFile string `json:"cache_file,omitempty"`
	// Deprecated: migrated to global cache file
	CacheID string `json:"cache_id,omitempty"`
	// Deprecated: migrated to global cache file
	StoreMode bool `json:"store_mode,omitempty"`
	// Deprecated: migrated to global cache file
	StoreSelected bool `json:"store_selected,omitempty"`
	// Deprecated: migrated to global cache file
	StoreFakeIP bool `json:"store_fakeip,omitempty"`
}

type V2RayAPIOptions struct {
	Listen string                    `json:"listen,omitempty"`
	Stats  *V2RayStatsServiceOptions `json:"stats,omitempty"`
}

type V2RayStatsServiceOptions struct {
	Enabled   bool     `json:"enabled,omitempty"`
	Inbounds  []string `json:"inbounds,omitempty"`
	Outbounds []string `json:"outbounds,omitempty"`
	Users     []string `json:"users,omitempty"`
}
