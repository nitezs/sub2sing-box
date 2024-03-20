package model

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type SortByNumber []Outbound

func (a SortByNumber) Len() int           { return len(a) }
func (a SortByNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByNumber) Less(i, j int) bool { return len(a[i].Outbounds) < len(a[j].Outbounds) }

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
