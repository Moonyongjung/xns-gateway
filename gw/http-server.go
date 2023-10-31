package gw

import (
	"net/http"

	"github.com/Moonyongjung/xns-gateway/service/controller"
	"github.com/Moonyongjung/xns-gateway/service/gateway"
	nftregistrar "github.com/Moonyongjung/xns-gateway/service/nft-registrar"
	"github.com/Moonyongjung/xns-gateway/service/registrar"
	"github.com/Moonyongjung/xns-gateway/service/resolver"
	"github.com/Moonyongjung/xns-gateway/types"
	gwtypes "github.com/Moonyongjung/xns-gateway/types/gw"
	"github.com/Moonyongjung/xns-gateway/util"

	"github.com/rs/cors"

	_ "github.com/Moonyongjung/xns-gateway/docs"
)

// @title XNS gateway
// @version 0.0.1
// @description XNS(XPLA Name Service) gateway provides chain naming service for XPLA

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// Start the HTTP server.
// Users can request to this server in order to query infos of XNS contracts
// by using HTTP REST API.
func StartHttpServer(xnsContext *types.XnsContext) error {
	port := xnsContext.GetApp().Config.XnsGateway.TcpPort
	if err := swaggerInit(xnsContext, port); err != nil {
		return util.LogErr(types.ErrGW, err)
	}

	util.LogInfo("gateway server start...")
	util.LogKV("port", port)

	manager := NewServicesManager(
		nftregistrar.NewService(),
		registrar.NewService(),
		controller.NewService(),
		resolver.NewService(),
		gateway.NewService(),
	)

	server := gwtypes.NewGWServer()

	for _, service := range manager.services {
		server.MuxAppend(service.Router(xnsContext, server.GetMux()))
	}

	server.MuxSwagger()

	handler := cors.Default().Handler(server.GetMux())
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		return util.LogErr(types.ErrGW, err)
	}

	return nil
}
