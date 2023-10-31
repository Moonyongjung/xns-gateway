package gw

import (
	"path"

	"github.com/Moonyongjung/xns-gateway/types/api"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

type GWServer struct {
	mux *chi.Mux
}

func NewGWServer() *GWServer {
	mux := chi.NewRouter()
	return &GWServer{
		mux: mux,
	}
}

func (g *GWServer) MuxAppend(mux *chi.Mux) *GWServer {
	g.mux = mux
	return g
}

func (g *GWServer) MuxSwagger() *GWServer {
	g.mux.Get(path.Join(api.SwaggerAPI, api.AnyLabel), httpSwagger.Handler())
	return g
}

func (g *GWServer) GetMux() *chi.Mux { return g.mux }
