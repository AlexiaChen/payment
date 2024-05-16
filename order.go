package payment

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/guid"
	"gitlab.landui.cn/gomod/logs"
	"sort"
	"time"
)

type Response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Info    *Info  `json:"info"`
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

type renewOrderParam struct {
	Stype             string   `json:"stype"`
	ProductType       uint32   `json:"productType"`
	Name              string   `json:"name"`
	InstanceName      string   `json:"instanceName"`
	InstanceID        string   `json:"InstanceID"`
	VpsIds            []string `json:"vpsIds"`
	Price             string   `json:"price"`
	BuyTime           string   `json:"buyTime"`
	CurrentExpireTime string   `json:"currentExpireTime"`
	RenewExpireTime   string   `json:"renewExpireTime"`
	CallBack          string   `json:"callback"`
	OrderDisposition  string   `json:"orderDisposition"`
}

// PlaceAnOrder 购买时创建订单
func (p *Payment) PlaceAnOrder() (*Response, error) {
	p.setAPIUrlPrefix()
	times := time.Now().Unix()
	randStr := guid.S()
	text := fmt.Sprintf("%d%s%s", times, randStr, p.APISignSecret)
	newText := sorts(text)
	ciphertext := gmd5.MustEncryptString(newText)
	var orderParams []placeOrderParam
	orderParam := placeOrderParam{
		Stype:       "cloud_other",
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
	}
	orderParams = append(orderParams, orderParam)
	param, _ := json.Marshal(orderParams)
	body := map[string]interface{}{
		"time_stamp":   fmt.Sprintf("%d", times),
		"nonce_str":    randStr,
		"sign":         ciphertext,
		"userid":       p.UserId,
		"username":     p.UserName,
		"ordermoney":   p.OrderAmount,
		"checkmoney":   p.PaidAmount,
		"param":        string(param),
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

// RenewOrder 续费时创建订单
func (p *Payment) RenewOrder() (*Response, error) {
	p.setAPIUrlPrefix()
	if p.RenewExpireTime == "" || p.CurrentExpireTime == "" {
		return nil, errors.New("RenewExpireTime or CurrentExpireTime cant be null")
	}
	layout := "2006-01-02 15:04:05"
	renewTime, err := time.Parse(layout, p.RenewExpireTime)
	if err != nil {
		return nil, errors.New("RenewExpireTime invalid")
	}
	expireTime, err := time.Parse(layout, p.CurrentExpireTime)
	if err != nil {
		return nil, errors.New("CurrentExpireTime invalid")
	}
	if expireTime.After(renewTime) || expireTime.Equal(renewTime) {
		return nil, errors.New("RenewExpireTime invalid")
	}

	times := time.Now().Unix()
	randStr := guid.S()
	text := fmt.Sprintf("%d%s%s", times, randStr, p.APISignSecret)
	newText := sorts(text)
	ciphertext := gmd5.MustEncryptString(newText)
	var orderParams []renewOrderParam
	orderParam := renewOrderParam{
		Stype:             "cloud_product_renew",
		ProductType:       p.ProductType,
		Name:              p.ProductName,
		InstanceID:        p.InstanceID,
		InstanceName:      p.InstanceName,
		VpsIds:            p.VspIds,
		Price:             p.PaidAmount.String(),
		BuyTime:           fmt.Sprintf("%d个月", p.Months),
		CallBack:          p.CallBackUrl,
		CurrentExpireTime: p.CurrentExpireTime,
		RenewExpireTime:   p.RenewExpireTime,
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
	}
	orderParams = append(orderParams, orderParam)
	param, _ := json.Marshal(orderParams)
	body := map[string]interface{}{
		"time_stamp":   fmt.Sprintf("%d", times),
		"nonce_str":    randStr,
		"sign":         ciphertext,
		"userid":       p.UserId,
		"username":     p.UserName,
		"ordermoney":   p.OrderAmount,
		"checkmoney":   p.PaidAmount,
		"param":        string(param),
		"ipaddress":    p.Ipaddress,
		"voucher_num":  "0",
		"coupon_id":    "0",
		"couponsmoney": "0",
		"source":       "pc端",
	}
	httpClient := resty.New()
	var resp Response
	logs.New().SetAdditionalInfo("body", body).SetAdditionalInfo("url", p.APIUriPrefix+RenewOrderUri).Info("记录下单的参数")
	bodyJson, _ := json.Marshal(body)
	fmt.Println(string(bodyJson))
	res, err := httpClient.R().SetBody(body).SetResult(&resp).Post(p.APIUriPrefix + RenewOrderUri)
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
