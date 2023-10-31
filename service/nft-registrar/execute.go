package nftregistrar

import (
	"github.com/Moonyongjung/xns-gateway/request"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	xtypes "github.com/xpladev/xpla.go/types"
)

func ExecuteNftRegistrar(xnsContext *types.XnsContext, msg, amount string) (*xtypes.TxRes, error) {
	address := xnsContext.GetApp().Contracts.NftRegistrarContract.ContractAddress
	return request.ExecuteContract(xnsContext, address, msg, amount)
}

func ExecSetRegistrar(xnsContext *types.XnsContext) error {
	registrarAddress := xnsContext.GetApp().Contracts.RegistrarContract.ContractAddress
	if registrarAddress == "" {
		return util.LogErr(types.ErrHaveNoAddress, "must have registrar contract address when execute set registrar")
	}

	msg, err := NewNftRegistrarSetRegistrarExecuteMsg(registrarAddress)
	if err != nil {
		return util.LogErr(types.ErrNewMsg, err)
	}

	res, err := ExecuteNftRegistrar(xnsContext, msg, types.ZeroAmount)
	if err != nil {
		return err
	}

	util.LogInfo(res.Response)

	return nil
}

func ExecSetResolver(xnsContext *types.XnsContext) error {
	resolverAddress := xnsContext.GetApp().Contracts.ResolverContract.ContractAddress
	if resolverAddress == "" {
		return util.LogErr(types.ErrHaveNoAddress, "must have resolver contract address when execute set resolver")
	}

	msg, err := NewNftRegistrarSetResolverExecuteMsg(resolverAddress)
	if err != nil {
		return util.LogErr(types.ErrNewMsg, err)
	}

	res, err := ExecuteNftRegistrar(xnsContext, msg, types.ZeroAmount)
	if err != nil {
		return err
	}

	util.LogInfo(res.Response)

	return nil
}
