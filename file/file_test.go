package file

import (
	"github.com/duzhenlin/pan/conf"
	"testing"
)

func TestFile_List(t *testing.T) {
	fileClient := NewFileClient(conf.TestData.AccessToken)
	option := fileClient.NewDefaultListOption()
	option.Dir = conf.TestData.Dir
	option.Start = 0
	option.Limit = 100
	res, err := fileClient.List(option)
	if err != nil {
		t.Errorf("TestList failed, err:%v", err)
	}
	t.Logf("TestList res: %+v", res)
}

func TestFile_Metas(t *testing.T) {
	fileClient := NewFileClient(conf.TestData.AccessToken)
	option := fileClient.NewMetasOption()
	option.FsIDs = []uint64{conf.TestData.FsID}
	res, err := fileClient.Metas(option)
	if err != nil {
		t.Errorf("TestMetas failed, err:%v", err)
	}
	t.Logf("TestMetas res: %+v", res)
}

func TestFile_Streaming(t *testing.T) {
	fileClient := NewFileClient(conf.TestData.AccessToken)
	res, err := fileClient.Streaming(conf.TestData.Path, conf.TestData.TranscodingType)
	if err != nil {
		t.Errorf("TestFile_Streaming failed, err:%v", err)
	}
	t.Logf("TestFile_Streaming res: %+v", res)
}
