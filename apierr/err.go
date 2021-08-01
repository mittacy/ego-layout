package apierr

import (
	"errors"
)

var (
	ErrParam       = errors.New("参数错误")
	ErrCopier      = errors.New("结构体转化错误")
	ErrJsonMarshal = errors.New("json序列化错误")

	// 注册登录
	ErrLoginExpire = errors.New("登录信息过期")
	ErrNoLogin     = errors.New("未登录")

	// 缓存
	ErrCacheNoExist = errors.New("查询的缓存不存在")
)

var errCode = map[error]int{
	ErrParam:  CodeParamErr,
	ErrCopier: CodeBackErr,
	ErrJsonMarshal: CodeJsonMarshalErr,

}

func ErrCode(err error) int {
	if v, ok := errCode[err]; ok {
		return v
	}
	return CodeParamErr
}
