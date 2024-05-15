package payment

import "github.com/shopspring/decimal"

const (
	CreateOrderUri = "/sweep/createOrder"
	RenewOrderUri  = "/sweep/renewOrder"
)

type Payment struct {
	UserId            uint
	UserName          string
	ProductType       uint32
	ProductSType      string
	ProductName       string
	OrderAmount       decimal.Decimal
	PaidAmount        decimal.Decimal
	Ipaddress         string
	Cpu               uint32
	Memory            uint32
	Bandwidth         uint32
	HardDisks         uint32
	Disks             uint32
	Months            uint32
	InstanceName      string
	InstanceID        string
	Version           string
	Type              string
	APIUriPrefix      string
	APISignSecret     string
	CallBackUrl       string
	CurrentExpireTime string // 当前过期时间，字符串格式：YYYY-MM-DD HH:ii:ss
	RenewExpireTime   string // 续费后过期时间，字符串格式：YYYY-MM-DD HH:ii:ss
}

// setAPIUrlPrefix 设置api schema和host，默认是官网st环境
func (p *Payment) setAPIUrlPrefix() {
	if p.APIUriPrefix == "" {
		p.APIUriPrefix = "https://www.st.landui.cn"
	}
}
