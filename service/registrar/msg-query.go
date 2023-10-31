package registrar

import (
	"encoding/json"
	"fmt"

	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"
)

type registrarGetConfigQueryMsg struct {
	Config struct{} `json:"config"`
}

func NewRegistrarGetConfigQueryMsg() (string, error) {
	var msg registrarGetConfigQueryMsg
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type registrarNftRegistrarAddressQueryMsg struct {
	NftRegistrar struct{} `json:"nft_registrar"`
}

func NewRegistrarNftRegistrarAddressQueryMsg() (string, error) {
	var msg registrarNftRegistrarAddressQueryMsg
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type registrarControllersQueryMsg struct {
	Controllers struct {
		StartAfter string `json:"start_after,omitempty"`
		Limit      uint32 `json:"limit,omitempty"`
	} `json:"controllers"`
}

func NewRegistrarControllersQueryMsg(startAfter interface{}, limit interface{}) (string, error) {
	var msg registrarControllersQueryMsg
	if startAfter != nil {
		startAfterAssertion, ok := startAfter.(string)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.Controllers.StartAfter = startAfterAssertion
	}

	if limit != nil {
		limitAssertion, ok := limit.(string)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		limitU32, err := util.FromStringToUint32(limitAssertion)
		if err != nil {
			return "", err
		}
		msg.Controllers.Limit = limitU32
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
