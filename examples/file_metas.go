package main

import (
	"fmt"
	"github.com/duzhenlin/pan/file"
)

func main() {
	accessToken := "121.0814add880b06b2b83aaf26bc9768a64.YBuxciRyVZ0o3JRsaX0g9xT1tGhoTth28q3ZhL8.tn1_EA"
	fileClient := file.NewFileClient(accessToken)
	fsIDs := []uint64{760976149817903}
	option := fileClient.NewMetasOption()
	option.FsIDs = fsIDs
	res, err := fileClient.Metas(option)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(res)
}
