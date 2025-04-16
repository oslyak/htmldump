package htmldump

import (
	"fmt"
	"reflect"
)

// mapToHTML dumps a map to HTML table.
func mapToHTML(doc *htmlDocument, reflectedMap reflect.Value) error {
	if !isMapOrPointerToMap(reflectedMap) {
		return fmt.Errorf(`[mapToHTML] only accepts map or pointers to map, got %s`, reflectedMap.Kind())
	}

	caption, err := mapCaption(reflectedMap)
	if err != nil {
		return err
	}

	table := new(tableT)
	table.Caption(caption).
		mapHeader(reflectedMap).
		mapBody(reflectedMap).
		toHTML(doc)

	return nil
}

// Returns the caption of HTML table for a map key and value types.
func mapCaption(reflectedMap reflect.Value) (string, error) {
	typeName, err := MapTypeName(reflectedMap.Type())
	if err != nil {
		return ``, err
	}

	return fmt.Sprintf(`%s (length: %d)`, typeName, reflectedMap.Len()), nil
}

func (table *tableT) mapHeader(reflectedMap reflect.Value) *tableT {
	var captions, types rowT

	captions.addCell(cellT{value: `map key`, key: true})
	types.addCell(cellT{value: reflectedMap.Type().Key().Name(), key: true})

	table.headerRow(reflectedMap.Type().Elem(), &captions, &types)

	return table
}

// Generate table body for a map.
func (table *tableT) mapBody(reflectedMap reflect.Value) *tableT {
	for _, key := range reflectedMap.MapKeys() {
		var row rowT

		row.addCell(cellT{value: formatValue(key), key: true})

		item := reflectedMap.MapIndex(key)
		if item.Kind() == reflect.Pointer {
			item = item.Elem()
		}

		switch {
		case !item.IsValid():
			row.addCell(cellT{value: NULL, colspan: table.columns - 1})
		case item.Kind() == reflect.Struct:
			structRow(&row, item)
		default:
			row.addCellStr(formatValue(item))
		}

		table.addBodyRow(row)
	}

	return table
}

func isMapOrPointerToMap(reflectedMap reflect.Value) bool {
	if reflectedMap.Kind() == reflect.Pointer {
		reflectedMap = reflectedMap.Elem()
	}

	return reflectedMap.Kind() == reflect.Map
}
