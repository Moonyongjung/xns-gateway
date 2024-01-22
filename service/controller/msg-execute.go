package controller

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/Moonyongjung/xns-gateway/types"
)

type controllerSetConfigExecuteMsg struct {
	SetConfig struct {
		TopPrice            uint64 `json:"top_price,omitempty"`
		MiddlePrice         uint64 `json:"middle_price,omitempty"`
		LowPrice            uint64 `json:"low_price,omitempty"`
		MinRegisterDuration uint64 `json:"min_register_duration,omitempty"`
		RegistrarAddress    string `json:"registrar_address,omitempty"`
	} `json:"set_config"`
}

func NewControllerSetConfigExecuteMsg(
	topPrice, middlePrice, lowPrice, minRegisterDuration, registrarAddress interface{},
) (string, error) {
	var msg controllerSetConfigExecuteMsg

	if topPrice != nil {
		topPriceAssertion, ok := topPrice.(uint64)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.SetConfig.TopPrice = topPriceAssertion
	}

	if middlePrice != nil {
		middlePriceAssertion, ok := middlePrice.(uint64)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.SetConfig.MiddlePrice = middlePriceAssertion
	}

	if lowPrice != nil {
		lowPriceAssertion, ok := lowPrice.(uint64)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.SetConfig.LowPrice = lowPriceAssertion
	}

	if minRegisterDuration != nil {
		minRegisterDurationAssertion, ok := minRegisterDuration.(uint64)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.SetConfig.MinRegisterDuration = minRegisterDurationAssertion
	}

	if registrarAddress != nil {
		registrarAddressAssertion, ok := registrarAddress.(string)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.SetConfig.RegistrarAddress = registrarAddressAssertion
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type controllerDomainRegisterExecuteMsg struct {
	Register struct {
		Label          string `json:"label"`
		ExpireDuration uint64 `json:"expire_duration"`
	} `json:"register"`
}

func NewControllerDomainRegisterExecuteMsg(label string, expireDuration uint64) (string, error) {
	var msg controllerDomainRegisterExecuteMsg

	msg.Register.Label = label
	msg.Register.ExpireDuration = expireDuration

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type controllerDomainRenewExecuteMsg struct {
	Renew struct {
		Label          string `json:"label"`
		ExpireDuration uint64 `json:"expire_duration"`
	} `json:"renew"`
}

func NewControllerDomainRenewExecuteMsg(label string, expireDuration uint64) (string, error) {
	var msg controllerDomainRenewExecuteMsg

	msg.Renew.Label = label
	msg.Renew.ExpireDuration = expireDuration

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type controllerEnrollSubdomainExecuteMsg struct {
	EnrollSubdomain struct {
		Label              string `json:"label"`
		SubdomainLabel     string `json:"subdomain_label"`
		TextData           string `json:"text_data,omitempty"`
		AccountAddressData string `json:"account_address_data,omitempty"`
		PublicKeyData      string `json:"public_key_data,omitempty"`
		ContentHashData    string `json:"content_hash_data,omitempty"`
	} `json:"enroll_subdomain"`
}

func NewControllerEnrollSubdomainExecuteMsg(label, subdomainLabel, textData, accountAddressData, publicKeyData, contentHashData string) (string, error) {
	var msg controllerEnrollSubdomainExecuteMsg

	msg.EnrollSubdomain.Label = label
	msg.EnrollSubdomain.SubdomainLabel = subdomainLabel
	msg.EnrollSubdomain.TextData = textData
	msg.EnrollSubdomain.AccountAddressData = accountAddressData
	msg.EnrollSubdomain.PublicKeyData = publicKeyData
	msg.EnrollSubdomain.ContentHashData = contentHashData

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type controllerRemoveSubdomainExecuteMsg struct {
	RemoveSubdomain struct {
		Label          string `json:"label"`
		SubdomainLabel string `json:"subdomain_label"`
	} `json:"remove_subdomain"`
}

func NewControllerRemoveSubdomainExecuteMsg(label, subdomainLabel string) (string, error) {
	var msg controllerRemoveSubdomainExecuteMsg

	msg.RemoveSubdomain.Label = label
	msg.RemoveSubdomain.SubdomainLabel = subdomainLabel

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type controllerChangeSubdomainLabelExecuteMsg struct {
	ChangeSubdomainLabel struct {
		Label                  string `json:"label"`
		PreviousSubdomainLabel string `json:"previous_subdomain_label"`
		NewSubdomainLabel      string `json:"new_subdomain_label"`
	} `json:"change_subdomain_label"`
}

func NewControllerChangeSubdomainLabelExecuteMsg(label, previousLabel, newLabel string) (string, error) {
	var msg controllerChangeSubdomainLabelExecuteMsg

	msg.ChangeSubdomainLabel.Label = label
	msg.ChangeSubdomainLabel.PreviousSubdomainLabel = previousLabel
	msg.ChangeSubdomainLabel.NewSubdomainLabel = newLabel

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type controllerSaveDomainDataExecuteMsg struct {
	SaveDomainData struct {
		Label              string `json:"label"`
		TextData           string `json:"text_data,omitempty"`
		AccountAddressData string `json:"account_address_data,omitempty"`
		PublicKeyData      string `json:"public_key_data,omitempty"`
		ContentHashData    string `json:"content_hash_data,omitempty"`
	} `json:"save_domain_data"`
}

func NewControllerSaveDomainDataExecuteMsg(label, textData, accountAddressData, publicKeyData, contentHashData string) (string, error) {
	var msg controllerSaveDomainDataExecuteMsg

	msg.SaveDomainData.Label = label
	msg.SaveDomainData.TextData = textData
	msg.SaveDomainData.AccountAddressData = accountAddressData
	msg.SaveDomainData.PublicKeyData = publicKeyData
	msg.SaveDomainData.ContentHashData = contentHashData

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type controllerWithdrawAllExecuteMsg struct {
	WithdrawAll struct{} `json:"withdraw_all"`
}

func NewControllerWithdrawAllExecuteMsg() (string, error) {
	var msg controllerWithdrawAllExecuteMsg

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type controllerWitdrawExecuteMsg struct {
	Withdraw struct {
		RequestWithdrawAmount *big.Int `json:"request_withdraw_amount"`
	} `json:"withdraw"`
}

func NewControllerWithdrawExecuteMsg(requestWithdrawAmount *big.Int) (string, error) {
	var msg controllerWitdrawExecuteMsg

	msg.Withdraw.RequestWithdrawAmount = requestWithdrawAmount

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type controllerPrimaryExecuteMsg struct {
	Primary struct {
		Label string `json:"label"`
	} `json:"primary"`
}

func NewControllerPrimaryExecuteMsg(label string) (string, error) {
	var msg controllerPrimaryExecuteMsg

	msg.Primary.Label = label

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
