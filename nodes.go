package filters

import (
	"regexp"
	"strings"
)

type filterOr []Filter

func (that filterOr) IsMatch(text string) bool {
	for _, re := range that {
		if re.IsMatch(text) {
			return true
		}
	}
	return false
}

type filterAnd []Filter

func (that filterAnd) IsMatch(text string) bool {
	for _, re := range that {
		if !re.IsMatch(text) {
			return false
		}
	}
	return true
}

type filterNot struct {
	filter Filter
}

func (that *filterNot) IsMatch(text string) bool {
	return !that.filter.IsMatch(text)
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
