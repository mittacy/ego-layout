package apierr

import (
	"github.com/mittacy/ego/library/apierr"
)

// 业务错误定义格式: 大模块:中间模块:业务模块
var (
	// 共用 code: 1000XX
	ResourceNoExist = &apierr.BizErr{Code: 100001, Msg: "资源不存在"}

	//用户相关 code: 2000XX
	LoginExpire = &apierr.BizErr{Code: 200001, Msg: "登录信息过期"}
	UserNoExist = &apierr.BizErr{Code: 200002, Msg: "用户不存在"}
)
