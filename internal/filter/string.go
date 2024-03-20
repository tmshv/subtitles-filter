package filter

import (
	"strings"
	"unicode/utf8"
)

type minlen struct {
	val int
}

func (f *minlen) Test(text string) bool {
    return utf8.RuneCountInString(text) >= f.val
}

func MinLen(val int) Filter {
    return &minlen{val}
}

type substr struct {
	val string
}

func (f *substr) Test(text string) bool {
	return strings.Contains(text, f.val)
}

func Substr(val string) Filter {
    return &substr{val}
}

