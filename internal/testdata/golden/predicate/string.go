package predicate

import "regexp"

// Special characters need to be escaped when adding to
// or querying from RediSearch. The simple approach here uses
// regexp to insert or remove backslashes to strings.

var _punctuationRe = regexp.MustCompile(`(\W)`)
var _escapedPunctuationRe = regexp.MustCompile(`\\(\W)`)

func Escape(s string) string {
	return _punctuationRe.ReplaceAllString(s, "\\$1")
}

func Unescape(s string) string {
	return _escapedPunctuationRe.ReplaceAllString(s, "$1")
}
