package service

import (
	"net/http"

	"github.com/Moonyongjung/xns-gateway/types"
	gwtypes "github.com/Moonyongjung/xns-gateway/types/gw"
)

func Consume(w http.ResponseWriter, res interface{}, err error, xnsContext *types.XnsContext) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(gwtypes.ErrorReturn(types.ErrHttpServer, err.Error()))

	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(gwtypes.DataReturn(res))
	}
}
