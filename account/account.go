package account

import (
	"encoding/json"
	"fmt"
	"github.com/duzhenlin/pan/conf"
	"github.com/duzhenlin/pan/error_pan"
	"github.com/duzhenlin/pan/utils/httpclient"
	"log"
	"net/url"
	"strconv"
)

func NewAccountClient(accessToken string) *Account {
	return &Account{
		AccessToken: accessToken,
	}
}

// UserInfo 获取网盘用户信息
func (a *Account) UserInfo() (UserInfoResponse, error) {
	ret := UserInfoResponse{}

	v := url.Values{}
	v.Add("access_token", a.AccessToken)
	query := v.Encode()

	requestUrl := conf.OpenApiDomain + UserInfoUri + "&" + query
	resp, err := httpclient.Get(requestUrl, map[string]string{})
	if err != nil {
		log.Println("httpclient.Get failed, err:", err)
		return ret, err
	}

	if resp.StatusCode != 200 {
		return ret, error_pan.NewBaiduPanError(
			ret.ErrorCode,
			fmt.Sprintf("HttpStatusCode is not equal to 200, httpStatusCode[%d], respBody[%s]", resp.StatusCode, string(resp.Body)),
			ret.RequestIDStr,
		)
	}

	if err := json.Unmarshal(resp.Body, &ret); err != nil {
		return ret, err
	}

	if ret.ErrorCode != 0 { //错误码不为0
		return ret, error_pan.NewBaiduPanError(
			ret.ErrorCode,
			fmt.Sprintf("error_code:%d, error_msg:%s", ret.ErrorCode, ret.ErrorMsg),
			ret.RequestIDStr,
		)
	}

	//兼容用户信息接口返回的request_id为string类型的问题
	ret.RequestID, err = strconv.Atoi(ret.RequestIDStr)
	if err != nil {
		log.Println("strconv.Atoi failed, err:", err)
		return ret, err
	}

	return ret, nil
}

// Quota 获取用户网盘容量信息
func (a *Account) Quota() (QuotaResponse, error) {
	ret := QuotaResponse{}

	v := url.Values{}
	v.Add("access_token", a.AccessToken)
	v.Add("checkfree", "1")
	v.Add("checkexpire", "1")
	query := v.Encode()

	requestUrl := conf.OpenApiDomain + QuotaUri + "?" + query
	resp, err := httpclient.Get(requestUrl, map[string]string{})
	if err != nil {
		log.Println("httpclient.Get failed, err:", err)
		return ret, err
	}

	if resp.StatusCode != 200 {
		return ret, error_pan.NewBaiduPanError(
			ret.ErrorCode,
			fmt.Sprintf("HttpStatusCode is not equal to 200, httpStatusCode[%d], respBody[%s]", resp.StatusCode, string(resp.Body)),
			strconv.FormatUint(ret.RequestID, 10),
		)
	}

	if err := json.Unmarshal(resp.Body, &ret); err != nil {
		return ret, err
	}

	if ret.ErrorCode != 0 { //错误码不为0
		return ret, error_pan.NewBaiduPanError(
			ret.ErrorCode,
			fmt.Sprintf("error_code:%d, error_msg:%s", ret.ErrorCode, ret.ErrorMsg),
			strconv.FormatUint(ret.RequestID, 10),
		)
	}

	return ret, nil
}
