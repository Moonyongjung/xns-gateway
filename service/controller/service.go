package controller

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
	return ControllerAPI
}

func (c *contractService) Router(xnsContext *types.XnsContext, mux *chi.Mux) *chi.Mux {
	mux.HandleFunc(path.Join(api.APIVersion, ControllerAPI, QueryCommitmentLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		label, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		domainOwner, err := util.ParseAPI2Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		secret, err := util.ParseAPI3Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryCommitment(xnsContext, label, domainOwner, secret)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, ControllerAPI, QueryBalanceLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryBalance(xnsContext)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, ControllerAPI, QueryConfigLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryConfig(xnsContext)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, ControllerAPI, QueryRegisterCostLabel, api.AnyLabel), func(w http.ResponseWriter, r *http.Request) {
		label, err := util.ParseAPI1Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		duration, err := util.ParseAPI2Data(r)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}
		res, err := QueryRegisterCost(xnsContext, label, duration)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, ControllerAPI, QueryTopDomainLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryTopDomain(xnsContext)
		service.Consume(w, res, err, xnsContext)
	})

	return mux
}
