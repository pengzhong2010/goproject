package util

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpClientPostIn struct {
	Url    string
	Body   []byte
	Header map[string]string
}

type HttpClientPostOut struct {
	StatusCode int64
	Content    []byte
	//Header     map[string]string
}

func HttpClientPost(c context.Context, in *HttpClientPostIn) (out *HttpClientPostOut, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", in.Url, bytes.NewBuffer(in.Body))
	if err != nil {
		return
	}
	for k, v := range in.Header {
		req.Header.Set(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	stateCode := int64(res.StatusCode)
	out = &HttpClientPostOut{
		StatusCode: stateCode,
		Content:    content,
	}
	return
}
func UrlEncode(in string) (out string) {
	return url.QueryEscape(in)
}

func UrlDecode(in string) (out string, err error) {
	return url.QueryUnescape(in)
}
