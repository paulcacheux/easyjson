//go:build !go1.18
// +build !go1.18

package tests

//easyjson:json
type GenericsSpecific struct {
}

var genericsValue GenericsSpecific = GenericsSpecific{}

var genericsValueString = `{}`
