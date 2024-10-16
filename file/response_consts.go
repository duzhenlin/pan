// Package file
// Created by Duzhenlin
// @Author   Duzhenlin
// @Email: duzhenlin@vip.qq.com
// @Date: 2024/10/15
// @Time: 15:49

package file

import "github.com/duzhenlin/pan/conf"

type CategoryListResponse struct {
	conf.CloudDiskResponseBase
	List []struct {
		Category       int               `json:"category"`
		FsID           uint64            `json:"fs_id"`
		IsDir          int               `json:"isdir"`
		LocalCtime     int               `json:"local_ctime"`
		LocalMtime     int               `json:"local_mtime"`
		Md5            string            `json:"md5"`
		Path           string            `json:"path"`
		ServerCtime    int               `json:"server_ctime"`
		ServerFilename string            `json:"server_filename"`
		ServerMtime    int               `json:"server_mtime"`
		Size           int               `json:"size"`
		Thumbs         map[string]string `json:"thumbs"`
	} `json:"list"`
	RequestId string `json:"request_id"`
}

type ListResponse struct {
	conf.CloudDiskResponseBase
	List []struct {
		FsID           uint64            `json:"fs_id"`
		Path           string            `json:"path"`
		ServerFileName string            `json:"server_filename"`
		Size           int               `json:"size"`
		IsDir          int               `json:"isdir"`
		Category       int               `json:"category"`
		Md5            string            `json:"md5"`
		DirEmpty       int               `json:"dir_empty"`
		Thumbs         map[string]string `json:"thumbs"`
		LocalCtime     int               `json:"local_ctime"`
		LocalMtime     int               `json:"local_mtime"`
		ServerCtime    int               `json:"server_ctime"`
		ServerMtime    int               `json:"server_mtime"`
	}
}

type MetasResponse struct {
	ErrorCode    int    `json:"errno"`
	ErrorMsg     string `json:"errmsg"`
	RequestID    int
	RequestIDStr string `json:"request_id"`
	List         []struct {
		FsID        uint64            `json:"fs_id"`
		Path        string            `json:"path"`
		Category    int               `json:"category"`
		FileName    string            `json:"filename"`
		IsDir       int               `json:"isdir"`
		Size        int               `json:"size"`
		Md5         string            `json:"md5"`
		DLink       string            `json:"dlink"`
		Thumbs      map[string]string `json:"thumbs"`
		ServerCtime int               `json:"server_ctime"`
		ServerMtime int               `json:"server_mtime"`
		DateTaken   int               `json:"date_taken"`
		Width       int               `json:"width"`
		Height      int               `json:"height"`
	}
}

type ManagerResponse struct {
	conf.CloudDiskResponseBase
	Info []struct {
		Path   string
		TaskID int
		Errno  int
	}
}

type CreateDirResponse struct {
	conf.CloudDiskResponseBase
	Path     string `json:"path"`
	Ctime    int    `json:"ctime"`
	Mtime    int    `json:"mtime"`
	FsID     uint64 `json:"fs_id"`
	IsDir    int    `json:"isdir"`
	Category int    `json:"category"`
}

/**/

type UploadResponse struct {
	conf.CloudDiskResponseBase
	Path  string `json:"path"`
	Size  int    `json:"size"`
	Ctime int    `json:"ctime"`
	Mtime int    `json:"mtime"`
	Md5   string `json:"md5"`
	FsID  uint64 `json:"fs_id"`
	IsDir int    `json:"isdir"`
}

type PreCreateResponse struct {
	conf.CloudDiskResponseBase
	UploadID   string         `json:"uploadid"`
	Path       string         `json:"path"`
	ReturnType int            `json:"return_type"`
	BlockList  []int          `json:"block_list"`
	Info       UploadResponse `json:"info"`
}

type SuperFile2UploadResponse struct {
	conf.PcsResponseBase
	Md5      string `json:"md5"`
	UploadID string `json:"uploadid"`
	PartSeq  string `json:"partseq"` //pcsapi PHP版本返回的是int类型，Go版本返回的是string类型
}
