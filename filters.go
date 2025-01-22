package filters

import (
	"regexp"
)

type Filter interface {
	IsMatch(text string) bool
}

var (
	AlwaysAllow Filter = &filterConst{allow: true}
	AlwaysDeny  Filter = &filterConst{allow: false}
)

func AND(filters ...Filter) Filter {
	return filterAnd(filters)
}

func OR(filters ...Filter) Filter {
	return filterOr(filters)
}

func NOT(filter Filter) Filter {
	return &filterNot{filter: filter}
}

func Regexp(text string) (Filter, error) {
	re, err := regexp.Compile(text)
	if err != nil {
		return nil, err
	}

	return &filterRegexp{re: re}, nil
}

func Exact(text string) Filter {
	return &filterExact{text: text}
}

func Prefix(text string) Filter {
	return &filterPrefix{text: text}
}

func Suffix(text string) Filter {
	return &filterSuffix{text: text}
}

func Allow(filter Filter) Filter {
	return &filterAllowDeny{
		allow: filter,
	}
}

func Deny(filter Filter) Filter {
	return &filterAllowDeny{
		deny: filter,
	}
}

func AllowDeny(allow Filter, deny Filter) Filter {
	return &filterAllowDeny{
		allow: allow,
		deny:  deny,
	}
}

func Must(f Filter, err error) Filter {
	if err != nil {
		panic(err)
	}
	return f
}
