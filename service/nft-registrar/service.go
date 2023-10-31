package nftregistrar

import (
	"net/http"
	"path"

	"github.com/Moonyongjung/xns-gateway/service"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/types/api"
	"github.com/Moonyongjung/xns-gateway/util"

	"github.com/go-chi/chi"
)

type contractService struct{}

func NewService() service.Service {
	return &contractService{}
}

func (c *contractService) Name() string {
	return NftRegistrarAPI
}

func (c *contractService) Router(xnsContext *types.XnsContext, mux *chi.Mux) *chi.Mux {
	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryRegistrarLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryRegistrar(xnsContext)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryResolverLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryResolver(xnsContext)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryAccountByDomainLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		domain, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		topDomain, err := util.ParseAPI2Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryAccountByDomain(xnsContext, domain, topDomain)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryDomainByAccountLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		account, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryDomainByAccount(xnsContext, account)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryDomainHistoryByAccountLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		account, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryDomainHistoryByAccount(xnsContext, account)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryExpiredLatestDomainLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		domain, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		topDomain, err := util.ParseAPI2Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryExpiredLatestDomain(xnsContext, domain, topDomain)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryTokenIDByDomainLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		domain, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		topDomain, err := util.ParseAPI2Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryTokenIDByDomain(xnsContext, domain, topDomain)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QuerySubdomainLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		subdomainLabel, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		label, err := util.ParseAPI2Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		topDomain, err := util.ParseAPI3Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QuerySubdomain(xnsContext, subdomainLabel, label, topDomain)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QuerySubdomainsLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		label, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		topDomain, err := util.ParseAPI2Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QuerySubdomains(xnsContext, label, topDomain)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryHashResultLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		hash, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryHashResult(xnsContext, hash)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryContractInfoLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryContractInfo(xnsContext)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryNftInfoLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		tokenID, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryNftInfo(xnsContext, tokenID)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryOwnerOfLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		tokenID, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryOwnerOf(xnsContext, tokenID, service.DefaultGWIncludeExpired)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryAllNftInfoLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		tokenID, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryAllNftInfo(xnsContext, tokenID, service.DefaultGWIncludeExpired)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryAllOperatorsLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		owner, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryAllOperators(xnsContext, owner, service.DefaultGWIncludeExpired, service.DefaultGWStartAfter, service.DefaultGWLimit)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryNumTokensLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryNumTokens(xnsContext)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryTokensLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		owner, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryTokens(xnsContext, owner, service.DefaultGWStartAfter, service.DefaultGWLimit)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryAllTokensLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryAllTokens(xnsContext, service.DefaultGWStartAfter, service.DefaultGWLimit)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryApprovalLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		tokenID, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}

		spender, err := util.ParseAPI2Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryApproval(xnsContext, tokenID, spender, service.DefaultGWIncludeExpired)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryApprovalsLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		tokenID, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryApprovals(xnsContext, tokenID, service.DefaultGWIncludeExpired)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, NftRegistrarAPI, QueryMinterLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryMinter(xnsContext)
		service.Consume(w, res, err, xnsContext)
	})

	return mux
}
