package main

import (
	"fmt"
	"github.com/duzhenlin/pan/error_pan"
	"github.com/duzhenlin/pan/file"
)

func main() {
	accessToken := "121.0814add880b06b2b83aaf26bc9768a64.YBuxciRyVZ0o3JRsaX0g9xT1tGhoTth28q3ZhL8.tn1_EA"
	fileClient := file.NewFileClient(accessToken)
	option := fileClient.NewCategoryListOption()
	option.Desc = 1
	option.Start = 0
	option.Limit = 100
	option.Order = "time"
	option.Category = []string{
		file.ListCategoryTypeVideo,
		file.ListCategoryTypeAudio,
		file.ListCategoryTypeImage,
		//file.ListCategoryTypeDoc,
	}
	option.ParentPath = "/"
	option.ShowDir = 0
	res, err := fileClient.CategoryList(option)
	if err != nil {
		fmt.Println("err:", err)
		fmt.Println("err code:", err.(*error_pan.BaiduPanError).GetCode())
		fmt.Println("err msg:", err.(*error_pan.BaiduPanError).GetMessage())
		fmt.Println("err request_id:", err.(*error_pan.BaiduPanError).GetRequestId())
		return
	}
	fmt.Println(res)
}
