package resolver

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
	return ResolverAPI
}

func (c *contractService) Router(xnsContext *types.XnsContext, mux *chi.Mux) *chi.Mux {
	mux.HandleFunc(path.Join(api.APIVersion, ResolverAPI, QueryConfigLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryConfig(xnsContext)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, ResolverAPI, QueryDomainStateLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
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
		res, err := QueryDomainState(xnsContext, label, topDomain)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, ResolverAPI, QueryAccountStateLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		account, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryAccountState(xnsContext, account)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, ResolverAPI, QueryFreeAccountStateLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		account, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryFreeAccountState(xnsContext, account)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, ResolverAPI, QuerySubdomainLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		domain, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		subdomain, err := util.ParseAPI2Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QuerySubdomain(xnsContext, domain, subdomain)
		service.Consume(w, res, err, xnsContext)
	})

	return mux
}
