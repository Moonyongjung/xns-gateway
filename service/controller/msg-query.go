package controller

import (
	"encoding/json"
)

type controllerGetConfigQueryMsg struct {
	Config struct{} `json:"config"`
}

func NewControllerGetConfigQueryMsg() (string, error) {
	var msg controllerGetConfigQueryMsg
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type controllerGetCommitmentQueryMsg struct {
	Commitment struct {
		Secret      string `json:"secret"`
		DomainOwner string `json:"domain_owner"`
		Label       string `json:"label"`
	} `json:"commitment"`
}

func NewControllerGetCommitmentQueryMsg(label, domainOwner, secret string) (string, error) {
	var msg controllerGetCommitmentQueryMsg

	msg.Commitment.Label = label
	msg.Commitment.DomainOwner = domainOwner
	msg.Commitment.Secret = secret

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

type controllerControllerBalanceQueryMsg struct {
	Balance struct{} `json:"balance"`
}

func NewControllerControllerBalanceQueryMsg() (string, error) {
	var msg controllerControllerBalanceQueryMsg
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type controllerRegisterAmountQueryMsg struct {
	RegisterCost struct {
		Label          string `json:"label"`
		ExpireDuration uint64 `json:"expire_duration"`
	} `json:"register_cost"`
}

func NewControllerRegisterAmountQueryMsg(label string, expireDuration uint64) (string, error) {
	var msg controllerRegisterAmountQueryMsg

	msg.RegisterCost.Label = label
	msg.RegisterCost.ExpireDuration = expireDuration

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type controllerTopDomainQueryMsg struct {
	TopDomain struct{} `json:"top_domain"`
}

func NewControllerTopDomainQueryMsg() (string, error) {
	var msg controllerTopDomainQueryMsg
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
