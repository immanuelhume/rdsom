package rdsom

import "time"

const prefixBar = "rdsom:rdsomgolden:Bar:"
const idxBar = "rdsom:rdsomgolden:idx:Bar:1642524233"

type Bar struct {
	BoolField    bool      `json:"boolField,omitempty"`
	FloatField   float64   `json:"floatField,omitempty"`
	FloatsField  []float64 `json:"floatsField,omitempty"`
	IntField     int       `json:"intField,omitempty"`
	IntsField    []int     `json:"intsField,omitempty"`
	StringField  string    `json:"stringField,omitempty"`
	StringsField []string  `json:"stringsField,omitempty"`
	TimeField    time.Time `json:"timeField,omitempty"`
}
