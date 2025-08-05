package alipay

type ContactInfo struct {
	// name｜联系人名字必选string(85)
	//【描述】联系人名字
	Name string `json:"name,omitempty"`
	// email｜电子邮箱string(128)
	//【描述】电子邮箱
	Email string `json:"email,omitempty"`
	// mobile｜手机号string(20)
	//【描述】手机号
	Mobile string `json:"mobile,omitempty"`
	// phone｜电话string(20)
	//【描述】电话
	Phone string `json:"phone,omitempty"`
	// id_card_no｜身份证号可选string(32)
	//【描述】身份证号
	IdCardNo string `json:"id_card_no,omitempty"`
}

type DefaultSettleRule struct {
	// default_settle_type必选string(64)
	//【描述】默认结算类型，可选值有bankCard/alipayAccount。bankCard表示结算到银行卡；alipayAccount表示结算到支付宝账号
	//【枚举值】
	//结算到银行卡: bankCard
	//结算到支付宝账号: alipayAccount
	DefaultSettleType string `json:"default_settle_type,omitempty"`
	// default_settle_target必选string(64)
	//【描述】默认结算目标。当默认结算类型为bankCard时填写银行卡卡号，其值需在进件填写的结算银行卡范围内；当默认结算类型为alipayAccount时填写支付宝账号登录号，其值需在进件填写的结算支付宝账号范围内。
	DefaultSettleTarget string `json:"default_settle_target,omitempty"`
}

type SettleCardInfo struct {
	// account_branch_name｜开户支行名必选string(64)
	//【描述】开户支行名
	AccountBranchName string `json:"account_branch_name,omitempty"`
	// account_holder_name｜卡户名必选string(64)
	//【描述】卡户名
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// account_inst_city｜开户行所在地-市必选string(10)
	//【描述】开户行所在地-市
	AccountInstCity string `json:"account_inst_city,omitempty"`
	// account_inst_id｜开户行简称缩写必选string(10)
	//【描述】开户行简称缩写
	AccountInstId string `json:"account_inst_id,omitempty"`
	// account_inst_name｜银行名称必选string(20)
	//【描述】银行名称
	AccountInstName string `json:"account_inst_name,omitempty"`
	// account_inst_province｜开户行所在地-省必选string(10)
	//【描述】开户行所在地-省
	AccountInstProvince string `json:"account_inst_province,omitempty"`
	// account_no｜银行卡号必选string(20)
	//【描述】银行卡号
	AccountNo string `json:"account_no,omitempty"`
	// account_type｜卡类型必选string(64)
	//【描述】卡类型
	//借记卡-DC
	//信用卡-CC
	//【枚举值】
	//借记卡: DC
	//信用卡: CC
	AccountType string `json:"account_type,omitempty"`
	// usage_type｜账号使用类型必选string(64)
	//【描述】账号使用类型
	//对公-01
	//对私-02
	//【枚举值】
	//对公: 01
	//对私: 02
	UsageType string `json:"usage_type,omitempty"`
	// bank_code｜联行号可选string(64)
	//【描述】联行号，详细需调用‘联行号关联分支行查询’API查看查询结果. https://opendocs.alipay.com/apis/0853v1
	BankCode string `json:"bank_code,omitempty"`
}

type AddressInfo struct {
	// city_code｜城市编码必选string(10)
	//【描述】城市编码。
	//蚂蚁店铺请按照蚂蚁店铺地区码 表格中填写。
	//直付通商户请按照直付通商户地区码 表格中内容填写。
	CityCode string `json:"city_code,omitempty"`
	// district_code｜区县编码必选string(10)
	//【描述】区县编码。
	//蚂蚁店铺请按照蚂蚁店铺地区码 表格中填写。
	//直付通商户请按照直付通商户地区码 表格中内容填写。
	DistrictCode string `json:"district_code,omitempty"`
	// address｜地址必选string(256)
	//【描述】地址。商户详细经营地址或人员所在地点
	Address string `json:"address,omitempty"`
	// province_code｜省份编码必选string(10)
	//【描述】省份编码。
	//蚂蚁店铺请按照蚂蚁店铺地区码 表格中填写。
	//直付通商户请按照直付通商户地区码 表格中内容填写。
	ProvinceCode string `json:"province_code,omitempty"`
	// poiid｜高德poiid可选string(16)
	//【描述】高德poiid
	Poiid string `json:"poiid,omitempty"`
	// longitude｜经度可选string(11)
	//【描述】经度，浮点型, 小数点后最多保留6位。
	//如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询
	Longitude string `json:"longitude,omitempty"`
	// latitude｜纬度可选string(10)
	//【描述】纬度，浮点型,小数点后最多保留6位
	//如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询
	Latitude string `json:"latitude,omitempty"`
	// type｜地址类型可选string(64)
	//【描述】地址类型。取值范围：BUSINESS_ADDRESS：经营地址（默认）
	//【枚举值】
	//经营地址: BUSINESS_ADDRESS
	AddressType string `json:"address_type,omitempty"`
}

