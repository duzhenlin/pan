package main

import (
	"fmt"
	"github.com/duzhenlin/pan/file"
)

func main() {
	accessToken := "122.b0a9ab31cc24b429d460cd3ce1f1af97.Yn53jGAwd_1elGgODFvYl1sp9qOYVUDRiVawin5.tbNcEw"
	fileClient := file.NewFileClient(accessToken)
	// 参数配置
	option := fileClient.NewCreateDirOption()
	option.Path = "/apps/书"
	option.RType = "0"
	res, err := fileClient.CreateDir(option)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(res)
}
