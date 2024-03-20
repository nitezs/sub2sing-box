package model

import (
	C "sub2sing-box/constant"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type SortByNumber []Outbound

func (a SortByNumber) Len() int      { return len(a) }
func (a SortByNumber) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByNumber) Less(i, j int) bool {
	var size1, size2 int
	if a[i].Type == C.TypeSelector {
		size1 = len(a[i].SelectorOptions.Outbounds)
	}
	if a[i].Type == C.TypeURLTest {
		size1 = len(a[j].URLTestOptions.Outbounds)
	}
	if a[j].Type == C.TypeSelector {
		size2 = len(a[j].SelectorOptions.Outbounds)
	}
	if a[j].Type == C.TypeURLTest {
		size2 = len(a[j].URLTestOptions.Outbounds)
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
