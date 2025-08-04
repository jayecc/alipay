package alipay

const (
	NotifyTypeTradeStatusSync = "trade_status_sync"
)

// TradeNotification
// Deprecated: use Notification instead.
type TradeNotification Notification

// Notification 通知响应参数 https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.8AmJwg&treeId=203&articleId=105286&docType=1
type Notification struct {
	AuthAppId           string      `json:"auth_app_id"`           // App Id
	NotifyTime          string      `json:"notify_time"`           // 通知时间
	NotifyType          string      `json:"notify_type"`           // 通知类型
	NotifyId            string      `json:"notify_id"`             // 通知校验ID
	AppId               string      `json:"app_id"`                // 开发者的app_id
	Charset             string      `json:"charset"`               // 编码格式
	Version             string      `json:"version"`               // 接口版本
	SignType            string      `json:"sign_type"`             // 签名类型
	Sign                string      `json:"sign"`                  // 签名
	TradeNo             string      `json:"trade_no"`              // 支付宝交易号
	OutTradeNo          string      `json:"out_trade_no"`          // 商户订单号
	OutBizNo            string      `json:"out_biz_no"`            // 商户业务号
	BuyerId             string      `json:"buyer_id"`              // 买家支付宝用户号
	BuyerLogonId        string      `json:"buyer_logon_id"`        // 买家支付宝账号
	SellerId            string      `json:"seller_id"`             // 卖家支付宝用户号
	SellerEmail         string      `json:"seller_email"`          // 卖家支付宝账号
	TradeStatus         TradeStatus `json:"trade_status"`          // 交易状态
	TotalAmount         string      `json:"total_amount"`          // 订单金额
	ReceiptAmount       string      `json:"receipt_amount"`        // 实收金额
	InvoiceAmount       string      `json:"invoice_amount"`        // 开票金额
	BuyerPayAmount      string      `json:"buyer_pay_amount"`      // 付款金额
	PointAmount         string      `json:"point_amount"`          // 集分宝金额
	RefundFee           string      `json:"refund_fee"`            // 总退款金额
	Subject             string      `json:"subject"`               // 商品的标题/交易标题/订单标题/订单关键字等，是请求时对应的参数，原样通知回来。
	Body                string      `json:"body"`                  // 商品描述
	GmtCreate           string      `json:"gmt_create"`            // 交易创建时间
	GmtPayment          string      `json:"gmt_payment"`           // 交易付款时间
	GmtRefund           string      `json:"gmt_refund"`            // 交易退款时间
	GmtClose            string      `json:"gmt_close"`             // 交易结束时间
	FundBillList        string      `json:"fund_bill_list"`        // 支付金额信息
	PassbackParams      string      `json:"passback_params"`       // 回传参数
	VoucherDetailList   string      `json:"voucher_detail_list"`   // 优惠券信息
	AgreementNo         string      `json:"agreement_no"`          // 支付宝签约号
	ExternalAgreementNo string      `json:"external_agreement_no"` // 商户自定义签约号
}

// MerchantReviewNotification 商户审核通知响应参数
// https://opendocs.alipay.com/pre-apis/msgapi_pre/ant.merchant.expand.indirect.zft.passed
// https://opendocs.alipay.com/pre-apis/msgapi_pre/ant.merchant.expand.indirect.zft.rejected
type MerchantReviewNotification struct {
	NotifyId     string             `json:"notify_id"`     //通知ID	5608cccc09ddb39d41c2e3c06e3d9fejh2
	UtcTimestamp string             `json:"utc_timestamp"` //消息发送时的服务端时间	1514210452731
	MsgMethod    string             `json:"msg_method"`    //消息接口名称	ant.merchant.expand.indirect.zft.passed
	AppId        string             `json:"app_id"`        //消息接受方的应用id	2014060600164699
	MsgType      string             `json:"msg_type"`      //消息类型。目前支持类型：sys：系统消息；usr，用户消息；app，应用消息	sys
	MsgUid       string             `json:"msg_uid"`       //消息归属的商户支付宝uid。用户消息和应用消息时非空	2088102165945162
	MsgAppId     string             `json:"msg_app_id"`    // 消息归属方的应用id。应用消息时非空	2016032301002387
	Version      string             `json:"version"`       //版本号(1.1版本为标准消息)	1.0或者1.1
	BizContent   MerchantBizContent `json:"biz_content"`   //是		消息报文	<a href="#s1">参见具体的消息接口文档</a>
	Sign         string             `json:"sign"`          //签名
	SignType     string             `json:"sign_type"`     //签名类型	RSA2
	EncryptType  string             `json:"encrypt_type"`  //加密算法	AES
	Charset      string             `json:"charset"`       //编码集，该字符集为验签和解密所需要的字符集	UTF-8
}

type MerchantBizContent struct {
	OrderId     string `json:"order_id"`      //申请单id	2018011900502000000005124744
	Smid        string `json:"smid"`          //二级商户id	2088301031335814
	ExternalId  string `json:"external_id"`   //外部商户id	105290059990194
	CardAliasNo string `json:"card_alias_no"` //银行卡编号	c6c0c7a1b9d54e5db9d49eed39f00e65
	Memo        string `json:"memo"`          //审核备注信息	符合要求，通过
	Reason      string `json:"reason"`        //拒绝原因	资料有问题，拒绝  拒绝时无smid 和 card_alias_no
}
