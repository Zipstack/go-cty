package gocty

import (
	"math/big"
	"reflect"
	"strings"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/set"
)

var valueType = reflect.TypeOf(cty.Value{})
var typeType = reflect.TypeOf(cty.Type{})

var setType = reflect.TypeOf(set.Set[interface{}]{})

var bigFloatType = reflect.TypeOf(big.Float{})
var bigIntType = reflect.TypeOf(big.Int{})

var emptyInterfaceType = reflect.TypeOf(interface{}(nil))

var stringType = reflect.TypeOf("")

// structTagIndices interrogates the fields of the given type (which must
// be a struct type, or we'll panic) and returns two maps, both from the
// pctsdk attribute names declared via struct tags:
// 1. to the indices of the fields holding those tags
// 2. to the detected tag values of the fields holding those tags
//
// This function will panic if two fields within the struct are tagged with
// the same pctsdk attribute name.
func structTagIndices(st reflect.Type) (map[string]int, map[string]map[string]bool) {
	ct := st.NumField()
	ret := make(map[string]int, ct)
	retTags := make(map[string]map[string]bool, ct)

	for i := 0; i < ct; i++ {
		field := st.Field(i)
		attrTags := field.Tag.Get("pctsdk")

		if attrTags != "" {
			parts := strings.Split(attrTags, ",")

			attrName := parts[0]

			ret[attrName] = i
			retTags[attrName] = map[string]bool{}

			for _, t := range parts[1:] {
				if t == "omitempty" {
					retTags[attrName]["omitempty"] = true
				}
			}
		}
	}

	return ret, retTags
}
