package rdsom

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/immanuelhume/rdsomgolden/predicate"
)

type BarFind struct {
	rdb *redis.Client
	p   predicate.Predicate
}

func NewBarFind(rdb *redis.Client) *BarFind {
	return &BarFind{rdb: rdb, p: predicate.TRUE}
}

func (f *BarFind) Where(ps ...predicate.Predicate) *BarFind {
	for _, p := range ps {
		f.p = f.p.And(p)
	}
	return f
}

func (f *BarFind) One(ctx context.Context) (*Bar, error) {
	if f.p.Falsy {
		return nil, nil
	}
	cmd := f.rdb.Do(ctx, "FT.SEARCH", idxBar, f.p.Query, "VERBATIM", "SORTBY", "id", "LIMIT", 0, 1)
	res, err := cmd.Slice()
	if err != nil {
		return nil, err
	}
	if len(res) == 1 {
		return nil, nil
	}
	newBar, err := f.fromRedis(res[2])
	if err != nil {
		return nil, err
	}
	return newBar, nil
}

func (f *BarFind) fromRedis(xs interface{}) (*Bar, error) {
	b := &Bar{}
	ys, ok := xs.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", xs)
	}
	for i := 0; i < len(ys); i += 2 {
		key := ys[i].(string)
		val := predicate.Unescape(ys[i+1].(string))
		f, ok := _fromRedisFuncs[key]
		if !ok {
			continue
		}
		err := f(b, val)
		if err != nil {
			return nil, err
		}
	}
	return b, nil
}

var _fromRedisFuncs = map[string]func(*Bar, string) error{
	"boolField": func(b *Bar, s string) error {
		if s == "0" {
			b.BoolField = false
		} else {
			b.BoolField = true
		}
		return nil
	},
	"floatField": func(b *Bar, s string) error {
		x, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		b.FloatField = x
		return nil
	},
	"floatsField": func(b *Bar, s string) error {
		var xs []float64
		err := json.Unmarshal([]byte(s), &xs)
		if err != nil {
			return err
		}
		b.FloatsField = xs
		return nil
	},
	"intField": func(b *Bar, s string) error {
		x, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		b.IntField = x
		return nil
	},
	"intsField": func(b *Bar, s string) error {
		var xs []int
		err := json.Unmarshal([]byte(s), &xs)
		if err != nil {
			return err
		}
		b.IntsField = xs
		return nil
	},
	"stringField": func(b *Bar, s string) error {
		b.StringField = s
		return nil
	},
	"stringsField": func(b *Bar, s string) error {
		var xs []string
		err := json.Unmarshal([]byte(s), &xs)
		if err != nil {
			return err
		}
		b.StringsField = xs
		return nil
	},
	"timeField": func(b *Bar, s string) error {
		x, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		b.TimeField = time.Unix(x, 0)
		return nil
	},
}
