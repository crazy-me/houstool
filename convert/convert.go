package convert

import "fmt"

// String Convert
func ToString(value interface{}) string {
	point := []string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64"}
	if isValueInList(fmt.Sprintf("%T", value), point) {
		return fmt.Sprintf("%d", value)
	}

	if fmt.Sprintf("%T", value) == "float32" || fmt.Sprintf("%T", value) == "float64" {
		return fmt.Sprintf("%.2f", value)
	}
	return ""
}

// Check Slice Element
func isValueInList(value string, s []string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}
