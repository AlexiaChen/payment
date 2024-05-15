package payment

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gitlab.landui.cn/gomod/global"
	"go.uber.org/zap"
	"testing"
)

func TestPlaceAnOrder(t *testing.T) {
	initLog()
	payCenter := Payment{
		UserId:        14610,
		UserName:      "86326328",
		ProductType:   1,
		ProductSType:  "cloud_other",
		ProductName:   "云数据库 MySQL",
		OrderAmount:   decimal.NewFromFloat(123.00),
		PaidAmount:    decimal.NewFromFloat(123.00),
		Ipaddress:     "10.99.255.73",
		Cpu:           1,
		Memory:        2,
		Bandwidth:     3000,
		HardDisks:     40,
		Disks:         10,
		Months:        1,
		InstanceName:  "st001",
		Version:       "5.7.43",
		Type:          "独享型",
		APIUriPrefix:  "https://www.st.landui.cn",
		APISignSecret: "JMY13PagXnhl3rpiI1ht1hBBOaSF7dSOf8ktJ95zmOx19PWayRlyCtCm7UT0mghJ",
		CallBackUrl:   "https://rds.st.landui.cn/auth/payment/callback",
	}

	res, err := payCenter.PlaceAnOrder()

	fmt.Println(res)
	fmt.Println(err)

}

func TestPayment_RenewOrder(t *testing.T) {
	initLog()
	payCenter := Payment{
		UserId:            14610,
		UserName:          "86326328",
		ProductType:       1,
		ProductSType:      "cloud_product_renew",
		ProductName:       "云数据库 MySQL",
		OrderAmount:       decimal.NewFromFloat(123.00),
		PaidAmount:        decimal.NewFromFloat(123.00),
		Ipaddress:         "10.99.255.73",
		Cpu:               1,
		Memory:            2,
		Bandwidth:         3000,
		HardDisks:         40,
		Disks:             10,
		Months:            1,
		InstanceName:      "st002",
		Version:           "5.7.43",
		Type:              "独享型",
		APIUriPrefix:      "http://www.lxy.dev.landui.cn",
		APISignSecret:     "JMY13PagXnhl3rpiI1ht1hBBOaSF7dSOf8ktJ95zmOx19PWayRlyCtCm7UT0mghJ",
		CallBackUrl:       "https://rds.st.landui.cn/auth/payment/callback",
		CurrentExpireTime: "2024-05-01 10:00:00",
		RenewExpireTime:   "2024-09-01 10:00:00",
	}

	res, err := payCenter.RenewOrder()

	fmt.Println(res)
	fmt.Println(err)
}

func initLog() {
	global.Logger = zap.NewExample() // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
}
