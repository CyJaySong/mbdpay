package mbdpay

import (
	"encoding/json"
	"errors"
	"strconv"
)

type AliPayResp struct {
	Body  string `json:"body"`
	Error string `json:"error"`
}

// AliPay 支付宝支付
// 请求参数 https://doc.mbd.pub/api/zhi-fu-bao-zhi-fu
func (c Client) AliPay(outTradeNo, description, callbackUrl, url string, amountTotal uint64) (resp *AliPayResp, err error) {
	param := make(map[string]string, 7)
	param["app_id"] = c.appId
	param["amount_total"] = strconv.FormatUint(amountTotal, 10)
	param["callback_url"] = callbackUrl
	param["description"] = description
	param["out_trade_no"] = outTradeNo
	param["url"] = url
	if param["sign"], err = c.sign(param); err != nil {
		return nil, err
	}
	body, err := c.doRequest(alipayPath, param)
	if err != nil {
		return nil, err
	}
	resp = new(AliPayResp)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, err
	} else if len(resp.Body) == 0 && len(resp.Error) == 0 {
		return nil, errors.New(string(body))
	}
	return
}
