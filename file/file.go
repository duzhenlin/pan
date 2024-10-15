package file

import (
	"encoding/json"
	"fmt"
	"github.com/duzhenlin/pan/conf"
	"github.com/duzhenlin/pan/error_pan"
	"github.com/duzhenlin/pan/utils/httpclient"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type File struct {
	AccessToken string
}

func NewFileClient(accessToken string) *File {
	return &File{
		AccessToken: accessToken,
	}
}
func (f *File) NewDefaultListOption() (option *ListOption) {
	option = &ListOption{
		Desc:      0,
		Dir:       "/",
		Folder:    0,
		Start:     1,
		Limit:     100,
		Order:     ListOptionTypeName,
		ShowEmpty: 0,
		Web:       0,
	}
	return option
}

// List 获取文件列表
func (f *File) List(option *ListOption) (ListResponse, error) {
	ret := ListResponse{}

	v := url.Values{}
	v.Add("access_token", f.AccessToken)
	v.Add("dir", option.Dir)
	v.Add("start", strconv.Itoa(option.Start))
	v.Add("limit", strconv.Itoa(option.Limit))
	v.Add("order", option.Order)
	v.Add("web", strconv.Itoa(option.Web))
	v.Add("folder", strconv.Itoa(option.Folder))
	v.Add("showempty", strconv.Itoa(option.ShowEmpty))
	v.Add("desc", strconv.Itoa(option.Desc))
	query := v.Encode()
	requestUrl := conf.OpenApiDomain + ListUri + "&" + query
	log.Println("requestUrl:", requestUrl)
	resp, err := httpclient.Get(requestUrl, map[string]string{})
	if err != nil {
		log.Println("httpclient.Get failed, err:", err)
		return ret, err
	}

	if resp.StatusCode != 200 {
		return ret, error_pan.NewBaiduPanError(
			ret.ErrorCode,
			fmt.Sprintf("HttpStatusCode is not equal to 200, httpStatusCode[%d], respBody[%s]", resp.StatusCode, string(resp.Body)),
			strconv.FormatUint(ret.RequestID, 10),
		)
	}

	if err := json.Unmarshal(resp.Body, &ret); err != nil {
		return ret, err
	}

	if ret.ErrorCode != 0 { //错误码不为0
		return ret, error_pan.NewBaiduPanError(
			ret.ErrorCode,
			fmt.Sprintf("error_code:%d, error_msg:%s", ret.ErrorCode, ret.ErrorMsg),
			strconv.FormatUint(ret.RequestID, 10),
		)
	}

	return ret, nil
}

func (f *File) NewCategoryListOption() (option *CategoryListOption) {
	option = &CategoryListOption{
		Category:   []string{"1", "2", "3"},
		ParentPath: "/",
		Recursion:  0,
		Start:      0,
		Limit:      100,
		Order:      ListOptionTypeTime,
	}
	return option
}
func (f *File) CategoryList(option *CategoryListOption) (CategoryListResponse, error) {
	ret := CategoryListResponse{}
	if option.Recursion == 1 && option.ShowDir == 1 {
		return ret, error_pan.NewBaiduPanError(
			-100,
			"recursion=1时不支持show_dir=1",
			strconv.FormatUint(ret.RequestID, 10),
		)
	}
	v := url.Values{}
	if len(option.Category) > 3 {
		return ret, error_pan.NewBaiduPanError(
			-100,
			"Category 最多可以选择3个",
			strconv.FormatUint(ret.RequestID, 10),
		)
	}
	category := strings.Join(option.Category, ",")
	v.Add("access_token", f.AccessToken)
	v.Add("category", category)
	v.Add("parent_path", option.ParentPath)
	v.Add("recursion", strconv.Itoa(option.Recursion))
	if option.Ext != "" {
		v.Add("ext", option.Ext)
	}
	v.Add("show_dir", strconv.Itoa(option.ShowDir))
	v.Add("start", strconv.Itoa(option.Start))
	v.Add("limit", strconv.Itoa(option.Limit))
	v.Add("order", option.Order)
	v.Add("desc", strconv.Itoa(option.Desc))
	fmt.Printf("Values : %v \n", v)

	query := v.Encode()
	fmt.Printf(" query : %v \n", query)
	requestUrl := conf.OpenApiDomain + CategoryListUri + "&" + query
	log.Println("requestUrl:", requestUrl)
	resp, err := httpclient.Get(requestUrl, map[string]string{})
	if err != nil {
		log.Println("httpclient.Get failed, err:", err)
		return ret, err
	}

	if resp.StatusCode != 200 {
		return ret, error_pan.NewBaiduPanError(
			ret.ErrorCode,
			fmt.Sprintf("HttpStatusCode is not equal to 200, httpStatusCode[%d], respBody[%s]", resp.StatusCode, string(resp.Body)),
			strconv.FormatUint(ret.RequestID, 10),
		)
	}

	if err := json.Unmarshal(resp.Body, &ret); err != nil {
		return ret, err
	}

	if ret.ErrorCode != 0 { //错误码不为0
		return ret, error_pan.NewBaiduPanError(
			ret.ErrorCode,
			fmt.Sprintf("error_code:%d, error_msg:%s", ret.ErrorCode, ret.ErrorMsg),
			strconv.FormatUint(ret.RequestID, 10),
		)
	}

	return ret, nil
}

func (f *File) NewCreateDirOption() (option *CreateDirOption) {
	option = &CreateDirOption{
		Mode:  "1",
		RType: "1",
		Path:  "/",
	}
	return option
}

// CreateDir 创建目录
func (f *File) CreateDir(option *CreateDirOption) (CreateDirResponse, error) {
	ret := CreateDirResponse{}
	// 参数校验
	if option.RType == "" {
		return ret, error_pan.NewBaiduPanError(-100, "RType 不能为空", "")
	}

	if option.Path == "" {
		return ret, error_pan.NewBaiduPanError(-100, "Path 不能为空", "")
	}

	if option.Mode == "" {
		return ret, error_pan.NewBaiduPanError(-100, "Mode 不能为空", "")
	}
	// 参数组合
	v := url.Values{}
	v.Add("path", option.Path)
	v.Add("rtype", option.RType)
	v.Add("mode", option.Mode)
	v.Add("isdir", "1")
	body := v.Encode()

	requestUrl := conf.OpenApiDomain + CreateUri + "&access_token=" + f.AccessToken

	headers := make(map[string]string)
	resp, err := httpclient.Post(requestUrl, headers, body)
	if err != nil {
		log.Println("httpclient.Post failed, err:", err)
		return ret, err
	}

	if err := json.Unmarshal(resp.Body, &ret); err != nil {
		log.Printf("json.Unmarshal failed, resp[%s], err[%v]", string(resp.Body), err)
		return ret, err
	}

	if ret.ErrorCode != 0 { //错误码不为0
		log.Println("file create failed, resp:", string(resp.Body))
		return ret, error_pan.NewBaiduPanError(
			ret.ErrorCode,
			fmt.Sprintf("error_code:%d, error_msg:%s", ret.ErrorCode, ret.ErrorMsg),
			strconv.FormatUint(ret.RequestID, 10),
		)
	}

	return ret, nil
}

func (f *File) NewMetasOption() (option *MetasOption) {
	option = &MetasOption{
		Path:      "",
		DLink:     1,
		Thumb:     1,
		Extra:     1,
		NeedMedia: 1,
		Detail:    1,
	}
	return option
}

// Metas 通过FsID获取文件信息
func (f *File) Metas(option *MetasOption) (MetasResponse, error) {
	ret := MetasResponse{}
	if len(option.FsIDs) > 100 {
		return ret, error_pan.NewBaiduPanError(-100, "FsIDs 数组大小上限是：100", "")
	}
	fsIDsByte, err := json.Marshal(option.FsIDs)
	if err != nil {
		return ret, err
	}

	v := url.Values{}
	v.Add("access_token", f.AccessToken)
	v.Add("fsids", string(fsIDsByte))
	v.Add("dlink", strconv.Itoa(option.DLink))
	v.Add("thumb", strconv.Itoa(option.Thumb))
	v.Add("extra", strconv.Itoa(option.Extra))
	v.Add("needmedia", strconv.Itoa(option.NeedMedia))
	v.Add("detail", strconv.Itoa(option.Detail))
	if option.Path != "" {
		v.Add("path", option.Path)
	}
	query := v.Encode()
	requestUrl := conf.OpenApiDomain + MetasUri + "&" + query
	log.Println("requestUrl:", requestUrl)
	resp, err := httpclient.Get(requestUrl, map[string]string{})
	if err != nil {
		log.Println("httpclient.Get failed, err:", err)
		return ret, err
	}

	if resp.StatusCode != 200 {
		return ret, error_pan.NewBaiduPanError(
			ret.ErrorCode,
			fmt.Sprintf("HttpStatusCode is not equal to 200, httpStatusCode[%d], respBody[%s]", resp.StatusCode, string(resp.Body)),
			ret.RequestIDStr,
		)
	}

	if err := json.Unmarshal(resp.Body, &ret); err != nil {
		return ret, err
	}

	if ret.ErrorCode != 0 { //错误码不为0
		return ret, error_pan.NewBaiduPanError(
			ret.ErrorCode,
			fmt.Sprintf("error_code:%d, error_msg:%s", ret.ErrorCode, ret.ErrorMsg),
			ret.RequestIDStr,
		)
	}

	ret.RequestID, err = strconv.Atoi(ret.RequestIDStr)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// Streaming 获取音视频在线播放地址，转码类型有M3U8_AUTO_480=>视频ts、M3U8_FLV_264_480=>视频flv、M3U8_MP3_128=>音频mp3、M3U8_HLS_MP3_128=>音频ts
func (f *File) Streaming(path string, transcodingType string) (string, error) {
	ret := ""

	v := url.Values{}
	v.Add("access_token", f.AccessToken)
	v.Add("path", path)
	v.Add("type", transcodingType)
	query := v.Encode()

	requestUrl := conf.OpenApiDomain + StreamingUri + "&" + query
	resp, err := httpclient.Get(requestUrl, map[string]string{})
	if err != nil {
		log.Println("httpclient.Get failed, err:", err)
		return ret, err
	}

	if resp.StatusCode != 200 {
		return ret, error_pan.NewBaiduPanError(
			resp.StatusCode,
			fmt.Sprintf("HttpStatusCode is not equal to 200, httpStatusCode[%d], respBody[%s]", resp.StatusCode, string(resp.Body)),
			"",
		)
	}

	return string(resp.Body), nil
}
