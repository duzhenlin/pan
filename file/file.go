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
)

type File struct {
	AccessToken string
}

func NewFileClient(accessToken string) *File {
	return &File{
		AccessToken: accessToken,
	}
}

// List 获取文件列表
func (f *File) List(dir string, start, limit int) (ListResponse, error) {
	ret := ListResponse{}

	v := url.Values{}
	v.Add("access_token", f.AccessToken)
	v.Add("dir", dir)
	v.Add("start", strconv.Itoa(start))
	v.Add("limit", strconv.Itoa(limit))
	query := v.Encode()

	requestUrl := conf.OpenApiDomain + ListUri + "&" + query
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

func (f *File) CategoryList(dir string, start, limit int) (ListResponse, error) {
	ret := ListResponse{}

	v := url.Values{}
	v.Add("access_token", f.AccessToken)
	v.Add("dir", dir)
	v.Add("start", strconv.Itoa(start))
	v.Add("limit", strconv.Itoa(limit))
	query := v.Encode()

	requestUrl := conf.OpenApiDomain + CategoryListUri + "&" + query
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

// CreateDir 创建目录
/*
 */
func (f *File) CreateDir(opts CreateDirOption) (CreateDirResponse, error) {

	rType := "1"
	path := "/"
	mode := "1"
	if opts.RType != "" {
		rType = opts.RType
	}
	if opts.Path != "" {
		path = opts.Path
	}
	if opts.Mode != "" {
		mode = opts.Mode
	}

	ret := CreateDirResponse{}
	// path urlencode
	v := url.Values{}
	v.Add("path", path)
	v.Add("isdir", "1")
	v.Add("rtype", rType)
	v.Add("mode", mode)
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

// Metas 通过FsID获取文件信息
func (f *File) Metas(fsIDs []uint64) (MetasResponse, error) {
	ret := MetasResponse{}

	fsIDsByte, err := json.Marshal(fsIDs)
	if err != nil {
		return ret, err
	}

	v := url.Values{}
	v.Add("access_token", f.AccessToken)
	v.Add("fsids", string(fsIDsByte))
	v.Add("dlink", "1")
	v.Add("thumb", "1")
	v.Add("extra", "1")
	query := v.Encode()

	requestUrl := conf.OpenApiDomain + MetasUri + "&" + query
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
