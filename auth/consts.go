// Package auth
// Created by Duzhenlin
// @Author   Duzhenlin
// @Email: duzhenlin@vip.qq.com
// @Date: 2024/10/15
// @Time: 15:46

package auth

type Auth struct {
	ClientID     string
	ClientSecret string
}

type AccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	SessionKey       string `json:"session_key"`
	SessionSecret    string `json:"session_secret"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type RefreshTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	SessionKey       string `json:"session_key"`
	SessionSecret    string `json:"session_secret"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type UserInfoResponse struct {
	OpenID       string `json:"openid"`
	UnionID      string `json:"unionid"` // 百度用户统一标识，对当前开发者帐号唯一
	UserID       string `json:"userid"`  // 老版百度用户的唯一标识，后续不在返回该字段，user_id字段对应account.UserInfo方法返回的uk
	UserName     string `json:"username"`
	SecureMobile int    `json:"securemobile"` // 当前用户绑定手机号，需要向百度开放平台单独申请权限
	Portrait     string `json:"portrait"`
	UserDetail   string `json:"userdetail"`
	Birthday     string `json:"birthday"`
	Marriage     string `json:"marriage"`
	Sex          string `json:"sex"`
	Blood        string `json:"blood"`
	IsBindMobile string `json:"is_bind_mobile"`
	IsRealName   string `json:"is_realname"`
	ErrorCode    int    `json:"errno"`
	ErrorMsg     string `json:"errmsg"`
}

const OAuthUri = "/oauth/2.0/authorize"
const OAuthTokenUri = "/oauth/2.0/token"
const UserInfoUri = "/rest/2.0/passport/users/getInfo"
