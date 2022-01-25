package apierr

// 服务错误定义
var (
	Success           = &BizErr{0, "成功"}
	Param             = &BizErr{1, "参数错误"}
	DebounceIntercept = &BizErr{2, "防抖拦截，此次访问无效"}
	RestyHttp         = &BizErr{3, "请求外部服务错误"}
	Unauthorized      = &BizErr{401, "未认证"}
	Forbidden         = &BizErr{401, "权限不足"}
	Unknown           = &BizErr{500, "未知错误"}
)

// 业务错误定义格式: 大模块:中间模块:业务模块
var (
	//用户相关 code: 1000XX
	LoginExpire = &BizErr{100001, "登录信息过期"}
	UserNoExist = &BizErr{100002, "用户不存在"}
)
