package payment

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/shopspring/decimal"
	"gitlab.landui.cn/gomod/logs"
	"sort"
	"time"
)

const (
	CreateOrderUri = "/sweep/createOrder"
)

type Payment struct {
	UserId        uint
	UserName      string
	ProductType   uint32
	ProductSType  string
	ProductName   string
	OrderAmount   decimal.Decimal
	PaidAmount    decimal.Decimal
	Ipaddress     string
	Cpu           uint32
	Memory        uint32
	Bandwidth     uint32
	HardDisks     uint32
	Disks         uint32
	Months        uint32
	InstanceName  string
	Version       string
	Type          string
	APIUriPrefix  string
	APISignSecret string
	CallBackUrl   string
}

type Response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Info    Info   `json:"info"`
}
type Info struct {
	Id string `json:"id"`
}

type placeOrderParam struct {
	Stype            string `json:"stype"`
	ProductType      uint32 `json:"productType"`
	Name             string `json:"name"`
	Price            string `json:"price"`
	BuyTime          string `json:"buyTime"`
	CallBack         string `json:"callback"`
	OrderDisposition string `json:"orderDisposition"`
}

func (p *Payment) PlaceAnOrder() (*Response, error) {
	p.setAPIUrlPrefix()
	times := time.Now().Unix()
	randStr := guid.S()
	text := fmt.Sprintf("%d%s%s", times, randStr, p.APISignSecret)
	newText := sorts(text)
	ciphertext := gmd5.MustEncryptString(newText)
	param, _ := json.Marshal(placeOrderParam{
		Stype:       p.ProductSType,
		ProductType: p.ProductType,
		Name:        p.ProductName,
		Price:       p.PaidAmount.String(),
		BuyTime:     fmt.Sprintf("%d个月", p.Months),
		CallBack:    p.CallBackUrl,
		OrderDisposition: fmt.Sprintf(
			"实例名称:%s|硬盘:%dG|版本:%s|IOPS:%d|CPU:%d核|内存:%dG|购买时长:%d个月|系列:%s",
			p.InstanceName,
			p.Disks,
			p.Version,
			p.Bandwidth,
			p.Cpu,
			p.Memory,
			p.Months,
			p.Type,
		),
	})
	body := map[string]interface{}{
		"time_stamp":   fmt.Sprintf("%d", times),
		"nonce_str":    randStr,
		"sign":         ciphertext,
		"userid":       p.UserId,
		"username":     p.UserName,
		"ordermoney":   p.OrderAmount,
		"checkmoney":   p.PaidAmount,
		"param":        param,
		"ipaddress":    p.Ipaddress,
		"voucher_num":  "0",
		"coupon_id":    "0",
		"couponsmoney": "0",
		"source":       "pc端",
	}
	httpClient := resty.New()
	var resp Response
	logs.New().SetAdditionalInfo("body", body).SetAdditionalInfo("url", p.APIUriPrefix+CreateOrderUri).Info("记录下单的参数")
	bodyJson, _ := json.Marshal(body)
	fmt.Println(string(bodyJson))
	res, err := httpClient.R().SetBody(body).SetResult(&resp).Post(p.APIUriPrefix + CreateOrderUri)
	if err != nil {
		fmt.Println("请求的body", body)
		fmt.Println(res.String())
		fmt.Println("请求失败", err)
		return nil, err
	}
	return &resp, nil
}

func sorts(text string) string {
	var array []string
	for _, v := range text {
		array = append(array, string(v))
	}
	sort.Strings(array)
	newText := ""
	for _, v := range array {
		newText += v
	}
	return newText
}
