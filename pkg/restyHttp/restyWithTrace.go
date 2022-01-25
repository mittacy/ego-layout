package restyHttp

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"time"
)

// GetWithTrace GET请求，返回数据为map结构
// @param host 域名，example: https://www.baidu.com
// @param uri example: /user
// @param timeout 超时控制  example: time.Second*5
// @return map[string]interface{}
// @return int 返回的业务状态码
// @return error
func GetWithTrace(c *gin.Context, host, uri string, timeout time.Duration) (map[string]interface{}, int, error) {
	url := fullUrl(host, uri)

	client := resty.New().SetTimeout(timeout)
	res := Reply{}

	resp, err := client.R().SetResult(&res).ForceContentType("application/json").Get(url)
	if err != nil {
		logger.ErrorwWithTrace(c, host+uri, "res", resp, "err", err)
		return nil, 0, errors.WithStack(err)
	} else {
		logger.InfowWithTrace(c, host+uri, "res", resp)
	}

	if !resp.IsSuccess() {
		return nil, 0, errors.New(resp.String())
	}

	if res.Code != 0 {
		return nil, res.Code, errors.New(res.Msg)
	}

	return resPackage(res), res.Code, nil
}

// GetParamsWithTrace GET请求，返回数据为map结构
// @param host 域名，example: https://www.baidu.com
// @param uri example: /user
// @param params 请求参数
// @param timeout 超时控制  example: time.Second*5
// @return map[string]interface{}
// @return int 返回的业务状态码
// @return error
func GetParamsWithTrace(c *gin.Context, host, uri string, params map[string]string, timeout time.Duration) (map[string]interface{}, int, error) {
	url := fullUrl(host, uri)

	client := resty.New().SetTimeout(timeout)
	res := Reply{}

	resp, err := client.R().SetQueryParams(params).SetResult(&res).ForceContentType("application/json").Get(url)
	if err != nil {
		logger.ErrorwWithTrace(c, host+uri, "res", resp, "err", err)
		return nil, 0, errors.WithStack(err)
	} else {
		logger.InfowWithTrace(c, host+uri, "res", resp)
	}

	if !resp.IsSuccess() {
		return nil, 0, errors.New(resp.String())
	}

	if res.Code != 0 {
		return nil, res.Code, errors.New(res.Msg)
	}

	return resPackage(res), res.Code, nil
}

// PostWithTrace POST请求
// @param host 域名，example: https://www.baidu.com
// @param uri example: /user
// @param body 请求体数据，struct/map/[]byte/……
// @param timeout 超时控制  example: time.Second*5
// @return map[string]interface{} 响应数据
// @return int 返回的业务状态码
// @return error
func PostWithTrace(c *gin.Context, host, uri string, body interface{}, timeout time.Duration) (map[string]interface{}, int, error) {
	url := fullUrl(host, uri)

	client := resty.New().SetTimeout(timeout)
	res := Reply{}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&res).
		Post(url)
	if err != nil {
		logger.ErrorwWithTrace(c, host+uri, "res", resp, "err", err)
		return nil, 0, errors.WithStack(err)
	} else {
		logger.InfowWithTrace(c, host+uri, "res", resp)
	}

	if !resp.IsSuccess() {
		return nil, 0, errors.New(resp.String())
	}

	if res.Code != 0 {
		return nil, res.Code, errors.New(res.Msg)
	}

	return resPackage(res), res.Code, nil
}
