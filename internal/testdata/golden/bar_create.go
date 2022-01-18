package rdsom

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/immanuelhume/rdsomgolden/predicate"
	uuid "github.com/lithammer/shortuuid/v4"
)

type BarCreate struct {
	rdb         *redis.Client
	values      map[string]interface{}
	errors      []error
	emptyFields []string
}

func NewBarCreate(rdb *redis.Client) *BarCreate {
	return &BarCreate{rdb: rdb, values: make(map[string]interface{}, 8), emptyFields: make([]string, 0, 1)}
}

func (c *BarCreate) SetBoolField(boolField bool) *BarCreate {
	c.values["boolField"] = boolField
	return c
}

func (c *BarCreate) SetFloatField(floatField float64) *BarCreate {
	c.values["floatField"] = floatField
	return c
}

func (c *BarCreate) SetFloatsField(floatsField []float64) *BarCreate {
	xs, err := json.Marshal(floatsField)
	if err != nil {
		c.errors = append(c.errors, err)
		return c
	}
	c.values["floatsField"] = predicate.Escape(string(xs))
	return c
}

func (c *BarCreate) SetIntField(intField int) *BarCreate {
	c.values["intField"] = intField
	return c
}

func (c *BarCreate) SetIntsField(intsField []int) *BarCreate {
	xs, err := json.Marshal(intsField)
	if err != nil {
		c.errors = append(c.errors, err)
		return c
	}
	c.values["intsField"] = predicate.Escape(string(xs))
	return c
}

func (c *BarCreate) SetStringField(stringField string) *BarCreate {
	if stringField == "" {
		c.emptyFields = append(c.emptyFields, "stringField")
	}
	c.values["stringField"] = predicate.Escape(stringField)
	return c
}

func (c *BarCreate) SetStringsField(stringsField []string) *BarCreate {
	xs, err := json.Marshal(stringsField)
	if err != nil {
		c.errors = append(c.errors, err)
		return c
	}
	c.values["stringsField"] = predicate.Escape(string(xs))
	return c
}

func (c *BarCreate) SetTimeField(timeField time.Time) *BarCreate {
	c.values["timeField"] = timeField.Unix()
	return c
}

func (c *BarCreate) Save(ctx context.Context) (string, error) {
	if len(c.emptyFields) != 0 {
		c.values["_empty"] = strings.Join(c.emptyFields, ",")
	}
	id := uuid.New()
	c.values["id"] = id
	key, err := c.genKey()
	if err != nil {
		return "", err
	}
	if err := c.rdb.HSet(ctx, key, c.values).Err(); err != nil {
		return "", err
	}
	return key, nil
}

func (c *BarCreate) genKey() (string, error) {
	id, ok := c.values["id"].(string)
	if !ok {
		return "", fmt.Errorf("id not set for %#v", c)
	}
	return prefixBar + id, nil
}
