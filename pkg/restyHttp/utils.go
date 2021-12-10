package restyHttp

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"	// 解析 time.Time 格式

// Decode 响应解析，map => struct
// 由于 mapstructure 解析时如果结构体存在时间将无法解析成功，所以此处添加一个钩子告诉mapstructure如何解析时间
// @param input map数据
// @param result struct结构体，需要指针
// @param timeFormat 解析时间格式，默认为2006-01-02 15:04:05
// @return error
func Decode(input interface{}, result interface{}, timeFormat *string) error {
	format := TimeFormat
	if timeFormat != nil {
		format = *timeFormat
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			toTimeHookFunc(format)),
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

func toTimeHookFunc(timeFormat string) mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(timeFormat, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
	}
}
