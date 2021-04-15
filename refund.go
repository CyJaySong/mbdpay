package mbdpay

import (
	"encoding/json"
	"errors"
)

type RefundBody struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}

type RefundResp struct {
	*RefundBody
	Error string `json:"error"`
}

// Refund 退款
// 请求参数 https://doc.mbd.pub/api/tui-kuan
func (c Client) Refund(orderId string) (resp *RefundResp, err error) {
	param := make(map[string]string, 3)
	param["app_id"] = c.appId
	param["order_id"] = orderId
	if param["sign"], err = c.sign(param); err != nil {
		return nil, err
	}
	body, err := c.doRequest(refundPath, param)
	if err != nil {
		return nil, err
	}
	resp = new(RefundResp)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, err
	} else if resp.RefundBody == nil && len(resp.Error) == 0 {
		return nil, errors.New(string(body))
	}
	return
}
