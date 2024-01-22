package contract

import (
	xtypes "github.com/xpladev/xpla.go/types"
)

type StoreResponseInfo struct {
	Contract string
	CodeID   string
}

func NewStoreResponseInfo(contract string, response *xtypes.TxRes) StoreResponseInfo {
	for _, attr := range response.Response.Logs[0].Events[1].Attributes {
		if attr.Key == "code_id" {
			return StoreResponseInfo{
				Contract: contract,
				CodeID:   attr.Value,
			}
		}
	}
	return StoreResponseInfo{}
}

type InstantiateResponseInfo struct {
	Contract string
	Address  string
}

func NewInstantiateResponseInfo(contract string, response *xtypes.TxRes) InstantiateResponseInfo {
	for _, attr := range response.Response.Logs[0].Events[0].Attributes {
		if attr.Key == "_contract_address" {
			return InstantiateResponseInfo{
				Contract: contract,
				Address:  attr.Value,
			}
		}
	}
	return InstantiateResponseInfo{}
}
