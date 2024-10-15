// Package auth
//
// 百度授权相关，wiki地址 https://openauth.baidu.com/doc/doc.html
package auth

import (
	"encoding/json"
	"fmt"
	"github.com/duzhenlin/pan/conf"
	"github.com/duzhenlin/pan/error_pan"
	"github.com/duzhenlin/pan/utils/httpclient"
	"log"
	"net/url"
)

func NewAuthClient(clientID string, clientSecret string) *Auth {
	return &Auth{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

// OAuthUrl 获取授权页网址
func (a *Auth) OAuthUrl(redirectUri string) string {
	oAuthUrl := ""

	v := url.Values{}
	v.Add("response_type", "code")
	v.Add("client_id", a.ClientID)
	v.Add("redirect_uri", redirectUri)
	v.Add("scope", "basic,netdisk")
	v.Add("state", "STATE")
	query := v.Encode()

	oAuthUrl = conf.BaiduOpenApiDomain + OAuthUri + "?" + query

	return oAuthUrl
}

// AccessToken 获取AccessToken
func (a *Auth) AccessToken(code, redirectUri string) (AccessTokenResponse, error) {
	ret := AccessTokenResponse{}

	v := url.Values{}
	v.Add("grant_type", "authorization_code")
	v.Add("code", code)
	v.Add("client_id", a.ClientID)
	v.Add("client_secret", a.ClientSecret)
	v.Add("redirect_uri", redirectUri)
	query := v.Encode()

	requestUrl := conf.BaiduOpenApiDomain + OAuthTokenUri + "?" + query

	resp, err := httpclient.Get(requestUrl, map[string]string{})
	if err != nil {
		log.Println("httpclient.Get failed, err:", err)
		return ret, err
	}

	if resp.StatusCode != 200 {
		return ret, error_pan.NewBaiduPanError(
			resp.StatusCode,
			fmt.Sprintf("HttpStatusCode is not equal to 200, httpStatusCode[%d], respBody[%s]", resp.StatusCode, string(resp.Body)),
			"",
		)
	}

	if err := json.Unmarshal(resp.Body, &ret); err != nil {
		return ret, err
	}

	if ret.Error != "" { //有错误

		return ret, error_pan.NewBaiduPanError(
			resp.StatusCode,
			fmt.Sprintf("error_code:%s, error_msg:%s", ret.Error, ret.ErrorDescription),
			"",
		)
	}

	return ret, nil
}

// RefreshToken 刷新AccessToken
func (a *Auth) RefreshToken(refreshToken string) (RefreshTokenResponse, error) {
	ret := RefreshTokenResponse{}

	v := url.Values{}
	v.Add("grant_type", "refresh_token")
	v.Add("refresh_token", refreshToken)
	v.Add("client_id", a.ClientID)
	v.Add("client_secret", a.ClientSecret)
	query := v.Encode()

	requestUrl := conf.BaiduOpenApiDomain + OAuthTokenUri + "?" + query

	resp, err := httpclient.Get(requestUrl, map[string]string{})
	if err != nil {
		log.Println("httpclient.Get failed, err:", err)
		return ret, err
	}

	if resp.StatusCode != 200 {
		return ret, error_pan.NewBaiduPanError(
			resp.StatusCode,
			fmt.Sprintf("HttpStatusCode is not equal to 200, httpStatusCode[%d], respBody[%s]", resp.StatusCode, string(resp.Body)),
			"",
		)
	}

	if err := json.Unmarshal(resp.Body, &ret); err != nil {
		return ret, err
	}

	if ret.Error != "" { //有错误
		return ret, error_pan.NewBaiduPanError(
			resp.StatusCode,
			fmt.Sprintf("Error is not empty, httpStatusCode[%d], ErrorDescription[%s]", resp.StatusCode, ret.ErrorDescription),
			"",
		)
	}

	return ret, nil
}

// UserInfo
// 获取授权用户的百度账号信息，可以通过unionid字段来识别多个百度产品授权的是否是同一用户
// 注：获取网盘账号信息请使用account.UserInfo方法
func (a *Auth) UserInfo(accessToken string) (UserInfoResponse, error) {
	ret := UserInfoResponse{}

	v := url.Values{}
	v.Add("access_token", accessToken)
	v.Add("get_unionid", "1") //需要获取unionid时，传递get_unionid = 1
	query := v.Encode()

	requestUrl := conf.BaiduOpenApiDomain + UserInfoUri + "?" + query

	resp, err := httpclient.Get(requestUrl, map[string]string{})
	if err != nil {
		log.Println("httpclient.Get failed, err:", err)
		return ret, err
	}

	if resp.StatusCode != 200 {
		return ret, error_pan.NewBaiduPanError(
			resp.StatusCode,
			fmt.Sprintf("HttpStatusCode is not equal to 200, httpStatusCode[%d], respBody[%s]", resp.StatusCode, string(resp.Body)),
			"",
		)
	}

	if err := json.Unmarshal(resp.Body, &ret); err != nil {
		return ret, err
	}

	if ret.ErrorCode != 0 { //有错误
		return ret, error_pan.NewBaiduPanError(
			ret.ErrorCode,
			fmt.Sprintf("ErrorCode is 0, ErrorCode[%d], ErrorMsg[%s]", ret.ErrorCode, ret.ErrorMsg),
			"",
		)
	}

	return ret, nil
}
