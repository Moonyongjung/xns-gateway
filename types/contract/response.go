package contract

import (
	xtypes "github.com/xpladev/xpla.go/types"
)

type StoreResponseInfo struct {
	Contract string
	CodeID   string
}

func NewStoreResponseInfo(contract string, response *xtypes.TxRes) StoreResponseInfo {
	codeID := response.Response.Logs[0].Events[1].Attributes[0].Value
	return StoreResponseInfo{
		Contract: contract,
		CodeID:   codeID,
	}
}

type InstantiateResponseInfo struct {
	Contract string
	Address  string
}

func NewInstantiateResponseInfo(contract string, response *xtypes.TxRes) InstantiateResponseInfo {
	address := response.Response.Logs[0].Events[0].Attributes[0].Value
	return InstantiateResponseInfo{
		Contract: contract,
		Address:  address,
	}
}
