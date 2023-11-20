package controller

import (
	"github.com/Moonyongjung/xns-gateway/request"
	"github.com/Moonyongjung/xns-gateway/service"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	xutil "github.com/xpladev/xpla.go/util"
)

// QueryBalance godoc
// @Summary      Balance of the controller
// @Description  Get balance of the controller contract
// @Tags         Controller
// @Accept       json
// @Produce      json
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /controller/balance [get]
func QueryBalance(xnsContext *types.XnsContext) (*service.QueryResponse, error) {
	msg, err := NewControllerControllerBalanceQueryMsg()
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryConfig godoc
// @Summary      Config of the controller
// @Description  Get the configuration of the controller contract
// @Tags         Controller
// @Accept       json
// @Produce      json
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /controller/config [get]
func QueryConfig(xnsContext *types.XnsContext) (*service.QueryResponse, error) {
	msg, err := NewControllerGetConfigQueryMsg()
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryRegisterCost godoc
// @Summary      Calculate cost amount
// @Description  Calculate cost to register domain
// @Tags         Controller
// @Accept       json
// @Produce      json
// @Param        label    path      string  true  "domain name for register"
// @Param        duration path      string  true  "expire duration"
// @Success      200      {object}  service.QueryResponse
// @Failure      400      {object}  service.QueryResponse
// @Router       /controller/register-cost/{label}/{duration} [get]
func QueryRegisterCost(xnsContext *types.XnsContext, label, duration string) (*service.QueryResponse, error) {
	durationU64, err := xutil.FromStringToUint64(duration)
	if err != nil {
		return nil, util.LogErr(types.ErrParseData, err)
	}

	msg, err := NewControllerRegisterCostQueryMsg(label, durationU64)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryTopDomain godoc
// @Summary      Get top domain
// @Description  Get top domain of the controller
// @Tags         Controller
// @Accept       json
// @Produce      json
// @Success      200      {object}  service.QueryResponse
// @Failure      400      {object}  service.QueryResponse
// @Router       /controller/top-domain [get]
func QueryTopDomain(xnsContext *types.XnsContext) (*service.QueryResponse, error) {
	msg, err := NewControllerTopDomainQueryMsg()
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

func Query(xnsContext *types.XnsContext, msg string) (*service.QueryResponse, error) {
	address := xnsContext.GetApp().Contracts.ControllerContract.ContractAddress
	res, err := request.QueryContract(xnsContext, address, msg)
	if err != nil {
		return nil, err
	}

	return service.ParseResponse(res)
}
