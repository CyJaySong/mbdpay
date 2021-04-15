package mbdpay

const (
	kContentType = "content-type: application/json; charset=utf-8"

	apiBaseUrl = "https://api.mianbaoduo.com"
	// 接口 Path

	wxpayPath  = "/release/wx/prepay"         // 微信 预支付
	alipayPath = "/release/alipay/pay"        // 支付宝支付
	refundPath = "/release/main/refund"       // 订单退款
	searchOrderPath = "/release/main/search_order" // 订单查询

)
