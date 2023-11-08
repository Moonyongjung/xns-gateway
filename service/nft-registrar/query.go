package nftregistrar

import (
	"github.com/Moonyongjung/xns-gateway/cmd/flag"
	"github.com/Moonyongjung/xns-gateway/request"
	"github.com/Moonyongjung/xns-gateway/service"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"
)

// QueryRegistrar godoc
// @Summary      Get registrar
// @Description  Get connected address of the registrar contract
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /nft-registrar/registrar [get]
func QueryRegistrar(xnsContext *types.XnsContext) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarGetRegistrarQueryMsg()
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryResolver godoc
// @Summary      Get resolver
// @Description  Get connected address of the resolver contract
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /nft-registrar/resolver [get]
func QueryResolver(xnsContext *types.XnsContext) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarGetResolverQueryMsg()
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryAccountByDomain godoc
// @Summary      Retrieve the account address by the domain
// @Description  Get account address that has the retrieved domain
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        label      path      string  true  "domain name to retrieve"
// @Param        top_domain path      string  true  "top level domain"
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /nft-registrar/account-by-domain/{label}/{top-domain} [get]
func QueryAccountByDomain(xnsContext *types.XnsContext, label, topDomain string) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarAccountByDomainQueryMsg(label, topDomain)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryDomainbyAccount godoc
// @Summary      Retrieve the domain by the account address
// @Description  Get domain by retrieving account address
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        account path      string  true  "account address to retrieve"
// @Success      200     {object}  service.QueryResponse
// @Failure      400     {object}  service.QueryResponse
// @Router       /nft-registrar/domain-by-account/{account} [get]
func QueryDomainByAccount(xnsContext *types.XnsContext, account string) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarDomainByAccountQueryMsg(account)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryHistoryDomainbyAccount godoc
// @Summary      Retrieve the domain historyby the account address
// @Description  Get domain history by retrieving account address
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        account path      string  true  "account address to retrieve"
// @Success      200     {object}  service.QueryResponse
// @Failure      400     {object}  service.QueryResponse
// @Router       /nft-registrar/domain-history-by-account/{account} [get]
func QueryDomainHistoryByAccount(xnsContext *types.XnsContext, account string) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarDomainHistoryByAccountQueryMsg(account)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryExpiredLatestDomain godoc
// @Summary      Get domain info with expiration
// @Description  Get domain info with expiration and the account who has the domain lastly
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        label      path      string  true  "domain name to retrieve"
// @Param        top_domain path      string  true  "top level domain"
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /nft-registrar/expired-latest-domain/{label}/{top-domain} [get]
func QueryExpiredLatestDomain(xnsContext *types.XnsContext, label, topDomain string) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarExpiredLatestDomainQueryMsg(label, topDomain)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryTokenIDByDomain godoc
// @Summary      Get token ID by domain
// @Description  Get token ID by domain
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        label      path      string  true  "domain name to retrieve"
// @Param        top_domain path      string  true  "top level domain"
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /nft-registrar/tokenid-by-domain/{label}/{top-domain} [get]
func QueryTokenIDByDomain(xnsContext *types.XnsContext, label, topDomain string) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarTokenIDByDomainQueryMsg(label, topDomain)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QuerySubdomain godoc
// @Summary      Get subdomain data
// @Description  Get subdomain mapped data
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        subdomain_label path      string  true  "subdomain label"
// @Param        label           path      string  true  "domain name to retrieve"
// @Param        top_domain      path      string  true  "top level domain"
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /nft-registrar/subdomain/{subdomain-label}/{label}/{top-domain} [get]
func QuerySubdomain(xnsContext *types.XnsContext, subdomainLabel, label, topDomain string) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarSubdomainQueryMsg(subdomainLabel, label, topDomain)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QuerySubdomains godoc
// @Summary      Get subdomains
// @Description  Get subdomains list of the domain
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        label           path      string  true  "domain name to retrieve"
// @Param        top_domain      path      string  true  "top level domain"
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /nft-registrar/subdomains/{label}/{top-domain} [get]
func QuerySubdomains(xnsContext *types.XnsContext, label, topDomain string) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarSubdomainsQueryMsg(label, topDomain)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryHashResult godoc
// @Summary      Get hash result
// @Description  Get hash result for an input
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        input  path      string  true  "any input"
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /nft-registrar/hash/{input} [get]
func QueryHashResult(xnsContext *types.XnsContext, input string) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarHashResultQueryMsg(input)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryContractInfo godoc
// @Summary      Get the contract info
// @Description  Get the NFT registrar contract info
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Success      200    {object}  service.QueryResponse
// @Failure      400    {object}  service.QueryResponse
// @Router       /nft-registrar/contract-info [get]
func QueryContractInfo(xnsContext *types.XnsContext) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarContractInfoQueryMsg()
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryNftInfo godoc
// @Summary      NFT info of the token ID
// @Description  Get the NFT info by retrieving the token ID
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        token-id path      string  true  "the NFT token ID"
// @Success      200      {object}  service.QueryResponse
// @Failure      400      {object}  service.QueryResponse
// @Router       /nft-registrar/nft-info/{token-id} [get]
func QueryNftInfo(xnsContext *types.XnsContext, tokenID string) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarNftInfoQueryMsg(tokenID)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryOwnerOf godoc
// @Summary      the owner info of the NFT
// @Description  Get the owner info of the NFT by retrieving the token ID
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        token-id path      string  true  "the NFT token ID"
// @Success      200      {object}  service.QueryResponse
// @Failure      400      {object}  service.QueryResponse
// @Router       /nft-registrar/owner-of/{token-id} [get]
func QueryOwnerOf(xnsContext *types.XnsContext, tokenID string, includeExpired bool) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarOwnerOfQueryMsg(tokenID, includeExpired)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryAllNftInfo godoc
// @Summary      All NFT info of the token ID
// @Description  Get the all NFT info by retrieving the token ID
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        token-id path      string  true  "the NFT token ID"
// @Success      200      {object}  service.QueryResponse
// @Failure      400      {object}  service.QueryResponse
// @Router       /nft-registrar/all-nft-info/{token-id} [get]
func QueryAllNftInfo(xnsContext *types.XnsContext, tokenID string, includeExpired bool) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarAllNftInfoQueryMsg(tokenID, includeExpired)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryAllOperators godoc
// @Summary      Operators info of the owner
// @Description  Get operators info of the owner
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        owner path      string  true  "the owner owns NFT"
// @Success      200   {object}  service.QueryResponse
// @Failure      400   {object}  service.QueryResponse
// @Router       /nft-registrar/all-operators/{owner} [get]
func QueryAllOperators(xnsContext *types.XnsContext, owner string, includeExpired bool, startAfter, limit string) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarAllOperatorsQueryMsg(owner, includeExpired, flag.StartAfterI(startAfter), flag.LimitI(limit))
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryNumTokens godoc
// @Summary      The number of the all tokens
// @Description  Get the number of the all tokens in the NFT registrar contract
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Success      200   {object}  service.QueryResponse
// @Failure      400   {object}  service.QueryResponse
// @Router       /nft-registrar/num-tokens [get]
func QueryNumTokens(xnsContext *types.XnsContext) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarNumTokensQueryMsg()
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryTokens godoc
// @Summary      Tokens that the owner has
// @Description  Get tokens that the owner has
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        owner path      string  true  "the owner owns NFT"
// @Success      200   {object}  service.QueryResponse
// @Failure      400   {object}  service.QueryResponse
// @Router       /nft-registrar/tokens/{owner} [get]
func QueryTokens(xnsContext *types.XnsContext, owner, startAfter, limit string) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarTokensQueryMsg(owner, flag.StartAfterI(startAfter), flag.LimitI(limit))
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryAllTokens godoc
// @Summary      All tokens in the contract
// @Description  Get all tokens in the NFT registrar contract
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Success      200   {object}  service.QueryResponse
// @Failure      400   {object}  service.QueryResponse
// @Router       /nft-registrar/all-tokens [get]
func QueryAllTokens(xnsContext *types.XnsContext, startAfter, limit string) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarAlltokensQueryMsg(flag.StartAfterI(startAfter), flag.LimitI(limit))
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryApproval godoc
// @Summary      Approval of the token ID
// @Description  Get approval of the token ID with the spender
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        token-id path      string  true  "the NFT token ID"
// @Param        spender  path      string  true  "the spender of the token ID"
// @Success      200   {object}  service.QueryResponse
// @Failure      400   {object}  service.QueryResponse
// @Router       /nft-registrar/approval/{token-id}/{spender} [get]
func QueryApproval(xnsContext *types.XnsContext, tokenID, spender string, includeExpired bool) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarApprovalQueryMsg(tokenID, spender, includeExpired)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryApprovals godoc
// @Summary      Approvals of the token ID
// @Description  Get approvals of the token ID
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Param        token-id path      string  true  "the owner owns NFT"
// @Success      200   {object}  service.QueryResponse
// @Failure      400   {object}  service.QueryResponse
// @Router       /nft-registrar/approvals/{token-id} [get]
func QueryApprovals(xnsContext *types.XnsContext, tokenID string, includeExpired bool) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarApprovalsQueryMsg(tokenID, includeExpired)
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

// QueryMinter godoc
// @Summary      Minter of the contract
// @Description  The Minter of the NFT registrar contract
// @Tags         NFT-Registrar
// @Accept       json
// @Produce      json
// @Success      200   {object}  service.QueryResponse
// @Failure      400   {object}  service.QueryResponse
// @Router       /nft-registrar/minter [get]
func QueryMinter(xnsContext *types.XnsContext) (*service.QueryResponse, error) {
	msg, err := NewNftRegistrarMinterQueryMsg()
	if err != nil {
		return nil, util.LogErr(types.ErrNewMsg, err)
	}

	return Query(xnsContext, msg)
}

func Query(xnsContext *types.XnsContext, msg string) (*service.QueryResponse, error) {
	address := xnsContext.GetApp().Contracts.NftRegistrarContract.ContractAddress
	res, err := request.QueryContract(xnsContext, address, msg)
	if err != nil {
		return nil, err
	}

	return service.ParseResponse(res)
}
