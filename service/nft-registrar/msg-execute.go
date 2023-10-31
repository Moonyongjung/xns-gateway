package nftregistrar

import (
	"encoding/json"
	"fmt"

	"github.com/Moonyongjung/xns-gateway/messages"
	"github.com/Moonyongjung/xns-gateway/types"
)

type nftRegistrarSetRegistrarExecuteMsg struct {
	Xns struct {
		SetRegistrar struct {
			RegistrarAddress string `json:"registrar_address"`
		} `json:"set_registrar"`
	} `json:"xns"`
}

func NewNftRegistrarSetRegistrarExecuteMsg(registrarAddress string) (string, error) {
	var msg nftRegistrarSetRegistrarExecuteMsg

	msg.Xns.SetRegistrar.RegistrarAddress = registrarAddress

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarSetResolverExecuteMsg struct {
	Xns struct {
		SetResolver struct {
			ResolverAddress string `json:"resolver_address"`
		} `json:"set_resolver"`
	} `json:"xns"`
}

func NewNftRegistrarSetResolverExecuteMsg(resolverAddress string) (string, error) {
	var msg nftRegistrarSetResolverExecuteMsg

	msg.Xns.SetResolver.ResolverAddress = resolverAddress

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarApproveExecuteMsg struct {
	Approve struct {
		Granter string           `json:"granter"`
		Spender string           `json:"spender"`
		TokenID string           `json:"token_id"`
		Expires messages.Expires `json:"expires,omitempty"`
	} `json:"approve"`
}

func NewNftRegistrarApproveExecuteMsg(granter, spender, tokenID string, atHeight interface{}, atTime interface{}) (string, error) {
	var msg nftRegistrarApproveExecuteMsg

	msg.Approve.Granter = granter
	msg.Approve.Spender = spender
	msg.Approve.TokenID = tokenID

	if atHeight != nil && atTime != nil {
		return "", fmt.Errorf(types.ErrMsgDuplicateExpire)
	}

	switch {
	case atHeight != nil:
		atHeightAssertion, ok := atHeight.(uint64)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.Approve.Expires.AtHeight = atHeightAssertion

	case atTime != nil:
		atTimeAssertion, ok := atTime.(string)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.Approve.Expires.AtTime = atTimeAssertion

	default:
		var never messages.Never
		msg.Approve.Expires.Never = never
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarRevokeExecuteMsg struct {
	Revoke struct {
		Granter string `json:"granter"`
		Spender string `json:"spender"`
		TokenID string `json:"token_id"`
	} `json:"revoke"`
}

func NewNftRegistrarRevokeExecuteMsg(granter, spender, tokenID string) (string, error) {
	var msg nftRegistrarRevokeExecuteMsg

	msg.Revoke.Granter = granter
	msg.Revoke.Spender = spender
	msg.Revoke.TokenID = tokenID

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarApproveAllExecuteMsg struct {
	ApproveAll struct {
		Granter  string           `json:"granter"`
		Operator string           `json:"operator"`
		Expires  messages.Expires `json:"expires,omitempty"`
	} `json:"approve_all"`
}

func NewNftRegistrarApproveAllExecuteMsg(granter, operator string, atHeight interface{}, atTime interface{}) (string, error) {
	var msg nftRegistrarApproveAllExecuteMsg

	msg.ApproveAll.Granter = granter
	msg.ApproveAll.Operator = operator

	if atHeight != nil && atTime != nil {
		return "", fmt.Errorf(types.ErrMsgDuplicateExpire)
	}

	switch {
	case atHeight != nil:
		atHeightAssertion, ok := atHeight.(uint64)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.ApproveAll.Expires.AtHeight = atHeightAssertion

	case atTime != nil:
		atTimeAssertion, ok := atTime.(string)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.ApproveAll.Expires.AtTime = atTimeAssertion

	default:
		var never messages.Never
		msg.ApproveAll.Expires.Never = never
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarRevokeAllExecuteMsg struct {
	RevokeAll struct {
		Granter  string `json:"granter"`
		Operator string `json:"operator"`
	} `json:"revoke_all"`
}

func NewNftRegistrarRevokeAllExecuteMsg(granter, operator string) (string, error) {
	var msg nftRegistrarRevokeAllExecuteMsg

	msg.RevokeAll.Granter = granter
	msg.RevokeAll.Operator = operator

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarTranferNftExecuteMsg struct {
	TransferNft struct {
		Sender    string `json:"sender"`
		Recipient string `json:"recipient"`
		TokenID   string `json:"token_id"`
	} `json:"transfer_nft"`
}

func NewNftRegistrarTransferNftExecuteMsg(sender, recipient, tokenID string) (string, error) {
	var msg nftRegistrarTranferNftExecuteMsg

	msg.TransferNft.Sender = sender
	msg.TransferNft.Recipient = recipient
	msg.TransferNft.TokenID = tokenID

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarSendNftExecuteMsg struct {
	SendNft struct {
		Sender   string `json:"sender"`
		Contract string `json:"contract"`
		TokenID  string `json:"token_id"`
		Msg      string `json:"msg"`
	} `json:"send_nft"`
}

func NewNftRegistrarSendNftExecuteMsg(sender, contract, tokenID, msgBinary string) (string, error) {
	var msg nftRegistrarSendNftExecuteMsg

	msg.SendNft.Sender = sender
	msg.SendNft.Contract = contract
	msg.SendNft.TokenID = tokenID
	msg.SendNft.Msg = msgBinary

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarBurnExecuteMsg struct {
	Burn struct {
		TokenID string `json:"token_id"`
	} `json:"burn"`
}

func NewNftRegistrarBurnExecuteMsg(tokenID string) (string, error) {
	var msg nftRegistrarBurnExecuteMsg

	msg.Burn.TokenID = tokenID

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
