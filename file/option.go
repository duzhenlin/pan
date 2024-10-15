// Package file
// Created by Duzhenlin
// @Author   Duzhenlin
// @Email: duzhenlin@vip.qq.com
// @Date: 2024/10/15
// @Time: 15:48

package file

// CreateDirOption 创建目录选项字段
type CreateDirOption struct {
	/*
		文件命名策略，默认0
		0 为不重命名，返回冲突
		1 为只要path冲突即重命名
	*/
	RType string `json:"r_type"`
	/*
		创建文件夹的绝对路径，需要urlencode
	*/
	Path string `json:"path"`
	/*
		上传方式
		1 手动、2 批量上传、3 文件自动备份、4 相册自动备份、5 视频自动备份
	*/
	Mode string `json:"mode"`
}

// CategoryListOption 分类列表选项字段
type CategoryListOption struct {
	/*
		文件类型，1 视频、2 音频、3 图片、4 文档、5 应用、6 其他、7 种子
		多个category使用英文逗号分隔，示例：3,4
	*/
	Category []string `json:"category,omitempty"`
	/*
		是否展示文件夹，0:否(默认) 1:是
	*/
	ShowDir int `json:"show_dir,omitempty"`
	/*
		目录名称，为空时，parent_path = "/" && recursion = 1 ；路径包含中文时需要进行UrlEncode编码
	*/
	ParentPath string `json:"parent_path,omitempty"`
	/*
		是否需要递归，0 不递归、1 递归，默认0 （注意recursion=1时不支持show_dir=1）
	*/
	Recursion int `json:"recursion,omitempty"`
	/*
		需要的文件格式，多个格式以英文逗号分隔，示例: txt,epub，默认为category下所有格式
	*/
	Ext string `json:"ext,omitempty"`
	/*
		起始位置，从0开始
	*/
	Start int `json:"start,omitempty"`
	/*
		查询数目，默认为1000，建议最大不超过1000
	*/
	Limit int `json:"limit,omitempty"`
	/*
		排序字段：默认为name；
		time按修改时间排序；
		name表示按文件名称排序；(注意，此处排序是按字符串排序的，如果用户有剧集排序需求，需要自行开发)
		size表示按文件大小排序。
	*/
	Order string `json:"order,omitempty"`
	/*
		默认为升序，设置为1实现降序 （注：排序的对象是当前目录下所有文件，不是当前分页下的文件）
	*/
	Desc int `json:"desc,omitempty"`
}

// ListOption 文件列表选项字段/*
type ListOption struct {
	/*
		需要list的目录，以/开头的绝对路径, 默认为/
		路径包含中文时需要UrlEncode编码
		给出的示例的路径是/测试目录的UrlEncode编码
	*/
	Dir string `json:"dir,omitempty" ` // 目录
	/*
		排序字段：默认为name；
		time按修改时间排序；
		name表示按文件名称排序；(注意，此处排序是按字符串排序的，如果用户有剧集排序需求，需要自行开发)
		size表示按文件大小排序。
	*/
	Order string `json:"order,omitempty"`
	/*
		默认为升序，设置为1实现降序 （注：排序的对象是当前目录下所有文件，不是当前分页下的文件）
	*/
	Desc int `json:"desc,omitempty"`
	/*
		起始位置，从0开始
	*/
	Start int `json:"start,omitempty"`
	/*
		查询数目，默认为1000，建议最大不超过1000
	*/
	Limit int `json:"limit,omitempty"`
	/*
		值为1时，返回dir_empty属性和缩略图数据
	*/
	Web int `json:"web,omitempty"`
	/*
		是否只返回文件夹，0 返回所有，1 只返回文件夹，且属性只返回path字段
	*/
	Folder int `json:"folder,omitempty"`
	/*
		是否返回dir_empty属性，0 不返回，1 返回
	*/
	ShowEmpty int `json:"show_empty,omitempty"`
}

type MetasOption struct {
	/*
		文件id数组，数组中元素是uint64类型，数组大小上限是：100
	*/
	FsIDs []uint64 `json:"fsids,omitempty"`
	/*
		是否需要下载地址，0为否，1为是，默认为0。获取到dlink后，参考下载文档进行下载操作
	*/
	DLink int `json:"dlink,omitempty"`
	/*
		查询共享目录或专属空间内文件时需要。
		共享目录格式： /uk-fsid
		其中uk为共享目录创建者id， fsid对应共享目录的fsid
		专属空间格式：/_pcs_.appdata/xpan/
	*/
	Path string `json:"path,omitempty"`
	/*
		是否需要缩略图地址，0为否，1为是，默认为0
	*/
	Thumb int `json:"thumb,omitempty"`
	/*
		图片是否需要拍摄时间、原图分辨率等其他信息，0 否、1 是，默认0
	*/
	Extra int `json:"extra,omitempty"`
	/*
		视频是否需要展示时长信息，needmedia=1时，返回 duration 信息时间单位为秒 （s），转换为向上取整。
		0 否、1 是，默认0
	*/
	NeedMedia int `json:"needmedia,omitempty"`
	/*
		视频是否需要展示长，宽等信息。
		0 否、1 是，默认0
	*/
	Detail int `json:"detail,omitempty"`
}

const ListOptionTypeName = "name"
const ListOptionTypeTime = "time"
const ListOptionTypeSize = "size"

const ListCategoryTypeVideo = "1"
const ListCategoryTypeAudio = "2"
const ListCategoryTypeImage = "3"
const ListCategoryTypeDoc = "4"
const ListCategoryTypeApp = "5"
const ListCategoryTypeOther = "6"
const ListCategoryTypeSeed = "7"
