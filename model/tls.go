package model

type OutboundTLSOptions struct {
	Enabled         bool                    `json:"enabled,omitempty"`
	DisableSNI      bool                    `json:"disable_sni,omitempty"`
	ServerName      string                  `json:"server_name,omitempty"`
	Insecure        bool                    `json:"insecure,omitempty"`
	ALPN            Listable[string]        `json:"alpn,omitempty"`
	MinVersion      string                  `json:"min_version,omitempty"`
	MaxVersion      string                  `json:"max_version,omitempty"`
	CipherSuites    Listable[string]        `json:"cipher_suites,omitempty"`
	Certificate     Listable[string]        `json:"certificate,omitempty"`
	CertificatePath string                  `json:"certificate_path,omitempty"`
	ECH             *OutboundECHOptions     `json:"ech,omitempty"`
	UTLS            *OutboundUTLSOptions    `json:"utls,omitempty"`
	Reality         *OutboundRealityOptions `json:"reality,omitempty"`
}

type OutboundECHOptions struct {
	Enabled                     bool             `json:"enabled,omitempty"`
	PQSignatureSchemesEnabled   bool             `json:"pq_signature_schemes_enabled,omitempty"`
	DynamicRecordSizingDisabled bool             `json:"dynamic_record_sizing_disabled,omitempty"`
	Config                      Listable[string] `json:"config,omitempty"`
	ConfigPath                  string           `json:"config_path,omitempty"`
}

type OutboundUTLSOptions struct {
	Enabled     bool   `json:"enabled,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
}

type OutboundRealityOptions struct {
	Enabled   bool   `json:"enabled,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
	ShortID   string `json:"short_id,omitempty"`
}
