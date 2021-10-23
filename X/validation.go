package x

func IsZeroValue(any interface{}) bool {
	if any == nil {
		return true
	}
	switch v := any.(type) {
	case string:
		return v == ``
	case int64:
		return v == 0
	case int32:
		return v == 0
	case int16:
		return v == 0
	case int8:
		return v == 0
	case int:
		return v == 0
	case float64:
		return v == 0
	case float32:
		return v == 0
	case bool:
		return !v
	case []byte:
		return v == nil
	case byte:
		return v == 0
	case uint64:
		return v == 0
	case uint32:
		return v == 0
	case uint16:
		return v == 0
	case uintptr:
		return v == 0
	default:
		return false
	}
}
