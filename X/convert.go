package X

func Dereference(any interface{}) interface{} {
	if any == nil {
		return any
	}
	switch v := any.(type) {
	case *string:
		return *v
	case *int64:
		return *v
	case *int32:
		return *v
	case *int16:
		return *v
	case *int8:
		return *v
	case *int:
		return *v
	case *float64:
		return *v
	case *float32:
		return *v
	case *bool:
		return *v
	case *[]byte:
		return *v
	case *byte:
		return *v
	case *uint64:
		return *v
	case *uint32:
		return *v
	case *uint16:
		return *v
	case *uintptr:
		return *v
	case *interface{}:
		return *v
	default:
		return any
	}
}
