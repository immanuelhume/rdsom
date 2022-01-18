package rdsom_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	rdsom "github.com/immanuelhume/rdsomgolden"
	"github.com/immanuelhume/rdsomgolden/bar"
	"github.com/immanuelhume/rdsomgolden/predicate"
)

func TestFindOneByEq(t *testing.T) {
	ctx := context.Background()
	bb := _client.Bar
	tests := []struct {
		desc   string
		create *rdsom.BarCreate
		where  predicate.Predicate
		want   *rdsom.Bar
		equals func(*rdsom.Bar, *rdsom.Bar) bool // a custom compare function
	}{
		{
			desc:   "by bool field [true]",
			create: bb.Create().SetBoolField(true),
			where:  bar.BoolFieldEq(true),
			want:   &rdsom.Bar{BoolField: true},
		},
		{
			desc:   "by bool field [false]",
			create: bb.Create().SetBoolField(false),
			where:  bar.BoolFieldEq(false),
			want:   &rdsom.Bar{BoolField: false},
		},

		{
			desc:   "by float64 field",
			create: bb.Create().SetFloatField(1.2),
			where:  bar.FloatFieldEq(1.2),
			want:   &rdsom.Bar{FloatField: 1.2},
		},
		{
			desc:   "by float64 field [negative]",
			create: bb.Create().SetFloatField(-1.2),
			where:  bar.FloatFieldEq(-1.2),
			want:   &rdsom.Bar{FloatField: -1.2},
		},
		{
			desc:   "by float64 field [zero]",
			create: bb.Create().SetFloatField(0),
			where:  bar.FloatFieldEq(0),
			want:   &rdsom.Bar{FloatField: 0},
		},

		{
			desc:   "by []float64 field",
			create: bb.Create().SetFloatsField([]float64{2.72, 3.14}),
			where:  bar.FloatsFieldEq([]float64{2.72, 3.14}),
			want:   &rdsom.Bar{FloatsField: []float64{2.72, 3.14}},
		},
		{
			desc:   "by []float64 field [zero values]",
			create: bb.Create().SetFloatsField([]float64{0, 0}),
			where:  bar.FloatsFieldEq([]float64{0, 0}),
			want:   &rdsom.Bar{FloatsField: []float64{0, 0}},
		},
		{
			desc:   "by []float64 field [empty slice]",
			create: bb.Create().SetFloatsField([]float64{}),
			where:  bar.FloatsFieldEq([]float64{}),
			want:   &rdsom.Bar{FloatsField: []float64{}},
		},
		{
			desc:   "by []float64 field [nil slice]",
			create: bb.Create().SetFloatsField(nil),
			where:  bar.FloatsFieldEq(nil),
			want:   &rdsom.Bar{},
		},

		{
			desc:   "by int field",
			create: bb.Create().SetIntField(1),
			where:  bar.IntFieldEq(1),
			want:   &rdsom.Bar{IntField: 1},
		},
		{
			desc:   "by int field [negative]",
			create: bb.Create().SetIntField(-1),
			where:  bar.IntFieldEq(-1),
			want:   &rdsom.Bar{IntField: -1},
		},
		{
			desc:   "by int field [zero]",
			create: bb.Create().SetIntField(0),
			where:  bar.IntFieldEq(0),
			want:   &rdsom.Bar{IntField: 0},
		},

		{
			desc:   "by []int field",
			create: bb.Create().SetIntsField([]int{1, 2}),
			where:  bar.IntsFieldEq([]int{1, 2}),
			want:   &rdsom.Bar{IntsField: []int{1, 2}},
		},
		{
			desc:   "by []int field [zero values]",
			create: bb.Create().SetIntsField([]int{0, 0}),
			where:  bar.IntsFieldEq([]int{0, 0}),
			want:   &rdsom.Bar{IntsField: []int{0, 0}},
		},
		{
			desc:   "by []int field [negative values]",
			create: bb.Create().SetIntsField([]int{-1, -2}),
			where:  bar.IntsFieldEq([]int{-1, -2}),
			want:   &rdsom.Bar{IntsField: []int{-1, -2}},
		},
		{
			desc:   "by []int field [empty slice]",
			create: bb.Create().SetIntsField([]int{}),
			where:  bar.IntsFieldEq([]int{}),
			want:   &rdsom.Bar{IntsField: []int{}},
		},
		{
			desc:   "by []int field [nil slice]",
			create: bb.Create().SetIntsField(nil),
			where:  bar.IntsFieldEq(nil),
			want:   &rdsom.Bar{},
		},

		{
			desc:   "by string field",
			create: bb.Create().SetStringField("foo bar"),
			where:  bar.StringFieldEq("foo bar"),
			want:   &rdsom.Bar{StringField: "foo bar"},
		},
		{
			desc:   "by string field [special characters]",
			create: bb.Create().SetStringField(`,.<>{}[]"':;!@#$%^&*()-+=~`),
			where:  bar.StringFieldEq(`,.<>{}[]"':;!@#$%^&*()-+=~`),
			want:   &rdsom.Bar{StringField: `,.<>{}[]"':;!@#$%^&*()-+=~`},
		},
		{
			desc:   "by string field [empty string]",
			create: bb.Create().SetStringField(""),
			where:  bar.StringFieldEq(""),
			want:   &rdsom.Bar{StringField: ""},
		},
		{
			desc:   "by string field [whitespace only]",
			create: bb.Create().SetStringField("  "),
			where:  bar.StringFieldEq("  "),
			want:   &rdsom.Bar{StringField: "  "},
		},

		{
			desc:   "by []string field",
			create: bb.Create().SetStringsField([]string{"foo", "bar"}),
			where:  bar.StringsFieldEq([]string{"foo", "bar"}),
			want:   &rdsom.Bar{StringsField: []string{"foo", "bar"}},
		},
		{
			desc:   "by []string field [empty strings]",
			create: bb.Create().SetStringsField([]string{"", ""}),
			where:  bar.StringsFieldEq([]string{"", ""}),
			want:   &rdsom.Bar{StringsField: []string{"", ""}},
		},
		{
			desc:   "by []string field [empty slice]",
			create: bb.Create().SetStringsField([]string{}),
			where:  bar.StringsFieldEq([]string{}),
			want:   &rdsom.Bar{StringsField: []string{}},
		},
		{
			desc:   "by []string field [nil slice]",
			create: bb.Create().SetStringsField(nil),
			where:  bar.StringsFieldEq(nil),
			want:   &rdsom.Bar{},
		},

		{
			desc:   "by time.Time field",
			create: bb.Create().SetTimeField(time.Now()),
			where:  bar.TimeFieldEq(time.Now()),
			want:   &rdsom.Bar{TimeField: time.Now()},
			equals: func(b1, b2 *rdsom.Bar) bool {
				return b1.TimeField.Sub(b2.TimeField) < time.Microsecond
			},
		},
	}

	// Ensure that migrations are ran.
	err := _client.MigrateAll(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.desc+" should return one", func(t *testing.T) {
			// First, insert so we have something to find.
			key, err := tt.create.Save(ctx)
			if err != nil {
				t.Fatal(err)
			}
			defer _client.Redis().Del(ctx, key) // make sure we delete it

			got, err := _client.Bar.Find().Where(tt.where).One(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if tt.equals != nil {
				if !tt.equals(got, tt.want) {
					t.Errorf("got %+v, want %+v", got, tt.want)
				}
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})

		// Now, test if it returns nil if nothing is found.
		// Since we deleted the hash, it should return nil.
		t.Run(tt.desc+" should return nil", func(t *testing.T) {
			got, err := _client.Bar.Find().Where(tt.where).One(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if got != nil {
				t.Errorf("got %+v, want nil", got)
			}
		})
	}
}
