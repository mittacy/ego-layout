package apierr

import (
	"github.com/mittacy/ego/library/apierr"
)

// 业务错误定义格式: 大模块:中间模块:业务模块
var (
	//用户相关 code: 1000XX
	LoginExpire = &apierr.BizErr{Code: 100001, Msg: "登录信息过期"}
	UserNoExist = &apierr.BizErr{Code: 100002, Msg: "用户不存在"}
)
