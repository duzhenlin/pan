// Package account
// Created by Duzhenlin
// @Author   Duzhenlin
// @Email: duzhenlin@vip.qq.com
// @Date: 2024/10/15
// @Time: 15:46

package account

import "github.com/duzhenlin/pan/conf"

type UserInfoResponse struct {
	BaiduName    string `json:"baidu_name"`
	NetdiskName  string `json:"netdisk_name"`
	AvatarUrl    string `json:"avatar_url"`
	VipType      int    `json:"vip_type"`
	Uk           int    `json:"uk"` //uk字段对应auth.UserInfo方法返回的user_id
	ErrorCode    int    `json:"errno"`
	ErrorMsg     string `json:"errmsg"`
	RequestID    int
	RequestIDStr string `json:"request_id"` //用户信息接口返回的request_id为string类型
}

type QuotaResponse struct {
	conf.CloudDiskResponseBase
	Total  int  `json:"total"`
	Used   int  `json:"used"`
	Free   int  `json:"free"`
	Expire bool `json:"expire"`
}

type Account struct {
	AccessToken string
}

const UserInfoUri = "/rest/2.0/xpan/nas?method=uinfo"
const QuotaUri = "/api/quota"
