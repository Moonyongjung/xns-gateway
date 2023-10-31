package service

import (
	"encoding/json"

	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"
)

type QueryResponse struct {
	Data interface{} `json:"data"`
}

func ParseResponse(res string) (*QueryResponse, error) {
	var q QueryResponse
	err := json.Unmarshal([]byte(res), &q)
	if err != nil {
		return nil, util.LogErr(types.ErrParseData, err)
	}

	return &q, nil
}
