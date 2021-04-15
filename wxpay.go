package mbdpay

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// CreateGetWxOpenIdUrl 创建获取用户openid Url
// 请求参数 https://doc.mbd.pub/api/huo-qu-yong-hu-openid
func (c Client) CreateGetWxOpenIdUrl(targetUrl string) string {
	return "https://mbd.pub/openid?app_id=" + c.appId + "&target_url=" + targetUrl
}

type WxJsPayParam struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}
type WxJsPayResp struct {
	*WxJsPayParam
	Error string `json:"error"`
}

// WxJsPay 微信 JSAPI 支付
// 请求参数 https://doc.mbd.pub/api/wei-xin-zhi-fu
func (c Client) WxJsPay(openid, outTradeNo, description, callbackUrl string, amountTotal uint64) (resp *WxJsPayResp, err error) {
	param := make(map[string]string, 7)
	param["app_id"] = c.appId
	param["amount_total"] = strconv.FormatUint(amountTotal, 10)
	param["openid"] = openid
	param["callback_url"] = callbackUrl
	param["description"] = description
	param["out_trade_no"] = outTradeNo
	if param["sign"], err = c.sign(param); err != nil {
		return nil, err
	}
	body, err := c.doRequest(wxpayPath, param)
	if err != nil {
		return nil, err
	}
	resp = new(WxJsPayResp)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, err
	} else if resp.WxJsPayParam == nil && len(resp.Error) == 0 {
		return nil, errors.New(string(body))
	}
	return
}

type WxH5PayResp struct {
	H5Url string `json:"h5_url"`
	Error string `json:"error"`
}

// WxH5Pay 微信 H5 支付
// 请求参数 https://doc.mbd.pub/api/wei-xin-h5-zhi-fu
func (c Client) WxH5Pay(outTradeNo, description string, amountTotal uint64) (resp *WxH5PayResp, err error) {
	param := make(map[string]string, 7)
	param["app_id"] = c.appId
	param["amount_total"] = strconv.FormatUint(amountTotal, 10)
	param["channel"] = "h5"
	param["description"] = description
	param["out_trade_no"] = outTradeNo
	if param["sign"], err = c.sign(param); err != nil {
		return nil, err
	}
	body, err := c.doRequest(wxpayPath, param)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s \n", body)
	resp = new(WxH5PayResp)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, err
	} else if len(resp.H5Url) == 0 && len(resp.Error) == 0 {
		return nil, errors.New(string(body))
	}
	return
}
