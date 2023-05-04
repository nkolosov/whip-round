package conv

import (
	"errors"
	"strconv"
)

func ConvertInterfaceToFloat64(i interface{}) (float64, error) {
	switch v := i.(type) {
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, errors.New("unknown type")
	}
}
