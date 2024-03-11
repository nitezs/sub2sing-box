package model

type ConvertRequest struct {
	Subscriptions []string          `form:"subscription" json:"subscription"`
	Proxies       []string          `form:"proxy" json:"proxy"`
	Template      string            `form:"template" json:"template"`
	Delete        string            `form:"delete" json:"delete"`
	Rename        map[string]string `form:"rename" json:"rename"`
}
