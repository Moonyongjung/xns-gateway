package request

import (
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	xtypes "github.com/xpladev/xpla.go/types"
)

func ExecuteContract(xnsContext *types.XnsContext, address, msg, amount string) (*xtypes.TxRes, error) {
	xplac := xnsContext.GetXplaClient()
	executeMsg := xtypes.ExecuteMsg{
		ContractAddress: address,
		Amount:          amount,
		ExecMsg:         msg,
	}

	txBytes, err := xplac.ExecuteContract(executeMsg).CreateAndSignTx()
	if err != nil {
		return nil, util.LogErr(types.ErrTx, err)
	}

	res, err := xplac.Broadcast(txBytes)
	if err != nil {
		return nil, util.LogErr(types.ErrBroadcast, err)
	}

	return res, nil
}

func QueryContract(xnsContext *types.XnsContext, address, msg string) (string, error) {
	xplac := xnsContext.GetXplaClient()
	queryMsg := xtypes.QueryMsg{
		ContractAddress: address,
		QueryMsg:        msg,
	}

	res, err := xplac.QueryContract(queryMsg).Query()
	if err != nil {
		return "", util.LogErr(types.ErrQuery, err)
	}

	return res, nil
}
