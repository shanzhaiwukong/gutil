package gutil

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// CStr2Int 字符串转数字（int）
func CStr2Int(value string, defaultValue int) int {
	if res, err := strconv.ParseInt(value, 0, 0); err != nil {
		return defaultValue
	} else {
		return int(res)
	}
}

// CStr2Int8 字符串转数字（int8）
func CStr2Int8(value string, defaultValue int8) int8 {
	if res, err := strconv.ParseInt(value, 0, 8); err != nil {
		return defaultValue
	} else {
		return int8(res)
	}
}

// CStr2Int16 字符串转数字（int16）
func CStr2Int16(value string, defaultValue int16) int16 {
	if res, err := strconv.ParseInt(value, 0, 16); err != nil {
		return defaultValue
	} else {
		return int16(res)
	}
}

// CStr2Int32 字符串转数字（int32）
func CStr2Int32(value string, defaultValue int32) int32 {
	if res, err := strconv.ParseInt(value, 0, 32); err != nil {
		return defaultValue
	} else {
		return int32(res)
	}
}

// CStr2Int64 字符串转数字（int64）
func CStr2Int64(value string, defaultValue int64) int64 {
	if res, err := strconv.ParseInt(value, 0, 64); err != nil {
		return defaultValue
	} else {
		return int64(res)
	}
}

// CStr2Float32 字符串转小数
func CStr2Float32(value string, defaultValue float32) float32 {
	if res, err := strconv.ParseFloat(value, 32); err != nil {
		return defaultValue
	} else {
		return float32(res)
	}
}

// CStr2Float64 字符串转小数
func CStr2Float64(value string, defaultValue float64) float64 {
	if res, err := strconv.ParseFloat(value, 64); err != nil {
		return defaultValue
	} else {
		return res
	}
}

// CStr2Bool 字符串转Boolean
func CStr2Bool(value string) bool {
	if value != "" && (value == "1" || strings.ToLower(value) == "true") {
		return true
	}
	return false
}

// CInt2Str 数字转字符串
func CInt2Str(value interface{}) string {
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", value)
	default:
		return ""
	}
}

// CFloat2Str 小数转字符串
func CFloat2Str(value interface{}, length int) string {
	switch value.(type) {
	case float32, float64:
		if length == 0 {
			return fmt.Sprintf("%f", value)
		}
		return fmt.Sprint("%."+CInt2Str(length)+"f", value)
	default:
		return ""
	}
}

// CBool2Str Boolean转字符串
func CBool2Str(value bool) string {
	if value {
		return "true"
	}
	return "false"
}

// CObject2JsonStr 对象转Json字符串
func CObject2JsonStr(obj interface{}) string {
	bs, _ := json.Marshal(obj)
	return string(bs)
}

// CStr2Time 字符串转时间
func CStr2Time(value, layout string, defaultValue time.Time) time.Time {
	if res, err := time.Parse(layout, value); err != nil {
		return defaultValue
	} else {
		return res
	}
}

// CInt2Time 时间戳转时间 支持10位时间戳和13位时间戳
func CInt2Time(_time int64, defaultValue time.Time) time.Time {
	var res time.Time
	if len(CInt2Str(_time)) == 10 {
		res = time.Unix(_time, 0)
	} else if len(CInt2Str(_time)) == 13 {
		res = time.Unix(_time/1e3, _time%1e3*1e6)
	} else {
		return defaultValue
	}
	return res
}
