package ast

import (
	"errors"
	"fmt"
)

func toDouble(v interface{}) (float64, error) {
	switch t := v.(type) {
	case int8:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case int:
		return float64(t), nil
	case uint8:
		return float64(t), nil
	case uint16:
		return float64(t), nil
	case uint32:
		return float64(t), nil
	case uint64:
		return float64(t), nil
	case uint:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return t, nil
	default:
		return 0.0, errors.New(fmt.Sprintf("Could not convert %v to double", v))
	}
}

func isBool(v interface{}) bool {
	var _, ok = v.(bool)
	return ok
}

func isString(v interface{}) bool {
	var _, ok = v.(string)
	return ok
}
