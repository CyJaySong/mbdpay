package mbdpay

import "encoding/json"

// ChargeSucceededBody 支付成功消息结果
type ChargeSucceededBody struct {
	Description string `json:"description"`  // 商品描述
	OutTradeNo  string `json:"out_trade_no"` // 订单号
	Amount      uint64 `json:"amount"`       // 订单金额，单位为分
	Openid      string `json:"openid"`       // 支付者 openid (仅微信支付)
	ChargeId    string `json:"charge_id"`    // 支付渠道流水号
	PayWay      int    `json:"payway"`       // 支付渠道，微信支付为 1 ，支付宝支付为 2
}

// ComplaintBody 订单投诉消息结果
type ComplaintBody struct {
	OutTradeNo      string `json:"out_trade_no"`     // 订单号
	ComplaintDetail string `json:"complaint_detail"` // 投诉详情
	Amount          uint64 `json:"amount"`           // 订单金额，单位为分
	PayerPhone      string `json:"payer_phone"`      // 投诉者电话
}

// UnmarshalChargeSucceededBody 解析 UnmarshalChargeSucceededBody
func UnmarshalChargeSucceededBody(data interface{}) (chargeSucceededBody *ChargeSucceededBody, err error) {
	chargeSucceededBody = new(ChargeSucceededBody)
	switch v := data.(type) {
	case []byte:
		err = json.Unmarshal(v, chargeSucceededBody)
	case string:
		err = json.Unmarshal([]byte(v), chargeSucceededBody)
	case map[string]interface{}:
		dataBytes, _ := json.Marshal(v)
		err = json.Unmarshal(dataBytes, chargeSucceededBody)
	default:
		return nil, err
	}
	return
}

// UnmarshalComplaintBody 解析 ComplaintBody
func UnmarshalComplaintBody(data interface{}) (complaintBody *ComplaintBody, err error) {
	complaintBody = new(ComplaintBody)
	switch v := data.(type) {
	case []byte:
		err = json.Unmarshal(v, complaintBody)
	case string:
		err = json.Unmarshal([]byte(v), complaintBody)
	case map[string]interface{}:
		dataBytes, _ := json.Marshal(v)
		err = json.Unmarshal(dataBytes, complaintBody)
	default:
		return nil, err
	}
	return
}
