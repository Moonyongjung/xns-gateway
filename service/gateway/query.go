package gateway

import (
	"github.com/Moonyongjung/xns-gateway/types"
)

// QueryXnsContractsInfo godoc
// @Summary      XNS contracts info
// @Description  Get XNS contracts info of the XNS gateway
// @Tags         XNS-gateway
// @Accept       json
// @Produce      json
// @Success      200          {object}  service.QueryResponse
// @Failure      400          {object}  service.QueryResponse
// @Router       /xns-contracts-info [get]
func QueryXnsContractsInfo(xnsContext *types.XnsContext) (types.Contracts, error) {
	return xnsContext.GetApp().Contracts, nil
}
