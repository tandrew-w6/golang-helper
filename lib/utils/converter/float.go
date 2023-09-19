package converter

func ConvertPointerToFloat32(value *float32) float32 {
	var result float32

	if value != nil {
		result = *value
	}

	return result
}

func ConvertFloat32ToPointer(value float32) *float32 {
	var result *float32 = &value

	return result
}

func ConvertPointerToFloat64(value *float64) float64 {
	var result float64

	if value != nil {
		result = *value
	}

	return result
}

func ConvertFloat64ToPointer(value float64) *float64 {
	var result *float64 = &value

	return result
}
