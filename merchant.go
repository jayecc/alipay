package alipay

// SecondaryMerchantSimpleCreate 二级商户标准进件 https://opendocs.alipay.com/solution/9434dd99_ant.merchant.expand.indirect.zft.simplecreate?scene=bf5951260023430e944c2e9afdf7f9e2&pathHash=d3136936
func (this *Client) SecondaryMerchantSimpleCreate(param SecondaryMerchantSimpleCreate) (result *SecondaryMerchantSimpleCreateRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// SecondaryMerchantSimpleQuery https://opendocs.alipay.com/solution/dfebf71b_ant.merchant.expand.indirect.zftorder.query?scene=common&pathHash=019bc47e
func (this *Client) SecondaryMerchantSimpleQuery(param SecondaryMerchantSimpleQuery) (result *SecondaryMerchantSimpleQueryRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// SecondaryMerchantSimpleCancel 二级商户作废 https://opendocs.alipay.com/solution/a4111a84_ant.merchant.expand.indirect.zft.delete?scene=common&pathHash=1abe3619
func (this *Client) SecondaryMerchantSimpleCancel(param SecondaryMerchantSimpleCancel) (result *SecondaryMerchantSimpleCancelRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}
