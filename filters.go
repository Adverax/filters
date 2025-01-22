package filters

import (
	"regexp"
	"strings"
)

type Filter interface {
	IsMatch(text string) bool
}

type filterMulti []Filter

func (that filterMulti) IsMatch(text string) bool {
	for _, re := range that {
		if re.IsMatch(text) {
			return true
		}
	}
	return false
}

type filterRegexp struct {
	re *regexp.Regexp
}

func (that *filterRegexp) IsMatch(text string) bool {
	return that.re.Match([]byte(text))
}

type filterExact struct {
	text string
}

func (that *filterExact) IsMatch(text string) bool {
	return that.text == text
}

type filterPrefix struct {
	text string
}

func (that *filterPrefix) IsMatch(text string) bool {
	return strings.HasPrefix(text, that.text)
}

type filterSuffix struct {
	text string
}

func (that *filterSuffix) IsMatch(text string) bool {
	return strings.HasSuffix(text, that.text)
}

type filterConst struct {
	allow bool
}

func (that *filterConst) IsMatch(text string) bool {
	return that.allow
}

type filterAllowDeny struct {
	allow Filter
	deny  Filter
}

func (that *filterAllowDeny) IsMatch(text string) bool {
	if that.deny != nil && that.deny.IsMatch(text) {
		return false
	}
	return that.allow == nil || that.allow.IsMatch(text)
}

var (
	AllowFilter Filter = &filterConst{allow: true}
	DenyFilter  Filter = &filterConst{allow: false}
)
