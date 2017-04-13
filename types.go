package pickle

import (
	"github.com/lomik/og-rek"
)

func ToList(value interface{}) ([]interface{}, bool) {
	if l, ok := value.([]interface{}); ok {
		return l, ok
	}

	if l, ok := value.(ogórek.Tuple); ok {
		return []interface{}(l), ok
	}

	return nil, false
}

func ToInt64(value interface{}) (int64, bool) {
	switch value := value.(type) {
	case float32:
		return int64(value), true
	case float64:
		return int64(value), true
	case int:
		return int64(value), true
	case int16:
		return int64(value), true
	case int32:
		return int64(value), true
	case int64:
		return int64(value), true
	case int8:
		return int64(value), true
	case uint:
		return int64(value), true
	case uint16:
		return int64(value), true
	case uint32:
		return int64(value), true
	case uint64:
		return int64(value), true
	case uint8:
		return int64(value), true
	default:
		return 0, false
	}
}

func ToFloat64(value interface{}) (float64, bool) {
	switch value := value.(type) {
	case float32:
		return float64(value), true
	case float64:
		return float64(value), true
	case int:
		return float64(value), true
	case int16:
		return float64(value), true
	case int32:
		return float64(value), true
	case int64:
		return float64(value), true
	case int8:
		return float64(value), true
	case uint:
		return float64(value), true
	case uint16:
		return float64(value), true
	case uint32:
		return float64(value), true
	case uint64:
		return float64(value), true
	case uint8:
		return float64(value), true
	default:
		return 0, false
	}
}