type MerchantInvoiceInfo struct {
	// auto_invoice必选boolean(10)
	//【描述】是否自动开票，值为true/false
	AutoInvoice bool `json:"auto_invoice,omitempty"`
	// accept_electronic必选boolean(10)
	//【描述】是否接受电子发票 true/false
	AcceptElectronic bool `json:"accept_electronic,omitempty"`
	// tax_payer_qualification必选string(10)
	//【描述】纳税人资格种类:01一般纳税人，02小规模纳税人。一般纳税人开的是专票
	TaxPayerQualification string `json:"tax_payer_qualification,omitempty"`
	// title必选string(9999)
	//【描述】发票抬头
	Title string `json:"title,omitempty"`
	// tax_no必选string(64)
	//【描述】纳税人识别号
	TaxNo string `json:"tax_no,omitempty"`
	// tax_payer_valid必选string(8)
	//【描述】纳税人资格开始时间,yyyyMMdd格式
	TaxPayerValid string `json:"tax_payer_valid,omitempty"`
	// address必选string(64)
	//【描述】开票地址
	Address string `json:"address,omitempty"`
	// telephone必选string(12)
	//【描述】开票电话
	Telephone string `json:"telephone,omitempty"`
	// bank_account必选string(32)
	//【描述】银行账号
	BankAccount string `json:"bank_account,omitempty"`
	// mail_name必选string(64)
	//【描述】收件人名称
	MailName string `json:"mail_name,omitempty"`
	// mail_address必选AddressInfo(64)
	//【描述】收件人地
	MailAddress *AddressInfo `json:"mail_address,omitempty"`
	// mail_telephone必选string(64)
	//【描述】寄送电话
	MailTelephone string `json:"mail_telephone,omitempty"`
	// bank_name必选string(64)
	//【描述】开户行名称
	BankName string `json:"bank_name,omitempty"`
}

type SiteInfo struct {
	// site_type必选string(10)
	//【描述】网站类型
	//【枚举值】
	//网站: 01
	//APP: 02
	//服务窗: 03
	//公众号: 04
	//其他: 05
	//支付宝小程序: 06
	//手机网站/H5: 07
	SiteType string `json:"site_type,omitempty"`
	// site_url特殊可选string(256)
	//【描述】站点地址
	//【必选条件】当传入service，且包含jsapi支付时，sites的site_type=06, site_url必填
	SiteURL string `json:"site_url,omitempty"`
	// site_name特殊可选string(128)
	//【描述】站点名称
	//【必选条件】当传入service，且包含jsapi支付、小程序支付时，sites的site_type=06, site_name必填
	SiteName string `json:"site_name,omitempty"`
	// screenshot_image特殊可选string(256)
	//【描述】截图照片；当传入交易场景trade_scene，且当传入trade_scene=WAP、trade_scene=APP、trade_scene=PC时该参数必传
	//【必选条件】当传入交易场景trade_scene，且当传入trade_scene=WAP、trade_scene=APP、trade_scene=PC时该参数必传
	ScreenshotImage string `json:"screenshot_image,omitempty"`
	// tiny_app_id特殊可选string(64)
	//【描述】小程序appId; 当传入service，且包含jsapi支付时，sites的site_type=06, tiny_app_id必填。
	//【必选条件】当传入service，且包含jsapi支付时，sites的site_type=06, tiny_app_id必填
	TinyAppID string `json:"tiny_app_id,omitempty"`
	// account可选string(64)
	//【描述】测试账号
	Account string `json:"account,omitempty"`
	// password可选string(128)
	//【描述】测试密码
	Password string `json:"password,omitempty"`
	// status可选string(20)
	//【描述】上架状态；
	//【枚举值】
	//已上线: ONLINE
	//已上线-内部: ONLINE_INNER
	//未上线: OFFLINE
	Status string `json:"status,omitempty"`
	// auth_letter_image可选string(256)
	//【描述】授权函照片
	AuthLetterImage string `json:"auth_letter_image,omitempty"`
	// remark可选string(512)
	//【描述】备注说明
	Remark string `json:"remark,omitempty"`
	// remark_image可选string(256)
	//【描述】备注说明图片
	RemarkImage string `json:"remark_image,omitempty"`
	// site_domain可选string(256)
	//【描述】网站域名
	SiteDomain string `json:"site_domain,omitempty"`
	// icp_service_name可选string(128)
	//【描述】ICP备案主体信息服务名称
	ICPServiceName string `json:"icp_service_name,omitempty"`
	// icp_no可选string(64)
	//【描述】ICP备案/许可证号
	ICPNo string `json:"icp_no,omitempty"`
	// icp_org_name可选string(256)
	//【描述】ICP备案主体主办单位名称
	ICPServiceOrgName string `json:"icp_service_org_name,omitempty"`
	// download可选string(256)
	//【描述】下载地址
	Download string `json:"download,omitempty"`
	// market可选string(128)
	//【描述】应用市场
	Market string `json:"market,omitempty"`
}

