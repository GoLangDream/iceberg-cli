package project

import (
	"github.com/GoLangDream/rgo/rstring"
	"github.com/flosch/pongo2/v5"
)

func initPongo2() {
	pongo2.RegisterFilter("underscore", underscore)
	pongo2.RegisterFilter("camelize", camelize)
}

func underscore(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	if !in.IsString() {
		return nil, &pongo2.Error{
			Sender: "only strings should be sent to the scream filter",
		}
	}

	s := in.String()
	s = rstring.Underscore(s)

	return pongo2.AsValue(s), nil
}

func camelize(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	if !in.IsString() {
		return nil, &pongo2.Error{
			Sender: "only strings should be sent to the scream filter",
		}
	}

	s := in.String()
	s = rstring.Camelize(s)

	return pongo2.AsValue(s), nil
}
