package mbdpay

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

var httpClient = http.DefaultClient

type Client struct {
	appId  string // APP ID，可在控制台查看
	appKey string // APP KEY，可在控制台查看
}

// New 初始化面包多客户端
func New(appId, appKey string) (*Client, error) {
	if len(appId) == 0 {
		return nil, errors.New("appId 不能为空")
	}
	if len(appKey) == 0 {
		return nil, errors.New("appKey 不能为空")
	}
	return &Client{appId: appId, appKey: appKey}, nil
}

// 发起请求
func (c Client) doRequest(path string, param map[string]string) (body []byte, err error) {
	paramBytes, _ := json.Marshal(param)
	resp, err := httpClient.Post(apiBaseUrl+path, kContentType, bytes.NewReader(paramBytes))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

// 内容签名
func (c Client) sign(param map[string]string) (string, error) {
	if len(c.appId) == 0 {
		return "", errors.New("appId 不能为空")
	}
	if len(c.appKey) == 0 {
		return "", errors.New("appKey 不能为空")
	}
	params := make([]string, 0, len(param)+1)
	for key, value := range param {
		if len(value) > 0 {
			params = append(params, key+"="+value)
		}
	}
	sort.Strings(params)
	params = append(params, "key="+c.appKey)
	paramStr := strings.Join(params, "&")
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(paramStr))
	return hex.EncodeToString(md5Ctx.Sum(nil)), nil
}
