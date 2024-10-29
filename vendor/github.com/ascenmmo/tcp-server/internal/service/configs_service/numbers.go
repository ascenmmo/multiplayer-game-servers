package configsService

func incrementValue(value any) any {
	switch v := value.(type) {
	case int:
		return v + 1
	case int32:
		return v + 1
	case int64:
		return v + 1
	case float32:
		return v + 1
	case float64:
		return v + 1
	default:
		return value
	}
}

func decrementValue(value any) any {
	switch v := value.(type) {
	case int:
		return v - 1
	case int32:
		return v - 1
	case int64:
		return v - 1
	case float32:
		return v - 1
	case float64:
		return v - 1
	default:
		return value
	}
}

func additionValues(v1, v2 any, expectedType string) any {
	switch expectedType {
	case "int":
		return int(toFloat64(v1) + toFloat64(v2))
	case "int64":
		return int64(toFloat64(v1) + toFloat64(v2))
	case "float64":
		return toFloat64(v1) + toFloat64(v2)
	default:
		return nil // В случае неподдерживаемого типа
	}
}

func subtractValues(v1, v2 any, expectedType string) any {
	switch expectedType {
	case "int":
		return int(toFloat64(v1) - toFloat64(v2))
	case "int64":
		return int64(toFloat64(v1) + toFloat64(v2))
	case "float64":
		return toFloat64(v1) - toFloat64(v2)
	default:
		return nil // В случае неподдерживаемого типа
	}
}

func toFloat64(value any) float64 {
	switch v := value.(type) {
	case int:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		return 0
	}
}
