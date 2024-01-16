package jsonschema

import (
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func NewProperties() *orderedmap.OrderedMap[string, *Type] {
	return orderedmap.New[string, *Type]()
}
