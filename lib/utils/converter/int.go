package converter

func ConvertPointerToInt32(value *int32) int32 {
	var result int32

	if value != nil {
		result = *value
	}

	return result
}

func ConvertInt32ToPointer(value int32) *int32 {
	var result *int32 = &value

	return result
}

func ConvertPointerToUint32(value *uint32) uint32 {
	var result uint32

	if value != nil {
		result = *value
	}

	return result
}

func ConvertUint32ToPointer(value uint32) *uint32 {
	var result *uint32 = &value

	return result
}
