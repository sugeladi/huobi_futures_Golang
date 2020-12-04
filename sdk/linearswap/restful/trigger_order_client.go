﻿package restful

import (
	"encoding/json"
	"fmt"
	"github.com/sugeladi/huobi_futures_Golang/sdk/linearswap"
	requesttiggerorder "github.com/sugeladi/huobi_futures_Golang/sdk/linearswap/restful/request/triggerorder"
	responsetriggerorder "github.com/sugeladi/huobi_futures_Golang/sdk/linearswap/restful/response/triggerorder"
	"github.com/sugeladi/huobi_futures_Golang/sdk/log"
	"github.com/sugeladi/huobi_futures_Golang/sdk/reqbuilder"
)

type TriggerOrderClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (toc *TriggerOrderClient) Init(accessKey string, secretKey string, host string) *TriggerOrderClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	toc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return toc
}

func (toc *TriggerOrderClient) PlaceOrderAsync(data chan responsetriggerorder.PlaceOrderResponse, request requesttiggerorder.PlaceOrderRequest) {
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_trigger_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.PlaceOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to PlaceOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CancelOrderAsync(data chan responsetriggerorder.CancelOrderResponse, contractCode string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_trigger_cancel", nil)
	if orderId == "" {
		url = toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_trigger_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) GetOpenOrderAsync(data chan responsetriggerorder.GetOpenOrderResponse, contractCode string, pageIndex int, pageSize int) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_trigger_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) GetHisOrderAsync(data chan responsetriggerorder.GetHisOrderResponse, contractCode string, tradeType int, status string, createDate int,
	pageIndex int, pageSize int) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_trigger_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": %d,\"status\": \"%s\",\"create_date\": %d", contractCode, tradeType, status, createDate)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}
