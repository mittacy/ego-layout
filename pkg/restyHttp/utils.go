package restyHttp

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"time"
)

var ErrParseTime = errors.New("时间格式不正确")

// Decode 响应解析，map => struct
// 由于 mapstructure 解析时如果结构体存在时间将无法解析成功，所以此处添加一个钩子告诉mapstructure如何解析时间
// @param input map数据
// @param result struct结构体，需要指针
// @param timeFormat 解析时间格式
//	默认为 []string{"2006-01-02 15:04:05", "2006-01-02T15:04:05Z07:00", "2006-01-02T15:04:05.999999999Z07:00"}
// @return error
func Decode(input interface{}, result interface{}, timeFormat ...string) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			toTimeHookFunc(timeFormat...)),
		Result: result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(input.(map[string]interface{})); err != nil {
		return err
	}
	return err
}

func toTimeHookFunc(timeFormat ...string) mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return ParseTime(data, timeFormat...)
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
	}
}

// ParseTime 解析时间字符串为time.Time
// @param data 时间
// @param format 可能的时间格式
// @return time.Time
// @return error 如果格式错误，将返回 ErrTimeFormat
func ParseTime(data interface{}, format ...string) (time.Time, error) {
	formats := make([]string, 0)

	if len(format) > 0 {
		formats = append(formats, format...)
	} else {
		formats = []string{
			"2006-01-02 15:04:05",
			time.RFC3339,
			time.RFC3339Nano,
		}
	}

	for _, v := range formats {
		if result, err := time.Parse(v, data.(string)); err != nil {
			return result, nil
		}
	}

	return time.Now(), ErrParseTime

}
