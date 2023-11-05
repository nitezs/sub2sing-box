package model

type SmuxStruct struct {
	Enabled bool `yaml:"enable"`
}

type HTTPOptions struct {
	Method  string              `proxy:"method,omitempty"`
	Path    []string            `proxy:"path,omitempty"`
	Headers map[string][]string `proxy:"headers,omitempty"`
}

type HTTP2Options struct {
	Host []string `proxy:"host,omitempty"`
	Path string   `proxy:"path,omitempty"`
}

type GrpcOptions struct {
	GrpcServiceName string `proxy:"grpc-service-name,omitempty"`
}

type RealityOptions struct {
	PublicKey string `proxy:"public-key"`
	ShortID   string `proxy:"short-id"`
}

type WSOptions struct {
	Path                string            `proxy:"path,omitempty"`
	Headers             map[string]string `proxy:"headers,omitempty"`
	MaxEarlyData        int               `proxy:"max-early-data,omitempty"`
	EarlyDataHeaderName string            `proxy:"early-data-header-name,omitempty"`
}

type Proxy struct {
	Name                string         `yaml:"name,omitempty"`
	Server              string         `yaml:"server,omitempty"`
	Port                int            `yaml:"port,omitempty"`
	Type                string         `yaml:"type,omitempty"`
	Cipher              string         `yaml:"cipher,omitempty"`
	Password            string         `yaml:"password,omitempty"`
	UDP                 bool           `yaml:"udp,omitempty"`
	UUID                string         `yaml:"uuid,omitempty"`
	Network             string         `yaml:"network,omitempty"`
	Flow                string         `yaml:"flow,omitempty"`
	TLS                 bool           `yaml:"tls,omitempty"`
	ClientFingerprint   string         `yaml:"client-fingerprint,omitempty"`
	Plugin              string         `yaml:"plugin,omitempty"`
	PluginOpts          map[string]any `yaml:"plugin-opts,omitempty"`
	Smux                SmuxStruct     `yaml:"smux,omitempty"`
	Sni                 string         `yaml:"sni,omitempty"`
	AllowInsecure       bool           `yaml:"allow-insecure,omitempty"`
	Fingerprint         string         `yaml:"fingerprint,omitempty"`
	SkipCertVerify      bool           `yaml:"skip-cert-verify,omitempty"`
	Alpn                []string       `yaml:"alpn,omitempty"`
	XUDP                bool           `yaml:"xudp,omitempty"`
	Servername          string         `yaml:"servername,omitempty"`
	WSOpts              WSOptions      `yaml:"ws-opts,omitempty"`
	AlterID             int            `yaml:"alterId,omitempty"`
	GrpcOpts            GrpcOptions    `yaml:"grpc-opts,omitempty"`
	RealityOpts         RealityOptions `yaml:"reality-opts,omitempty"`
	Protocol            string         `yaml:"protocol,omitempty"`
	Obfs                string         `yaml:"obfs,omitempty"`
	ObfsParam           string         `yaml:"obfs-param,omitempty"`
	ProtocolParam       string         `yaml:"protocol-param,omitempty"`
	Remarks             []string       `yaml:"remarks,omitempty"`
	HTTPOpts            HTTPOptions    `yaml:"http-opts,omitempty"`
	HTTP2Opts           HTTP2Options   `yaml:"h2-opts,omitempty"`
	PacketAddr          bool           `yaml:"packet-addr,omitempty"`
	PacketEncoding      string         `yaml:"packet-encoding,omitempty"`
	GlobalPadding       bool           `yaml:"global-padding,omitempty"`
	AuthenticatedLength bool           `yaml:"authenticated-length,omitempty"`
	UDPOverTCP          bool           `yaml:"udp-over-tcp,omitempty"`
	UDPOverTCPVersion   int            `yaml:"udp-over-tcp-version,omitempty"`
	SubName             string         `yaml:"-"`
	Up                  string         `yaml:"up,omitempty"`
	Down                string         `yaml:"down,omitempty"`
	CustomCA            string         `yaml:"ca,omitempty"`
	CustomCAString      string         `yaml:"ca-str,omitempty"`
	CWND                int            `yaml:"cwnd,omitempty"`
}
