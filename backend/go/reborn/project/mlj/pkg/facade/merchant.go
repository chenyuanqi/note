package facade

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"mlj/pkg/common/helpers"
)

const BaseURL = "https://madminv2.mlj130.com"

type MerchantResponse struct {
	Result  string          `json:"result"`
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type MerchantItemResponse struct {
	Count string              `json:"count"`
	List  []*MerchantItemInfo `json:"list"`
}

type MerchantItemInfo struct {
	ItemID         string `json:"item_id"`
	OriginPrice    string `json:"origin_price"`
	Price          string `json:"price"`
	CommissionRate string `json:"commission_rate"`
	CouponUrl      string `json:"coupon_url"`
	CouponAmount   string `json:"coupon_amount"`
	CouponStart    string `json:"coupon_start"`
	CouponEnd      string `json:"coupon_end"`
	CouponID       string `json:"coupon_id"`
}

func GetMerchantCouponList(page, pageSize int) (*MerchantItemResponse, error) {
	requestURL := BaseURL + "/api/order/coupon-list"
	params := map[string]string{
		"page":     strconv.Itoa(page),
		"per-page": strconv.Itoa(pageSize),
	}
	res, err := helpers.Get(requestURL, params)
	if err != nil {
		return nil, err
	}

	var r *MerchantResponse
	if err := json.Unmarshal(res, &r); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if r.Result == "success" {
		var data *MerchantItemResponse
		if err := json.Unmarshal(r.Data, &data); err != nil {
			fmt.Println(err)
			return nil, err
		}
		return data, nil
	}
	return nil, errors.New(r.Message)
}
