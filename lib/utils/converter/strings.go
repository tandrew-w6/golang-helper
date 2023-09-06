package converter

import "time"

func ConvertPointerToString(value *string) string {
	result := ""

	if value != nil {
		result = *value
	}

	return result
}

func ConvertStringToPointer(value string) *string {
	var result *string

	if value != "" {
		result = &value
	}

	return result
}

func ConvertStringToPointerTime(value string, layout string) *time.Time {
	var result *time.Time

	if layout == "" {
		layout = time.RFC3339
	}

	if value != "" {
		t, _ := time.Parse(layout, value)
		result = &t
	}

	return result
}
