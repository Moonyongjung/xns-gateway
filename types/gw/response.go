package gw

import (
	"encoding/json"
	"strings"

	"github.com/Moonyongjung/xns-gateway/types"
)

type HttpResponse struct {
	ResCode uint64      `json:"res_code"`
	ResMsg  string      `json:"res_msg"`
	ResData interface{} `json:"res_data"`
}

func HttpResponseByte(res HttpResponse) []byte {
	responseByte, err := json.Marshal(res)
	if err != nil {
		return []byte(err.Error())
	}

	return responseByte
}

// Return success data.
func DataReturn(data interface{}) []byte {
	httpResponse := HttpResponse{
		ResCode: types.Success.Code(),
		ResMsg:  types.Success.Desc(),
		ResData: data,
	}

	return HttpResponseByte(httpResponse)
}

// Return error data.
func ErrorReturn(errType types.XnsError, errMsg string) []byte {
	httpResponse := HttpResponse{
		ResCode: errType.Code(),
		ResMsg:  errType.Desc(),
		ResData: parsingErrorMessage(errMsg),
	}

	return HttpResponseByte(httpResponse)
}

func parsingErrorMessage(errMsg string) string {
	errMsgConv := strings.ReplaceAll(errMsg, "\n", "")
	errMsgConv = strings.ReplaceAll(errMsgConv, "\"", "")
	errMsgConv = strings.ReplaceAll(errMsgConv, "  ", "")
	errMsgConv = strings.ReplaceAll(errMsgConv, ",", ", ")

	return errMsgConv
}
