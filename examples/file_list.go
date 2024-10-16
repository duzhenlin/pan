package main

import (
	"fmt"
	"github.com/duzhenlin/pan/error_pan"
	"github.com/duzhenlin/pan/file"
)

func main() {
	accessToken := "122.b0a9ab31cc24b429d460cd3ce1f1af97.Yn53jGAwd_1elGgODFvYl1sp9qOYVUDRiVawin5.tbNcEw"
	fileClient := file.NewFileClient(accessToken)
	option := fileClient.NewDefaultListOption()
	option.Desc = 1
	option.Order = "time"
	option.Dir = "/"
	res, err := fileClient.List(option)
	if err != nil {
		fmt.Println("err:", err)
		fmt.Println("err code:", err.(*error_pan.BaiduPanError).GetCode())
		fmt.Println("err msg:", err.(*error_pan.BaiduPanError).GetMessage())
		fmt.Println("err request_id:", err.(*error_pan.BaiduPanError).GetRequestId())
		return
	}
	fmt.Println(res)
}
