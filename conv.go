package conv

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func ToFloat64(v interface{}) float64 {
	if v == nil {
		return 0
	}
	str := ""
	switch v.(type) {
	case []byte:
		str = string(v.([]byte))
	default:
		str = fmt.Sprint(v)
	}
	ret, _ := strconv.ParseFloat(str, 64)
	return ret
}

func ToFloat32(v interface{}) float32 {
	return float32(ToFloat64(v))
}

func ToInt(v interface{}) int {
	return int(ToFloat64(v))
}

func ToInt32(v interface{}) int32 {
	return int32(ToInt(v))
}

func ToInt64(v interface{}) int64 {
	return int64(ToFloat64(v))
}

func ToStr(v interface{}) string {
	if v == nil {
		return ""
	}

	return fmt.Sprint(v)
}

func ToJsonByte(v interface{}) []byte {
	data, _ := json.Marshal(v)
	return data
}

func ToJsonString(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func JsonToObject(data interface{}, obj interface{}) error {
	var d []byte
	switch md := data.(type) {
	case string:
		d = []byte(md)
	case []byte:
		d = md
	default:
		return fmt.Errorf("unkonw type")
	}

	return json.Unmarshal(d, obj)
}
