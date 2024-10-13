// Package file
// Created by Duzhenlin
// @Author   Duzhenlin
// @Email: duzhenlin@vip.qq.com
// @Date: 2024/10/13
// @Time: 14:12

package file

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

// Md5File md5_file()
func Md5File(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil)), nil
}
