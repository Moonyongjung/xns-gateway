package registrar

import (
	"github.com/Moonyongjung/xns-gateway/request"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	xtypes "github.com/xpladev/xpla.go/types"
)

func ExecuteRegistrar(xnsContext *types.XnsContext, msg, amount string) (*xtypes.TxRes, error) {
	address := xnsContext.GetApp().Contracts.RegistrarContract.ContractAddress
	return request.ExecuteContract(xnsContext, address, msg, amount)
}

func ExecSetController(xnsContext *types.XnsContext, topDomain string) error {
	controllerAddress := xnsContext.GetApp().Contracts.ControllerContract.ContractAddress
	if controllerAddress == "" {
		return util.LogErr(types.ErrHaveNoAddress, "must have controller contract address when execute set controller")
	}

	msg, err := NewRegistrarSetControllerExecuteMsg(controllerAddress, topDomain)
	if err != nil {
		return util.LogErr(types.ErrNewMsg, err)
	}

	res, err := ExecuteRegistrar(xnsContext, msg, types.ZeroAmount)
	if err != nil {
		return err
	}

	util.LogInfo(res.Response)

	return nil
}

func ExecSetResolver(xnsContext *types.XnsContext) error {
	resolverAddress := xnsContext.GetApp().Contracts.ResolverContract.ContractAddress
	if resolverAddress == "" {
		return util.LogErr(types.ErrHaveNoAddress, "must have controller contract address when execute set controller")
	}

	msg, err := NewRegistrarSetResolverExecuteMsg(resolverAddress)
	if err != nil {
		return util.LogErr(types.ErrNewMsg, err)
	}

	res, err := ExecuteRegistrar(xnsContext, msg, types.ZeroAmount)
	if err != nil {
		return err
	}

	util.LogInfo(res.Response)

	return nil
}
