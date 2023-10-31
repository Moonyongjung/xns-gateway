package gateway

import (
	"net/http"
	"path"

	"github.com/Moonyongjung/xns-gateway/service"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/types/api"

	"github.com/go-chi/chi"
)

type contractService struct{}

func NewService() service.Service {
	return &contractService{}
}

func (c *contractService) Name() string {
	return GatewayAPI
}

func (c *contractService) Router(xnsContext *types.XnsContext, mux *chi.Mux) *chi.Mux {
	mux.HandleFunc(path.Join(api.APIVersion, QueryXnsContractsInfoLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryXnsContractsInfo(xnsContext)
		if err != nil {
			service.Consume(w, nil, err, xnsContext)
			return
		}

		service.Consume(w, res, err, xnsContext)
	})

	return mux
}
