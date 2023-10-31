package registrar

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
	return RegistrarAPI
}

func (c *contractService) Router(xnsContext *types.XnsContext, mux *chi.Mux) *chi.Mux {
	mux.HandleFunc(path.Join(api.APIVersion, RegistrarAPI, QueryConfigLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryConfig(xnsContext)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, RegistrarAPI, QueryNftRegistrarLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryNftRegistrar(xnsContext)
		service.Consume(w, res, err, xnsContext)
	})

	mux.HandleFunc(path.Join(api.APIVersion, RegistrarAPI, QueryControllersLabel), func(w http.ResponseWriter, r *http.Request) {
		res, err := QueryControllers(xnsContext, service.DefaultGWStartAfter, service.DefaultGWLimit)
		service.Consume(w, res, err, xnsContext)
	})

	return mux
}
