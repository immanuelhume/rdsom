package rdsom_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	rdsom "github.com/immanuelhume/rdsomgolden"
	"github.com/immanuelhume/rdsomgolden/predicate"
)

func TestCreateBar(t *testing.T) {
	ctx := context.Background()
	bar := _client.Bar
	tests := []struct {
		desc      string
		fieldName string
		builder   *rdsom.BarCreate
		want      string
	}{
		{
			desc:      "set bool field [true]",
			fieldName: "boolField",
			builder:   bar.Create().SetBoolField(true),
			want:      "1",
		},
		{
			desc:      "set bool field [false]",
			fieldName: "boolField",
			builder:   bar.Create().SetBoolField(false),
			want:      "0",
		},

		{
			desc:      "set float64 field",
			fieldName: "floatField",
			builder:   bar.Create().SetFloatField(3.14),
			want:      "3.14",
		},
		{
			desc:      "set float64 field [zero]",
			fieldName: "floatField",
			builder:   bar.Create().SetFloatField(0.0),
			want:      "0",
		},
		{
			desc:      "set float64 field [negative value]",
			fieldName: "floatField",
			builder:   bar.Create().SetFloatField(-3.14),
			want:      "-3.14",
		},

		{
			desc:      "set []float64 field",
			fieldName: "floatsField",
			builder:   bar.Create().SetFloatsField([]float64{2.72, 3.14}),
			want:      predicate.Escape("[2.72,3.14]"),
		},
		{
			desc:      "set []float64 field [with zero values]",
			fieldName: "floatsField",
			builder:   bar.Create().SetFloatsField([]float64{0.0, 0.0}),
			want:      predicate.Escape("[0,0]"),
		},
		{
			desc:      "set []float64 field [empty slice]",
			fieldName: "floatsField",
			builder:   bar.Create().SetFloatsField([]float64{}),
			want:      predicate.Escape("[]"),
		},
		{
			desc:      "set []float64 field [nil slice]",
			fieldName: "floatsField",
			builder:   bar.Create().SetFloatsField(nil),
			want:      "null",
		},

		{
			desc:      "set int field",
			fieldName: "intField",
			builder:   bar.Create().SetIntField(1),
			want:      "1",
		},
		{
			desc:      "set int field [zero]",
			fieldName: "intField",
			builder:   bar.Create().SetIntField(0),
			want:      "0",
		},
		{
			desc:      "set int field [negative]",
			fieldName: "intField",
			builder:   bar.Create().SetIntField(-1),
			want:      "-1",
		},

		{
			desc:      "set []int field",
			fieldName: "intsField",
			builder:   bar.Create().SetIntsField([]int{1, 2, 3}),
			want:      predicate.Escape("[1,2,3]"),
		},
		{
			desc:      "set []int field [with zero values]",
			fieldName: "intsField",
			builder:   bar.Create().SetIntsField([]int{0, 0}),
			want:      predicate.Escape("[0,0]"),
		},
		{
			desc:      "set []int field [empty slice]",
			fieldName: "intsField",
			builder:   bar.Create().SetIntsField([]int{}),
			want:      predicate.Escape("[]"),
		},
		{
			desc:      "set []int field [nil slice]",
			fieldName: "intsField",
			builder:   bar.Create().SetIntsField(nil),
			want:      "null",
		},

		{
			desc:      "set string field",
			fieldName: "stringField",
			builder:   bar.Create().SetStringField("foo bar"),
			want:      predicate.Escape("foo bar"),
		},
		{
			desc:      "set string field [empty string]",
			fieldName: "stringField",
			builder:   bar.Create().SetStringField(""),
			want:      "",
		},
		{
			desc:      "set string field [special characters]",
			fieldName: "stringField",
			builder:   bar.Create().SetStringField(`,.<>{}[]"':;!@#$%^&*()-+=~`),
			want:      predicate.Escape(`,.<>{}[]"':;!@#$%^&*()-+=~`),
		},
		{
			desc:      "set string field [whitespace only]",
			fieldName: "stringField",
			builder:   bar.Create().SetStringField("  "),
			want:      predicate.Escape("  "),
		},

		{
			desc:      "set []string field",
			fieldName: "stringsField",
			builder:   bar.Create().SetStringsField([]string{"foo", "bar"}),
			want:      predicate.Escape(`["foo","bar"]`),
		},
		{
			desc:      "set []string field [empty strings]",
			fieldName: "stringsField",
			builder:   bar.Create().SetStringsField([]string{"", ""}),
			want:      predicate.Escape(`["",""]`),
		},
		{
			desc:      "set []string [empty slice]",
			fieldName: "stringsField",
			builder:   bar.Create().SetStringsField([]string{}),
			want:      predicate.Escape("[]"),
		},
		{
			desc:      "set []string field [nil slice]",
			fieldName: "stringsField",
			builder:   bar.Create().SetStringsField(nil),
			want:      "null",
		},

		{
			desc:      "set time.Time field",
			fieldName: "timeField",
			builder:   bar.Create().SetTimeField(time.Now()),
			want:      fmt.Sprintf("%d", time.Now().Unix()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			key, err := tt.builder.Save(ctx)
			if err != nil {
				t.Fatal(err)
			}
			defer _client.Redis().Del(ctx, key)

			got, err := _client.Redis().HGet(ctx, key, tt.fieldName).Result()
			if err != nil {
				t.Fatal(err)
			}

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
