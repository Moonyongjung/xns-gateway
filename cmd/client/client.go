package client

import (
	"github.com/Moonyongjung/xns-gateway/key"
	"github.com/Moonyongjung/xns-gateway/types"

	"github.com/xpladev/xpla.go/client"
	"github.com/xpladev/xpla.go/provider"
)

func NewXplaClient(xnsContext *types.XnsContext, isExecute bool) (provider.XplaClient, error) {
	app := xnsContext.GetApp()
	xplaClient := client.NewXplaClient(app.Config.TargetChain.ChainID).
		WithURL(app.Config.TargetChain.LcdURL).
		WithGasAdjustment(app.Config.TargetChain.GasAdj).
		WithGasLimit(app.Config.TargetChain.GasLimit).
		WithBroadcastMode(app.Config.TargetChain.BroadcastMode)

	if isExecute {
		privKey, err := key.ExtractKey(xnsContext, app.Home)
		if err != nil {
			return nil, err
		}

		xplaClient.WithPrivateKey(privKey)
	}

	return xplaClient, nil
}
