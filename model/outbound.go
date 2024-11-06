package model

import (
	C "github.com/nitezs/sub2sing-box/constant"
	"github.com/sagernet/sing-box/option"
)

type Outbound struct {
	option.Outbound
}

func (h *Outbound) GetOutbounds() []string {
	if h.Type == C.TypeSelector {
		return h.SelectorOptions.Outbounds
	}
	if h.Type == C.TypeURLTest {
		return h.URLTestOptions.Outbounds
	}
	return nil
}

func (h *Outbound) SetOutbounds(outbounds []string) {
	if h.Type == C.TypeSelector {
		h.SelectorOptions.Outbounds = outbounds
	}
	if h.Type == C.TypeURLTest {
		h.URLTestOptions.Outbounds = outbounds
	}
}
