package resolver

import (
	"github.com/Moonyongjung/xns-gateway/request"
	"github.com/Moonyongjung/xns-gateway/service"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"
)

// QueryConfig godoc
// @Summary      Config of the contract
// @Description  Get configuration of the resolver contract
// @Tags         Resolver
// @Accept       json
// @Produce      json
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /resolver/config [get]
func QueryConfig(xnsContext *types.XnsContext) (*service.QueryResponse, error) {
	msg, err := NewRegistrarGetConfigQueryMsg()
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryDomainState godoc
// @Summary      XNS state for the domain
// @Description  Get XNS state for the domain
// @Tags         Resolver
// @Accept       json
// @Produce      json
// @Param        hashed_label path      string  true  "hashed label"
// @Param        top_domain   path      string  true  "top level domain"
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /resolver/domain_state/{hashed_label}/{top_domain} [get]
func QueryDomainState(xnsContext *types.XnsContext, hashedLabel, topDomain string) (*service.QueryResponse, error) {
	msg, err := NewResolverGetDomainStateMsg(hashedLabel, topDomain)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryAccountState godoc
// @Summary      XNS state for the account
// @Description  Get XNS state for the account
// @Tags         Resolver
// @Accept       json
// @Produce      json
// @Param        account path      string  true  "account address"
// @Success      200     {object}  service.QueryResponse
// @Failure      400     {object}  service.QueryResponse
// @Router       /resolver/account_state/{account} [get]
func QueryAccountState(xnsContext *types.XnsContext, account string) (*service.QueryResponse, error) {
	msg, err := NewResolverGetAccountStateMsg(account)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryFreeAccountState godoc
// @Summary      XNS state for the free domain of the account
// @Description  Get XNS state for the free domain of the account
// @Tags         Resolver
// @Accept       json
// @Produce      json
// @Param        account path      string  true  "account address"
// @Success      200     {object}  service.QueryResponse
// @Failure      400     {object}  service.QueryResponse
// @Router       /resolver/free_account_state/{account} [get]
func QueryFreeAccountState(xnsContext *types.XnsContext, account string) (*service.QueryResponse, error) {
	msg, err := NewResolverGetFreeAccountStateMsg(account)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QuerySubdomain godoc
// @Summary      XNS state for a subdomain
// @Description  Get XNS state for a subdomain
// @Tags         Resolver
// @Accept       json
// @Produce      json
// @Param        hashed_label           path      string  true  "hashed label"
// @Param        hashed_subdomain_label path      string  true  "hashed subdomain label"
// @Success      200              {object}  service.QueryResponse
// @Failure      400              {object}  service.QueryResponse
// @Router       /resolver/subdomain/{hashed_label}/{hashed_subdomain_label} [get]
func QuerySubdomain(xnsContext *types.XnsContext, hashedLabel, hashedSubdomainLabel string) (*service.QueryResponse, error) {
	msg, err := NewResolverGetSubdomainMsg(hashedLabel, hashedSubdomainLabel)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

func Query(xnsContext *types.XnsContext, msg string) (*service.QueryResponse, error) {
	address := xnsContext.GetApp().Contracts.ResolverContract.ContractAddress
	res, err := request.QueryContract(xnsContext, address, msg)
	if err != nil {
		return nil, err
	}

	return service.ParseResponse(res)
}
