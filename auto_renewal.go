package payment

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/guid"
)

type AutoRenewalPayment struct {
	apiUrlPrefix  string
	apiSignSecret string
}

func MakeAutoRenewalPayment(apiUrlPrefix, apiSignSecret string) *AutoRenewalPayment {
	if apiUrlPrefix == "" {
		panic("apiUrlPrefix is required")
	}
	if apiSignSecret == "" {
		panic("apiSignSecret is required")
	}

	return &AutoRenewalPayment{
		apiUrlPrefix:  apiUrlPrefix,
		apiSignSecret: apiSignSecret,
	}
}

type CommonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type SuccessResponse struct {
	CommonResponse
	Info []interface{} `json:"info"`
}

type ErrorResponse struct {
	CommonResponse
	Info struct {
		ErrorCode string `json:"error_code"`
	} `json:"info"`
}

// RenewalVPS 续费VPS
// username: 用户名
// vpsId: VPS ID
// months: 续费月数
// price: 价格
// endTime: 到期时间（Unix时间戳
// https://yapi.landui.cn/project/37/interface/api/4127
func (s *AutoRenewalPayment) RenewalVPS(username string, vpsId string, months uint32, price string, endTime int64) (string, error) {
	if username == "" {
		return "", errors.New("username is required")
	}
	if vpsId == "" {
		return "", errors.New("vpsId is required")
	}
	if months == 0 {
		return "", errors.New("months is required")
	}
	if price == "" {
		return "", errors.New("price is required")
	}
	if endTime == 0 {
		return "", errors.New("endTime is required")
	}

	// 创建resty客户端
	client := resty.New()

	times := time.Now().Unix()
	randStr := guid.S()
	text := fmt.Sprintf("%d%s%s", times, randStr, s.apiSignSecret)
	newText := sorts(text)
	ciphertext := gmd5.MustEncryptString(newText)

	// 发送表单请求
	resp, err := client.R().
		SetBody(map[string]string{
			"time_stamp": fmt.Sprintf("%d", times),
			"nonce_str":  randStr,
			"sign":       ciphertext,
			"username":   username,
			"vps_id":     vpsId,
			"months":     fmt.Sprintf("%d", months),
			"price":      price,
			"end_time":   fmt.Sprintf("%d", endTime),
		}).
		Post(s.apiUrlPrefix + "/api/autoRenew")

	if err != nil {
		return "", fmt.Errorf("resty HTTP request failed: %v", err)
	}

	if resp.StatusCode() != 200 {
		return resp.String(), errors.New("HTTP status code is not 200")
	}

	var commonRespone CommonResponse

	err = json.Unmarshal(resp.Body(), &commonRespone)
	if err != nil {
		return string(resp.Body()), fmt.Errorf("json.Unmarshal common response failed: %v", err)
	}

	if commonRespone.Status == "y" {
		var successResponse SuccessResponse
		err = json.Unmarshal(resp.Body(), &successResponse)
		if err != nil {
			return string(resp.Body()), fmt.Errorf("json.Unmarshal success response failed: %v", err)
		}
		return string(resp.Body()), nil
	} else {
		var errorResponse ErrorResponse
		err = json.Unmarshal(resp.Body(), &errorResponse)
		if err != nil {
			return string(resp.Body()), fmt.Errorf("json.Unmarshal error response failed: %v", err)
		}
		return string(resp.Body()), fmt.Errorf("renewal failed, error_code: %s msg: %s code: %d",
			errorResponse.Info.ErrorCode, errorResponse.Message, errorResponse.Code)
	}

}