type IndustryQualificationInfo struct {
	// industry_qualification_type｜商户行业资质类型必选string(6)
	//【描述】商户行业资质类型，具体选值参见文档 https://gw.alipayobjects.com/os/bmw-prod/7aa3a36b-2bc2-4d57-815f-08edd55ef67e.xlsx
	//【枚举值】
	//金融许可证: 323
	IndustryQualificationType string `json:"industry_qualification_type"`
	// industry_qualification_image｜商户行业资质图片必选string(256)
	//【描述】商户行业资质图片。其值为通过ant.merchant.expand.indirect.image.upload上传图片得到的image_id
	IndustryQualificationImage string `json:"industry_qualification_image"`
}

type ZFTWithholdingInfo struct {
	// withholding_service_feature_name可选string(64)
	//【描述】代扣产品业务模式；
	//【枚举值】
	//通用扣款: GENERAL_WITHHOLDING_P
	//周期性扣款: CYCLE_PAY_AUTH_P
	//单次授权: ONE_TIME_AUTH_P
	//免密支付: NOPWD_P
	WithholdingServiceFeatureName string `json:"withholding_service_feature_name"`
	// sign_scene可选string(64)
	//【描述】代扣签约场景码
	SignScene string `json:"sign_scene"`
}

// SecondaryMerchantSimpleCreate 二级商户标准进件 https://opendocs.alipay.com/solution/9434dd99_ant.merchant.expand.indirect.zft.simplecreate?scene=bf5951260023430e944c2e9afdf7f9e2&pathHash=5e709b90
type SecondaryMerchantSimpleCreate struct {
	AppAuthToken string `json:"-"` // 可选
	// external_id必选string(64)
	//【描述】商户编号，由一级商户定义，保证在一级商户下唯一即可
	ExternalId string `json:"external_id,omitempty"`
	// alias_name必选string(512)
	//【描述】商户别名。支付宝收银台及账单中的商户名称会展示此处设置的别名。如果涉及支付宝APP内的支付，支付结果页也会展示该别名；如果涉及线下当面付场景，请填写线下店铺名称
	AliasName string `json:"alias_name,omitempty"`
	// contact_infos必选ContactInfo
	//【描述】商户联系人信息。在本业务中，ContactInfo对象中联系人姓名、手机号必填，其他选填
	ContactInfos *ContactInfo `json:"contact_infos,omitempty"`
	// default_settle_rule必选DefaultSettleRule
	//【描述】默认结算规则。当调用收单接口，settle_info中设置默认结算规则（defaultSettle）时，交易资金将结算至此处设置的默认结算目标账户中。其详细描述及收单接口传参示例参考功能包文档
	DefaultSettleRule *DefaultSettleRule `json:"default_settle_rule,omitempty"`
	// service必选string[](200)
	//【描述】商户使用服务，可选值有：当面付、jsapi支付、app支付、wap支付、电脑支付、预授权支付、商户代扣、小程序支付、订单码支付。其值会影响其他字段必填性，详见其他字段描述
	//【枚举值】
	//当面付: 当面付
	//jsapi支付: jsapi支付
	//app支付: app支付
	//wap支付: wap支付
	//电脑支付: 电脑支付
	//预授权支付: 预授权支付
	//商户代扣: 商户代扣
	//小程序支付: 小程序支付
	//订单码支付: 订单码支付
	Service []string `json:"service,omitempty"`
	// mcc必选string(10)
	//【描述】商户类别码 mcc，可查看 进件MCC与资质要求 202212.xlsx，特殊行业要按照MCC说明中的资质一栏上传辅助资质，辅助资质要在 qualifications 中上传，会有人工审核。
	Mcc string `json:"mcc"`
	// info_source_uidstring(16)
	//【描述】（已废弃，请使用info_source_open_id）。信息关联的uid
	InfoSourceUid string `json:"info_source_uid,omitempty"`
	// info_source_open_idstring(128)
	//【描述】（平替原来的info_source_uid字段，如果能拿到openId，请传本字段，原字段留空）。信息关联的openId  详情可查看 openid简介 https://opendocs.alipay.com/mini/0ai2i6?pathHash=13dd5946
	InfoSourceOpenId string `json:"info_source_open_id,omitempty"`
	// oversea_settle_accountstring(32)
	//【描述】（已废弃，请使用oversea_settle_open_id）境外结算账号
	OverseaSettleAccount string `json:"oversea_settle_account,omitempty"`
	// oversea_settle_open_idstring(128)
	//【描述】（平替原来的oversea_settle_open_id字段，如能够获取到该场景的open_id，请传本字段，原字段留空）境外结算账号  详情可查看 openid简介 https://opendocs.alipay.com/mini/0ai2i6?pathHash=13dd5946
	OverseaSettleOpenId string `json:"oversea_settle_open_id,omitempty"`
	// alipay_logon_id特殊可选string(64)
	//【描述】结算支付宝账号，结算账号使用支付宝账号时必填，本字段指定交易资金结算的具体支付宝账号，与binding_alipay_logon_id同主体的支付宝账号即可
	AlipayLogonId string `json:"alipay_logon_id,omitempty"`
	// biz_cards特殊可选SettleCardInfo
	//【描述】结算银行卡信息，结算账号使用银行卡时必填。本业务当前只允许传入一张结算卡。个人类型商户不允许结算到银行卡
	//【必选条件】结算银行卡信息，结算账号使用银行卡时必填。本业务当前只允许传入一张结算卡。个人类型商户不允许结算到银行卡
	BizCards []*SettleCardInfo `json:"biz_cards,omitempty"`
	// business_address特殊可选AddressInfo
	//【描述】经营地址。使用当面付服务时必填。地址对象中省、市、区、地址必填，其余选填
	BusinessAddress *AddressInfo `json:"business_address,omitempty"`
	// out_door_images特殊可选string(256)
	//【描述】门头照，使用当面付服务时必填。其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。
	OutDoorImages []string `json:"out_door_images,omitempty"`
	// in_door_images特殊可选string(256)
	//【描述】内景照，使用当面付服务时必填。其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。
	InDoorImages []string `json:"in_door_images,omitempty"`
	// binding_alipay_logon_id可选string(100)
	//【描述】签约支付宝账户。需使用实名认证支付宝账号，使用该支付宝账号签约二级商户及后续服务，商户主体与该支付宝账号主体相同
	BindingAlipayLogonId string `json:"binding_alipay_logon_id,omitempty"`
	// service_phone可选string(20)
	//【描述】商户客服电话
	ServicePhone string `json:"service_phone,omitempty"`
	// name可选string(128)
	//【描述】进件的二级商户名称。一般情况下要与证件的名称相同。个体工商户类型可以放宽到法人名称
	Name string `json:"name,omitempty"`
	// license_auth_letter_image可选string(256)
	//【描述】授权函。当商户名与结算卡户名不一致。《说明函》模板参考。涉及外籍法人（这种情况上传任意能证明身份的图片）时必填，其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。（商户类型为个体工商户时，本字段仅需上传营业执照非授权函）
	LicenseAuthLetterImage string `json:"license_auth_letter_image,omitempty"`
	// invoice_info可选MerchantInvoiceInfo
	//【描述】开票资料信息
	InvoiceInfo *MerchantInvoiceInfo `json:"invoice_info,omitempty"`
	// sites可选SiteInfo
	//【描述】商户站点信息，包括网站、app、小程序。商户使用服务包含电脑支付、wap支付时，必须填充一个类型为01(网站)的SiteInfo对象，site_type/site_url/site_name必填；当包含app支付时，必须至少填充类型为02(APP)或06(支付宝小程序)中一种类型的SiteInfo对象，site_type/site_name必填；当包含jsapi支付时，必须填充一个类型为06(支付宝小程序)的SiteInfo对象；
	Sites *SiteInfo `json:"sites,omitempty"`
	// qualifications可选IndustryQualificationInfo[]
	//【描述】商户行业资质图片，当商户的经营类目选择了特殊行业时该字段必填，需要特殊行业资质文件。每项行业资质信息中，industry_qualification_type和industry_qualification_image均必填。
	Qualifications []*IndustryQualificationInfo `json:"qualifications,omitempty"`
	// additional_cert_no可选string(20)
	//【描述】补充证件号，与additional_cert_type+additional_cert_image同时提供。当商户类型为个人时，使用当面付收款有限额，补充这组证件信息可提额。目前仅允许个人类型商户传入本字段。
	AdditionalCertNo string `json:"additional_cert_no,omitempty"`
	// additional_cert_type可选string(20)
	//【描述】补充证件类型，与additional_cert_no+additional_cert_image同时提供。可选值有201（营业执照号）、204（民办非企业登记证书）、206（社会团体法人登记证书）、218（事业单位法人证书）、219（党政机关批准设立文件/行政执法主体资格证）。当商户类型为个人时，使用当面付收款有限额，补充这组证件信息可提额。目前仅允许个人类型商户传入本字段。
	//【枚举值】
	//营业执照: 201
	//事业单位法人证书: 218
	//民办非企业登记证书: 204
	//社会团体法人登记证书: 206
	//党政机关批准设立文件/行政执法主体资格证: 219
	AdditionalCertType string `json:"additional_cert_type,omitempty"`
	// additional_cert_image可选string(256)
	//【描述】补充证件图片，与additional_cert_no、additional_cert_type同时提供。当商户类型为个人时，使用当面付收款有限额，补充这组证件信息可提额。目前仅允许个人类型商户传入。其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。
	AdditionalCertImage string `json:"additional_cert_image,omitempty"`
	// sign_time_with_isv可选string(20)
	//【描述】二级商户与服务商的签约时间。
	SignTimeWithIsv string `json:"sign_time_with_isv,omitempty"`
	// zft_withholding_info可选ZFTWithholdingInfo
	//【描述】代扣产品信息
	ZftWithholdingInfo *ZFTWithholdingInfo `json:"zft_withholding_info,omitempty"`
	// trade_scene可选string[](32)
	//【描述】交易场景；
	//【枚举值】
	//小程序支付场景: TINY_APP
	//H5场景: WAP
	//线下当面付场景: OFFLINE
	//APP支付场景: APP
	//网站支付场景: PC
	TradeScene string `json:"trade_scene,omitempty"`
}

