package bar

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/immanuelhume/rdsomgolden/predicate"
)

func BoolFieldEq(boolField bool) predicate.Predicate {
	var p predicate.Predicate
	if boolField {
		p.Query = `@boolField:[1 1]`
		return p
	}
	p.Query = `@boolField:[0 0]`
	return p
}

func FloatFieldEq(floatField float64) predicate.Predicate {
	var p predicate.Predicate
	p.Query = fmt.Sprintf(`@floatField:[%f %[1]f]`, floatField)
	return p
}

func FloatsFieldEq(floatsField []float64) predicate.Predicate {
	var p predicate.Predicate
	if floatsField == nil {
		p.Query = fmt.Sprintf(`@floatsField:null`)
		return p
	}
	var ss []string
	for _, f := range floatsField {
		ss = append(ss, strconv.FormatFloat(f, 'f', -1, 64))
	}
	p.Query = fmt.Sprintf(`@floatsField:\[%s\]`, predicate.Escape(strings.Join(ss, ",")))
	return p
}

func IntFieldEq(intField int) predicate.Predicate {
	var p predicate.Predicate
	p.Query = fmt.Sprintf(`@intField:[%d %[1]d]`, intField)
	return p
}

func IntsFieldEq(intsField []int) predicate.Predicate {
	var p predicate.Predicate
	if intsField == nil {
		p.Query = fmt.Sprintf(`@intsField:null`)
		return p
	}
	var ss []string
	for _, i := range intsField {
		ss = append(ss, strconv.Itoa(i))
	}
	p.Query = fmt.Sprintf(`@intsField:\[%s\]`, predicate.Escape(strings.Join(ss, ",")))
	return p
}

func StringFieldEq(stringField string) predicate.Predicate {
	var p predicate.Predicate
	if stringField == "" {
		p.Query = "@_empty:{stringField}"
		return p
	}
	p.Query = fmt.Sprintf(`@stringField:%s`, predicate.Escape(stringField))
	return p
}

func StringsFieldEq(stringsField []string) predicate.Predicate {
	var p predicate.Predicate
	if stringsField == nil {
		p.Query = fmt.Sprintf(`@stringsField:null`)
		return p
	}
	var ss []string
	for _, s := range stringsField {
		ss = append(ss, fmt.Sprintf(`"%s"`, s))
	}
	p.Query = fmt.Sprintf(`@stringsField:\[%s\]`, predicate.Escape(strings.Join(ss, ",")))
	return p
}

func TimeFieldEq(timeField time.Time) predicate.Predicate {
	var p predicate.Predicate
	p.Query = fmt.Sprintf(`@timeField:[%d %[1]d]`, timeField.Unix())
	return p
}
