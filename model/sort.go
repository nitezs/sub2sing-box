package model

import (
	"github.com/sagernet/sing-box/option"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type SortByNumber []Outbound

func (a SortByNumber) Len() int      { return len(a) }
func (a SortByNumber) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByNumber) Less(i, j int) bool {
	var size1, size2 int
	switch v := a[i].Options.(type) {
	case option.SelectorOutboundOptions:
		size1 = len(v.Outbounds)
	case option.URLTestOutboundOptions:
		size1 = len(v.Outbounds)
	}
	switch v := a[j].Options.(type) {
	case option.SelectorOutboundOptions:
		size2 = len(v.Outbounds)
	case option.URLTestOutboundOptions:
		size2 = len(v.Outbounds)
	}
	return size1 < size2
}

type SortByTag []Outbound

func (a SortByTag) Len() int      { return len(a) }
func (a SortByTag) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByTag) Less(i, j int) bool {
	tags := []language.Tag{
		language.English,
		language.Chinese,
	}
	matcher := language.NewMatcher(tags)
	bestMatch, _, _ := matcher.Match(language.Make("zh"))
	c := collate.New(bestMatch)
	return c.CompareString(a[i].Tag, a[j].Tag) < 0
}
