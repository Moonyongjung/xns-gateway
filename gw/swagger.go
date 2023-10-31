package gw

import (
	"github.com/Moonyongjung/xns-gateway/docs"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/types/api"

	"github.com/cosmos/cosmos-sdk/server"
)

func swaggerInit(xnsContext *types.XnsContext, port string) error {
	var err error

	swaggerHost := xnsContext.GetApp().Config.SwaggerOption.ServerHostIp
	if swaggerHost == "" {
		swaggerHost, err = server.ExternalIP()
		if err != nil {
			return err
		}
	}

	docs.SwaggerInfo.Host = swaggerHost + ":" + port
	docs.SwaggerInfo.BasePath = api.APIVersion[:len(api.APIVersion)]

	return nil
}
