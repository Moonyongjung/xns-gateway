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

type controllerRegisterCostQueryMsg struct {
	RegisterCost struct {
		Label          string `json:"label"`
		ExpireDuration uint64 `json:"expire_duration"`
	} `json:"register_cost"`
}

func NewControllerRegisterCostQueryMsg(label string, expireDuration uint64) (string, error) {
	var msg controllerRegisterCostQueryMsg

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