func (s SecondaryMerchantSimpleCreate) APIName() string {
	return "ant.merchant.expand.indirect.zft.simplecreate"
}

func (s SecondaryMerchantSimpleCreate) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = s.AppAuthToken
	return m
}

type SecondaryMerchantSimpleCreateRsp struct {
	Error
	// order_id必选string(64)
	//【描述】申请单id
	OrderId string `json:"order_id"`
}

// SecondaryMerchantSimpleQuery 二级商户入驻进度查询 https://opendocs.alipay.com/solution/dfebf71b_ant.merchant.expand.indirect.zftorder.query?scene=common&pathHash=019bc47e
type SecondaryMerchantSimpleQuery struct {
	AppAuthToken string `json:"-"`           // 可选
	OrderId      string `json:"order_id"`    // 申请单id
	ExternalId   string `json:"external_id"` // 进件申请时的外部商户id，与order_id二选一必填
}

func (s SecondaryMerchantSimpleQuery) APIName() string {
	return "ant.merchant.expand.indirect.zftorder.query"
}

func (s SecondaryMerchantSimpleQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = s.AppAuthToken
	return m
}

type SecondaryMerchantSimpleQueryRsp struct {
	Error
	// orders必选ZftSubMerchantOrder[]
	//【描述】直付通二级商户进件申请单信息
	Orders []*ZftSubMerchantOrder `json:"orders"`
}

