package model

type Hysteria2Obfs struct {
	Type     string `json:"type,omitempty"`
	Password string `json:"password,omitempty"`
}

type Hysteria2 struct {
	Type        string              `json:"type"`
	Tag         string              `json:"tag,omitempty"`
	Server      string              `json:"server"`
	ServerPort  uint16              `json:"server_port"`
	UpMbps      int                 `json:"up_mbps,omitempty"`
	DownMbps    int                 `json:"down_mbps,omitempty"`
	Obfs        *Hysteria2Obfs      `json:"obfs,omitempty"`
	Password    string              `json:"password,omitempty"`
	Network     string              `json:"network,omitempty"`
	TLS         *OutboundTLSOptions `json:"tls,omitempty"`
	BrutalDebug bool                `json:"brutal_debug,omitempty"`
}

// func (h *Hysteria2OutboundOptions) MarshalJSON() ([]byte, error) {
// 	val := reflect.ValueOf(h)
// 	out := make(map[string]interface{})
// 	typ := val.Type()
// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)
// 		fieldType := typ.Field(i)
// 		if field.Kind() == reflect.Struct {
// 			for j := 0; j < field.NumField(); j++ {
// 				subField := field.Field(j)
// 				subFieldType := fieldType.Type.Field(j)
// 				jsonTag := subFieldType.Tag.Get("json")
// 				if jsonTag != "" && jsonTag != "-" {
// 					out[jsonTag] = subField.Interface()
// 				}
// 			}
// 		} else {
// 			jsonTag := fieldType.Tag.Get("json")
// 			if jsonTag != "" && jsonTag != "-" {
// 				out[jsonTag] = field.Interface()
// 			}
// 		}
// 	}
// 	return json.Marshal(out)
// }
