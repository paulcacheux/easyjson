//go:build go1.18

package tests

type GenericsBaseTemplate[T any] struct {
	A T `json:"a"`
}

//easyjson:json
type GenericsSpecific struct {
	B GenericsBaseTemplate[int] `json:"b"`
}

var genericsValue GenericsSpecific = GenericsSpecific{
	B: GenericsBaseTemplate{
		A: 12,
	},
}

var genericsValueString = `{"b":{"a":12}}`
