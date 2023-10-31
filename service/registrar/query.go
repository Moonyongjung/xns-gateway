package registrar

import (
	"github.com/Moonyongjung/xns-gateway/cmd/flag"
	"github.com/Moonyongjung/xns-gateway/request"
	"github.com/Moonyongjung/xns-gateway/service"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"
)

// QueryConfig godoc
// @Summary      Config of the contract
// @Description  Get configuration of the registrar contract
// @Tags         Registrar
// @Accept       json
// @Produce      json
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /registrar/config [get]
func QueryConfig(xnsContext *types.XnsContext) (*service.QueryResponse, error) {
	msg, err := NewRegistrarGetConfigQueryMsg()
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryNftRegistrar godoc
// @Summary      linked NFT registrar
// @Description  Get linked the NFT registrar contract
// @Tags         Registrar
// @Accept       json
// @Produce      json
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /registrar/nft-registrar [get]
func QueryNftRegistrar(xnsContext *types.XnsContext) (*service.QueryResponse, error) {
	msg, err := NewRegistrarNftRegistrarAddressQueryMsg()
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryController godoc
// @Summary      linked controllers
// @Description  Get linked controller contracts
// @Tags         Registrar
// @Accept       json
// @Produce      json
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /registrar/controllers [get]
func QueryControllers(xnsContext *types.XnsContext, startAfter, limit string) (*service.QueryResponse, error) {
	msg, err := NewRegistrarControllersQueryMsg(flag.StartAfterI(startAfter), flag.LimitI(limit))
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

func Query(xnsContext *types.XnsContext, msg string) (*service.QueryResponse, error) {
	address := xnsContext.GetApp().Contracts.RegistrarContract.ContractAddress
	res, err := request.QueryContract(xnsContext, address, msg)
	if err != nil {
		return nil, err
	}

	return service.ParseResponse(res)
}
