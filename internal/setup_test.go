package internal_test

import "github.com/immanuelhume/rdsom/internal"

var _barSchema = internal.Schema{
	Name: "Bar",
	Fields: []internal.Field{
		{
			Name:           "BoolField",
			JsonName:       "boolField",
			GoType:         "bool",
			RediSearchOpts: []string{"NUMERIC"},
		},
		{
			Name:           "FloatField",
			JsonName:       "floatField",
			GoType:         "float64",
			RediSearchOpts: []string{"NUMERIC"},
		},
		{
			Name:           "FloatsField",
			JsonName:       "floatsField",
			GoType:         "[]float64",
			RediSearchOpts: []string{"TEXT"},
		},
		{
			Name:           "IntField",
			JsonName:       "intField",
			GoType:         "int",
			RediSearchOpts: []string{"NUMERIC"},
		},
		{
			Name:           "IntsField",
			JsonName:       "intsField",
			GoType:         "[]int",
			RediSearchOpts: []string{"TEXT"},
		},
		{
			Name:           "StringField",
			JsonName:       "stringField",
			GoType:         "string",
			RediSearchOpts: []string{"TEXT"},
		},
		{
			Name:           "StringsField",
			JsonName:       "stringsField",
			GoType:         "[]string",
			RediSearchOpts: []string{"TEXT"},
		},
		{
			Name:           "TimeField",
			JsonName:       "timeField",
			GoType:         "time.Time",
			RediSearchOpts: []string{"NUMERIC"},
		},
	},
}

const codeGenAt uint64 = 1642524233
