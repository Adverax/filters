package filters

import (
	"regexp"
)

type Builder struct {
	allow filterOr
	deny  filterOr
	err   error
}

func (that *Builder) Allow(filter ...Filter) *Builder {
	that.allow = append(that.allow, filter...)
	return that
}

func (that *Builder) Deny(filter ...Filter) *Builder {
	that.deny = append(that.deny, filter...)
	return that
}

func (that *Builder) AllowExact(text string) *Builder {
	return that.Allow(&filterExact{text: text})
}

func (that *Builder) DenyExact(text string) *Builder {
	return that.Deny(&filterExact{text: text})
}

func (that *Builder) AllowRegexp(text string) *Builder {
	re, err := regexp.Compile(text)
	if err != nil {
		that.error(err)
		return that
	}
	return that.Allow(&filterRegexp{re: re})
}

func (that *Builder) DenyRegexp(text string) *Builder {
	re, err := regexp.Compile(text)
	if err != nil {
		that.error(err)
		return that
	}
	return that.Deny(&filterRegexp{re: re})
}

func (that *Builder) AllowPrefix(text string) *Builder {
	return that.Allow(&filterPrefix{text: text})
}

func (that *Builder) DenyPrefix(text string) *Builder {
	return that.Deny(&filterPrefix{text: text})
}

func (that *Builder) AllowSuffix(text string) *Builder {
	return that.Allow(&filterSuffix{text: text})
}

func (that *Builder) DenySuffix(text string) *Builder {
	return that.Deny(&filterSuffix{text: text})
}

func (that *Builder) Build() (Filter, error) {
	if that.err != nil {
		return nil, that.err
	}

	return &filterAllowDeny{
		allow: that.build(that.allow),
		deny:  that.build(that.deny),
	}, nil
}

func (that *Builder) build(filters filterOr) Filter {
	switch len(filters) {
	case 0:
		return nil
	case 1:
		return filters[0]
	default:
		return filters
	}
}

func (that *Builder) error(err error) {
	if that.err == nil {
		that.err = err
	}
}

func NewBuilder() *Builder {
	return &Builder{}
}
