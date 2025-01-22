package filters

import (
	"fmt"
	"regexp"
)

const (
	MatchFilterExact FilterType = iota
	MatchFilterRegexp
	MatchFilterPrefix
	MatchFilterSuffix
)

type FilterType int

type def struct {
	text string
	typ  FilterType
}

type Builder struct {
	allow []*def
	deny  []*def
}

func (that *Builder) Allow(typ FilterType, text string) *Builder {
	that.allow = append(that.allow, &def{text: text, typ: typ})
	return that
}

func (that *Builder) Deny(typ FilterType, text string) *Builder {
	that.deny = append(that.deny, &def{text: text, typ: typ})
	return that
}

func (that *Builder) AllowExact(text string) *Builder {
	return that.Allow(MatchFilterExact, text)
}

func (that *Builder) DenyExact(text string) *Builder {
	return that.Deny(MatchFilterExact, text)
}

func (that *Builder) AllowRegexp(text string) *Builder {
	return that.Allow(MatchFilterRegexp, text)
}

func (that *Builder) DenyRegexp(text string) *Builder {
	return that.Deny(MatchFilterRegexp, text)
}

func (that *Builder) AllowPrefix(text string) *Builder {
	return that.Allow(MatchFilterPrefix, text)
}

func (that *Builder) DenyPrefix(text string) *Builder {
	return that.Deny(MatchFilterPrefix, text)
}

func (that *Builder) AllowSuffix(text string) *Builder {
	return that.Allow(MatchFilterSuffix, text)
}

func (that *Builder) DenySuffix(text string) *Builder {
	return that.Deny(MatchFilterSuffix, text)
}

func (that *Builder) Build() (Filter, error) {
	allow, err := that.build(that.allow)
	if err != nil {
		return nil, fmt.Errorf("NewFilter: %w", err)
	}

	deny, err := that.build(that.deny)
	if err != nil {
		return nil, fmt.Errorf("NewFilter: %w", err)
	}

	return &filterAllowDeny{
		allow: allow,
		deny:  deny,
	}, nil
}

func (that *Builder) build(
	defs []*def,
) (filter Filter, err error) {
	switch len(defs) {
	case 0:
		return nil, nil
	case 1:
		return that.newFilter(defs[0])
	default:
		filters := make(filterOr, 0)
		for _, def := range defs {
			filter, err := that.newFilter(def)
			if err != nil {
				return nil, fmt.Errorf("newFilter: %w", err)
			}
			filters = append(filters, filter)
		}

		return filters, nil
	}
}

func (that *Builder) newFilter(def *def) (Filter, error) {
	switch def.typ {
	case MatchFilterRegexp:
		re, err := regexp.Compile(def.text)
		if err != nil {
			return nil, fmt.Errorf("regexp.Compile: %w", err)
		}
		return &filterRegexp{re: re}, nil
	case MatchFilterPrefix:
		return &filterPrefix{text: def.text}, nil
	case MatchFilterSuffix:
		return &filterSuffix{text: def.text}, nil
	default:
		return &filterExact{text: def.text}, nil
	}
}

func NewBuilder() *Builder {
	return &Builder{}
}
