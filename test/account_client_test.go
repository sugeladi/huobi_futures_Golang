package test

import (
	"github.com/sugeladi/huobi_futures_Golang/sdk/linearswap/restful"
	"github.com/sugeladi/huobi_futures_Golang/sdk/linearswap/restful/response/account"
	"testing"
)

var acClient restful.AccountClient

func init() {
	acClient = restful.AccountClient{}
	acClient.Init(config.AccessKey, config.SecretKey, config.Host)
}

func TestAccountClient_GetAccountAssetsAsync(t *testing.T) {

	data := make(chan account.GetAccountAssetsResponse)

	go acClient.GetAccountAssetsAsync(data, "BTC-USDT", 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go acClient.GetAccountAssetsAsync(data, "BTC-USDT", config.SubUid)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetAccountPositionAsync(t *testing.T) {

	data := make(chan account.GetAccountPositionResponse)

	go acClient.GetAccountPositionAsync(data, "BTC-USDT", 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go acClient.GetAccountPositionAsync(data, "BTC-USDT", config.SubUid)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetAllSubAssetsAsync(t *testing.T) {

	data := make(chan account.GetAllSubAssetsResponse)

	go acClient.GetAllSubAssetsAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_AccountTransferAsync(t *testing.T) {

	data := make(chan account.AccountTransferResponse)

	go acClient.AccountTransferAsync(data, "USDT", "BTC-USDT", "ETH-USDT", 1, config.SubUid, "master_to_sub")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go acClient.AccountTransferAsync(data, "USDT", "BTC-USDT", "ETH-USDT", 1, 0, "")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetAccountTransHisAsync(t *testing.T) {

	data := make(chan account.GetAccountTransHisResponse)

	go acClient.GetAccountTransHisAsync(data, "BTC-USDT", false, "34", 10, 1, 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go acClient.GetAccountTransHisAsync(data, "BTC-USDT", true, "34", 10, 1, 10)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetValidLeverRateAsync(t *testing.T) {

	data := make(chan account.GetValidLeverRateResponse)

	go acClient.GetValidLeverRateAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetOrderLimitAsync(t *testing.T) {

	data := make(chan account.GetOrderLimitResponse)

	go acClient.GetOrderLimitAsync(data, "limit", "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetFeeAsync(t *testing.T) {

	data := make(chan account.GetFeeResponse)

	go acClient.GetFeeAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetTransferLimitAsync(t *testing.T) {

	data := make(chan account.GetTransferLimitResponse)

	go acClient.GetTransferLimitAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetPositionLimitAsync(t *testing.T) {

	data := make(chan account.GetPositionLimitResponse)

	go acClient.GetPositionLimitAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetApiTradingStatusAsync(t *testing.T) {

	data := make(chan account.GetApiTradingStatusResponse)

	go acClient.GetApiTradingStatusAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}
