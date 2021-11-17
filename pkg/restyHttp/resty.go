package restyHttp

import (
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"time"
)

// Get GET请求，返回数据为map结构
// @param host 域名，example: https://www.baidu.com
// @param uri example: /user
// @param timeout 超时控制  example: time.Second*5
// @return map[string]interface{}
// @return error
func Get(host, uri string, timeout time.Duration) (map[string]interface{}, error) {
	url := fullUrl(host, uri)

	client := resty.New().SetTimeout(timeout)
	res := Reply{}

	resp, err := client.R().SetResult(&res).ForceContentType("application/json").Get(url)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !resp.IsSuccess() {
		err = errors.New(resp.String())
		return nil, err
	}
	if res.Code != 0 {
		return nil, errors.New(res.Msg)
	}

	if v, ok := res.Data.(map[string]interface{}); ok {
		return v, nil
	}

	return map[string]interface{}{}, nil
}

// GetParams GET请求，返回数据为map结构
// @param host 域名，example: https://www.baidu.com
// @param uri example: /user
// @param params 请求参数
// @param timeout 超时控制  example: time.Second*5
// @return map[string]interface{}
// @return error
func GetParams(host, uri string, params map[string]string, timeout time.Duration) (map[string]interface{}, error) {
	url := fullUrl(host, uri)

	client := resty.New().SetTimeout(timeout)
	res := Reply{}

	resp, err := client.R().SetQueryParams(params).SetResult(&res).ForceContentType("application/json").Get(url)

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		err = errors.New(resp.String())
		return nil, err
	}
	if res.Code != 0 {
		return nil, errors.New(res.Msg)
	}

	if v, ok := res.Data.(map[string]interface{}); ok {
		return v, nil
	}

	return map[string]interface{}{}, nil
}

// Post POST请求
// @param host 域名，example: https://www.baidu.com
// @param uri example: /user
// @param body 请求体数据，struct/map/[]byte/……
// @param timeout 超时控制  example: time.Second*5
// @return map[string]interface{} 响应数据
// @return error
func Post(host, uri string, body interface{}, timeout time.Duration) (map[string]interface{}, error) {
	url := fullUrl(host, uri)

	client := resty.New().SetTimeout(timeout)
	res := Reply{}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&res).
		Post(url)

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		err = errors.New(resp.String())
		return nil, err
	}
	if res.Code != 0 {
		return nil, errors.New(res.Msg)
	}

	if v, ok := res.Data.(map[string]interface{}); ok {
		return v, nil
	}

	return map[string]interface{}{}, nil
}

func fullUrl(host, uri string) string {
	return host + uri
}
