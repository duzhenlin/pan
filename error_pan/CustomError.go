// Package error
// Created by Duzhenlin
// @Author   Duzhenlin
// @Email: duzhenlin@vip.qq.com
// @Date: 2024/10/15
// @Time: 14:47

package error_pan

import (
	"fmt"
)

type BaiduPanError struct {
	Code      int
	Message   string
	RequestId string
}

func (e *BaiduPanError) Error() string {
	if e.RequestId == "" {
		return fmt.Sprintf("[BaiduPanError] Code=%d, Message=%s", e.Code, e.Message)
	}
	return fmt.Sprintf("[BaiduPanError] Code=%d, Message=%s, RequestId=%s", e.Code, e.Message, e.RequestId)
}

func NewBaiduPanError(code int, message string, requestId string) error {
	return &BaiduPanError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}

func (e *BaiduPanError) GetCode() int {
	return e.Code
}

func (e *BaiduPanError) GetMessage() string {
	return e.Message
}

func (e *BaiduPanError) GetRequestId() string {
	return e.RequestId
}