type ZftSubMerchantOrder struct {
	// order_id必选string(32)
	//【描述】申请单id
	OrderId string `json:"order_id,omitempty"`
	// external_id必选string(32)
	//【描述】外部商户id
	ExternalId string `json:"external_id,omitempty"`
	// apply_time必选date(20)
	//【描述】申请单创建时间
	ApplyTime string `json:"apply_time,omitempty"`
	// apply_type必选string(64)
	//【描述】本申请单的请求类型。
	//【枚举值】
	//商户预校验: ZHIFUTONG_CONSULT
	//商户创建: ZHIFUTONG_CREATE
	//商户修改: ZHIFUTONG_MODIFY
	ApplyType string `json:"apply_type,omitempty"`
	// merchant_name可选string(256)
	//【描述】进件时填写的商户名称
	MerchantName string `json:"merchant_name,omitempty"`
	// status可选string(12)
	//【描述】申请总体状态。99:已完结;-1:失败;031:审核中
	Status string `json:"status,omitempty"`
	// fk_audit可选string(10)
	//【描述】风控审核状态。CREATE：已创建待审批、SKIP：跳过风控审批步骤、PASS：风控审核通过、REJECT：风控审批拒绝
	FkAudit string `json:"fk_audit,omitempty"`
	// fk_audit_memo可选string(64)
	//【描述】风控审批备注，如有则返回
	FkAuditMemo string `json:"fk_audit_memo,omitempty"`
	// kz_audit可选string(64)
	//【描述】客资审核状态。CREATE：已创建待审批、SKIP：跳过客资审批步骤、PASS：客资审核通过、REJECT：客资审批拒绝
	KzAudit string `json:"kz_audit,omitempty"`
	// kz_audit_memo可选string(64)
	//【描述】客资审批备注，如有则返回
	KzAuditMemo string `json:"kz_audit_memo,omitempty"`
	// sub_confirm可选string(10)
	//【描述】二级商户确认状态。CREATE：已发起二级商户确认、SKIP：无需确认、FAIL：签约失败、NOT_CONFIRM：商户未确认、FINISH签约完成
	SubConfirm string `json:"sub_confirm,omitempty"`
	// card_alias_no可选string(32)
	//【描述】进件生成的卡编号，在发起结算时可以作为结算账号
	CardAliasNo string `json:"card_alias_no,omitempty"`
	// safe_binding_logon_id｜签约支付宝账号可选string(128)
	//【描述】签约支付宝账号（脱敏）
	SafeBindingLogonId string `json:"safe_binding_logon_id,omitempty"`
	// smid可选string(32)
	//【描述】二级商户id。当总体申请状态status为99时，smid才算进件完成
	Smid string `json:"smid,omitempty"`
	// app_pre_auth可选string(5)
	//【描述】是否开通线上预授权
	AppPreAuth string `json:"app_pre_auth,omitempty"`
	// face_pre_auth可选string(5)
	//【描述】是否开通线下预授权
	FacePreAuth string `json:"face_pre_auth,omitempty"`
	// is_face_limit可选string(5)
	//【描述】判断个人当面付权限版本，返回TRUE时表示是标准版，返回FALSE表示受限版
	IsFaceLimit string `json:"is_face_limit,omitempty"`
	// reason可选string(256)
	//【描述】申请单处理失败时，通过此此段返回具体的失败理由；与kf_audit_memo和kz_audit_memo配合使用
	Reason string `json:"reason,omitempty"`
	// sub_sign_qr_code_url可选string(256)
	//【描述】在快速进件场景下要求二级商户自助签约确认，通过此链接展示的二维码扫码进入支付宝小程序。（仅quickcreate接口返回）
	SubSignQrCodeUrl string `json:"sub_sign_qr_code_url,omitempty"`
	// sub_sign_short_chain_url可选string(256)
	//【描述】页面跳转的短链接，用于二级商户完成签约确认，可以支持平台商发送此短链接给二级商户，通过链接跳转的方式进入支付宝小程序完成确认。（仅quickcreate接口返回）
	SubSignShortChainUrl string `json:"sub_sign_short_chain_url,omitempty"`
}

// SecondaryMerchantSimpleCancel 二级商户作废 https://opendocs.alipay.com/solution/a4111a84_ant.merchant.expand.indirect.zft.delete?scene=common&pathHash=1abe3619
type SecondaryMerchantSimpleCancel struct {
	AppAuthToken string `json:"-"`
	// smid必选string[0,99]
	//【描述】二级商户smid
	Smid string `json:"smid"`
}

func (s SecondaryMerchantSimpleCancel) APIName() string {
	return "ant.merchant.expand.indirect.zft.delete"
}

func (s SecondaryMerchantSimpleCancel) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = s.AppAuthToken
	return m
}

type SecondaryMerchantSimpleCancelRsp struct {
	Error
}
