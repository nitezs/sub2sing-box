package model

import (
	"github.com/sagernet/sing-box/option"
)

type Outbound struct {
	option.Outbound
}

func (h *Outbound) GetOutbounds() []string {
	switch v := h.Options.(type) {
	case option.URLTestOutboundOptions:
		return v.Outbounds
	case option.SelectorOutboundOptions:
		return v.Outbounds
	}
	return nil
}

func (h *Outbound) SetOutbounds(outbounds []string) {
	switch v := h.Options.(type) {
	case option.URLTestOutboundOptions:
		v.Outbounds = outbounds
		h.Options = v
	case option.SelectorOutboundOptions:
		v.Outbounds = outbounds
		h.Options = v
	}
}
