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

func IsAlpha(alphabet string) (Filter, error) {
	return Regex(`^[` + regexp.QuoteMeta(alphabet) + `]+$`)
}

func IsNumeric() Filter {
	filter, _ := Regex(`^[-+]?[0-9]*\.?[0-9]+$`)
	return filter
}

func Regex(text string) (Filter, error) {
	re, err := regexp.Compile(text)
	if err != nil {
		return nil, err
	}

	return &filterRegex{re: re}, nil
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

func MinLength(minLen int) Filter {
	return &filterMinLength{minLen: minLen}
}

func MaxLength(maxLen int) Filter {
	return &filterMaxLength{maxLen: maxLen}
}

func Allow(filter Filter) Filter {
	if filter == nil {
		return AlwaysAllow
	}

	return filter
}

func Deny(filter Filter) Filter {
	if filter == nil {
		return AlwaysAllow
	}

	return &filterNot{filter: filter}
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
