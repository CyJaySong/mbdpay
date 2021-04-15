package mbdpay

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SearchOrderBody struct {
	OrderId      string `json:"order_id"`      // 订单号
	ChargeId     string `json:"charge_id"`     // 支付渠道流水号
	Description  string `json:"description"`   // 商品描述
	ShareId      string `json:"share_id"`      // 结算ID
	ShareState   int    `json:"share_state"`   // 结算状态
	Amount       uint64 `json:"amount"`        // 支付金额，单位为分
	State        int    `json:"state"`         // 支付状态，0-未支付，1-已支付，2-已结算，3-投诉中，4-投诉完结，5-投诉超时，6-投诉中(买家处理中)
	CreateTime   int    `json:"create_time"`   // 支付时间（时间戳）
	PayWay       int    `json:"payway"`        // 支付渠道，1为微信支付，2为支付宝
	RefundState  int    `json:"refund_state"`  // 退款状态，0为无退款，1为部分退款，2为全部退款
	RefundAmount uint64 `json:"refund_amount"` // 已退款金额，单位为分
	Plusinfo     string `json:"plusinfo"`      // 附加参数（json格式）
}

type SearchOrderResp struct {
	*SearchOrderBody
	Error string `json:"error"`
}

// SearchOrder 订单查询
// 请求参数 https://doc.mbd.pub/api/ding-dan-cha-xun
func (c Client) SearchOrder(outTradeNo string) (resp *SearchOrderResp, err error) {
	param := make(map[string]string, 3)
	param["app_id"] = c.appId
	param["out_trade_no"] = outTradeNo
	if param["sign"], err = c.sign(param); err != nil {
		return nil, err
	}
	body, err := c.doRequest(searchOrderPath, param)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body :%s\n", body)
	resp = new(SearchOrderResp)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, err
	} else if resp.SearchOrderBody == nil && len(resp.Error) == 0 {
		return nil, errors.New(string(body))
	}
	return
}
