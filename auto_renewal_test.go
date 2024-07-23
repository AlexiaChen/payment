package payment

import (
	"testing"
	"time"
)

func TestAutoRenewalPayment_RenewalVPS(t *testing.T) {
	// Create a new instance of AutoRenewalPayment
	payment := MakeAutoRenewalPayment("http://www.fcr.dev.landui.cn", "JMY13PagXnhl3rpiI1ht1hBBOaSF7dSOf8ktJ95zmOx19PWayRlyCtCm7UT0mghJ")

	// Test case 1: Valid input
	username := "86326328"
	vpsId := "120"
	months := uint32(6)
	price := "10.99"
	endTime := time.Now().AddDate(0, 6, 0).Unix()

	response, err := payment.RenewalVPS(username, vpsId, months, price, endTime)

	if err != nil {
		t.Errorf("Request failed: %s respone: %s", err.Error(), response)
	} else {
		t.Logf("sucecess Response: %s", response)
	}

	username = "86326328"
	vpsId = "120"
	months = uint32(6)
	price = "10.99"
	endTime = time.Now().AddDate(0, 1, 0).Unix()

	response, err = payment.RenewalVPS(username, vpsId, months, price, endTime)

	if err != nil {
		t.Errorf("Request failed: %s respone: %s", err.Error(), response)
	}

}
