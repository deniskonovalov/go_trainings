package validation

func IsValidNumber(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return true
	default:
		return false
	}
}

func IsValidMap(v any) bool {
	switch v.(type) {
	case map[string]int,
		map[string]int32,
		map[string]int64,
		map[string]float64,
		map[string]string,
		map[string]bool,
		map[string]any,
		map[string]byte,
		map[string]error:
		return true
	default:
		return false
	}
}

func IsValidSlice(v any) bool {
	switch v.(type) {
	case []int,
		[]int64,
		[]float64,
		[]string,
		[]bool,
		[]byte,
		[]rune,
		[]any:
		return true
	default:
		return false
	}
}
