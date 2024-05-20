package payment

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestPlaceAnOrder(t *testing.T) {
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

	if err != nil {
		t.Errorf("请求失败: %s", err.Error())
	}

	t.Logf("返回结果: %+v", res)

}

func TestPayment_RenewOrder(t *testing.T) {
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
		VspIds:            []string{"122", "123"},
	}

	res, err := payCenter.RenewOrder()

	if err != nil {
		t.Errorf("请求失败: %s", err.Error())
	}

	t.Logf("返回结果: %+v", res)
}
