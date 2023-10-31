package resolver

import (
	"encoding/json"
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

type resolverGetDomainStateMsg struct {
	DomainState struct {
		HashedLabel string `json:"hashed_label"`
		TopDomain   string `json:"top_domain"`
	} `json:"domain_state"`
}

func NewResolverGetDomainStateMsg(hashedLabel, topDomain string) (string, error) {
	var msg resolverGetDomainStateMsg

	msg.DomainState.HashedLabel = hashedLabel
	msg.DomainState.TopDomain = topDomain

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type resolverGetAccountStateMsg struct {
	AccountState struct {
		Account string `json:"account"`
	} `json:"account_state"`
}

func NewResolverGetAccountStateMsg(account string) (string, error) {
	var msg resolverGetAccountStateMsg

	msg.AccountState.Account = account

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type resolverGetFreeAccountStateMsg struct {
	FreeAccountState struct {
		Account string `json:"account"`
	} `json:"free_account_state"`
}

func NewResolverGetFreeAccountStateMsg(account string) (string, error) {
	var msg resolverGetFreeAccountStateMsg

	msg.FreeAccountState.Account = account

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type resolverGetSubdomainMsg struct {
	Subdomain struct {
		HashedLabel          string `json:"hashed_label"`
		HashedSubdomainLabel string `json:"hashed_subdomain_label"`
	} `json:"subdomain"`
}

func NewResolverGetSubdomainMsg(hashedLabel, hashedSubdomainLabel string) (string, error) {
	var msg resolverGetSubdomainMsg

	msg.Subdomain.HashedLabel = hashedLabel
	msg.Subdomain.HashedSubdomainLabel = hashedSubdomainLabel

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
