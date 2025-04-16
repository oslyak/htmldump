package htmldump

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// MapTypeName returns the name of map key, value types.
func MapTypeName(mapType reflect.Type) (string, error) {
	result := ``

	if mapType.Kind() == reflect.Pointer {
		mapType = mapType.Elem()
		result = `*`
	}

	if mapType.Kind() != reflect.Map {
		return ``, errors.New(`[mapTypeName] input parameter is not a map`)
	}

	result += fmt.Sprintf(`map[%s]%s`, mapType.Key().String(), mapType.Elem().String())

	return result, nil
}

// MapTypeName returns the name of map key, value types.
func SliceTypeName(sliceType reflect.Type) (string, error) {
	var result strings.Builder

	if sliceType.Kind() == reflect.Pointer {
		sliceType = sliceType.Elem()

		result.WriteString(`*`)
	}

	if sliceType.Kind() != reflect.Slice {
		return ``, errors.New(`[SliceTypeName] the input parameter is not a slice or pointer to slice`)
	}

	result.WriteString(`[]`)
	result.WriteString(sliceType.Elem().String())

	return result.String(), nil
}
