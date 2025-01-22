# Filters

`github.com/adverax/filters` is a lightweight and efficient library for applying and managing filters for strings. This package is especially useful for building filtering mechanisms in web applications, APIs, or any system that requires flexible and reusable filter definitions.

## Features

- Simple and declarative filter definitions.
- Support for chaining and combining filters.
- Flexible configuration options for custom logic.
- Built-in utilities for common filtering tasks.

## Installation

Install the package using `go get`:

```bash
go get github.com/adverax/filters
```

## Overview
The filter mechanism based on allow/deny relations works by defining rules that either allow or deny certain text patterns. These rules are combined to determine whether a given text should be accepted or rejected. Here is a brief description of the components involved:  
Filter Interface: This interface defines a method IsMatch(text string) bool that checks if a given text matches the filter criteria.  
Filter Types: Various filter types implement the Filter interface, such as:  
- filterExact: Matches text exactly.
- filterRegexp: Matches text using regular expressions.
- filterPrefix: Matches text with a specific prefix.
- filterSuffix: Matches text with a specific suffix.
Filter Builder: A builder pattern is used to create complex filters by combining multiple allow and deny rules. The builder provides methods to add different types of filters.  
Allow/Deny Logic: The filterAllowDeny type combines allow and deny filters. It first checks the deny filters; if any deny filter matches, the text is rejected. If no deny filter matches, it checks the allow filters; if any allow filter matches, the text is accepted.  
Constant Filters: AllowFilter and DenyFilter are predefined filters that always allow or deny text, respectively.

## Usage
See filters_test.go file.

---

Happy filtering!
