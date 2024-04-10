package payment

// setAPIUrlPrefix 设置api schema和host，默认是官网st环境
func (p *Payment) setAPIUrlPrefix() {
	if p.APIUriPrefix == "" {
		p.APIUriPrefix = "https://www.st.landui.cn"
	}
}
