package internal

import (
	"reflect"
	"regexp"
	"strings"
	"text/template"
)

var funcs template.FuncMap = template.FuncMap{
	"toCamelCase":  toCamelCase,
	"toPascalCase": toPascalCase,
	"toSnakeCase":  toSnakeCase,
	"toInitCase":   toInitCase,
	"toLower":      strings.ToLower,
	"title":        strings.Title,
	"addInts":      addInts,
}

var matchSymbolStart = regexp.MustCompile("[^0-9A-Za-z]([0-9A-Za-z])")
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toCamelCase(s string) string {
	camel := matchSymbolStart.ReplaceAllStringFunc(s, func(s string) string {
		return strings.ToUpper(string(s[1]))
	})
	return strings.ToLower(string(camel[0])) + camel[1:]
}

func toPascalCase(s string) string {
	return strings.Title(toCamelCase(s))
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake = strings.ReplaceAll(snake, " ", "_")
	return strings.ToLower(snake)
}

// toInitCase returns the first character of a string in lower case.
// E.g. "Foo" -> "f".
func toInitCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(string(s[0]))
}

// Returns the name of a thing.
func name(x interface{}) string {
	t := reflect.TypeOf(x)
	if reflect.TypeOf(x).Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}

func addInts(xs ...int) int {
	var ret int
	for _, x := range xs {
		ret += x
	}
	return ret
}
