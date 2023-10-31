package controller

import (
	"github.com/Moonyongjung/xns-gateway/request"
	"github.com/Moonyongjung/xns-gateway/types"

	xtypes "github.com/xpladev/xpla.go/types"
)

func ExecuteController(xnsContext *types.XnsContext, msg, amount string) (*xtypes.TxRes, error) {
	address := xnsContext.GetApp().Contracts.ControllerContract.ContractAddress
	return request.ExecuteContract(xnsContext, address, msg, amount)
}
