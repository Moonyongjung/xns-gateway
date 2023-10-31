package nftregistrar

import (
	"encoding/json"
	"fmt"

	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"
)

type nftRegistrarGetRegistrarQueryMsg struct {
	Xns struct {
		Registrar struct{} `json:"registrar"`
	} `json:"xns"`
}

func NewNftRegistrarGetRegistrarQueryMsg() (string, error) {
	var msg nftRegistrarGetRegistrarQueryMsg
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarGetResolverQueryMsg struct {
	Xns struct {
		Resolver struct{} `json:"resolver"`
	} `json:"xns"`
}

func NewNftRegistrarGetResolverQueryMsg() (string, error) {
	var msg nftRegistrarGetResolverQueryMsg
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarAccountByDomainQueryMsg struct {
	Xns struct {
		AccountByDomain struct {
			Label     string `json:"label"`
			TopDomain string `json:"top_domain"`
		} `json:"account_by_domain"`
	} `json:"xns"`
}

func NewNftRegistrarAccountByDomainQueryMsg(label, topDomain string) (string, error) {
	var msg nftRegistrarAccountByDomainQueryMsg

	msg.Xns.AccountByDomain.Label = label
	msg.Xns.AccountByDomain.TopDomain = topDomain

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarDomainByAccountQueryMsg struct {
	Xns struct {
		DomainByAccount struct {
			AccountAddress string `json:"account_address"`
		} `json:"domain_by_account"`
	} `json:"xns"`
}

func NewNftRegistrarDomainByAccountQueryMsg(account string) (string, error) {
	var msg nftRegistrarDomainByAccountQueryMsg

	msg.Xns.DomainByAccount.AccountAddress = account

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarDomainHistoryByAccountQueryMsg struct {
	Xns struct {
		DomainHistoryByAccount struct {
			AccountAddress string `json:"account_address"`
		} `json:"domain_history_by_account"`
	} `json:"xns"`
}

func NewNftRegistrarDomainHistoryByAccountQueryMsg(account string) (string, error) {
	var msg nftRegistrarDomainHistoryByAccountQueryMsg

	msg.Xns.DomainHistoryByAccount.AccountAddress = account

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarExpiredLatestDomainQueryMsg struct {
	Xns struct {
		ExpiredLatestDomain struct {
			Label     string `json:"label"`
			TopDomain string `json:"top_domain"`
		} `json:"expired_latest_domain"`
	} `json:"xns"`
}

func NewNftRegistrarExpiredLatestDomainQueryMsg(label, topDomain string) (string, error) {
	var msg nftRegistrarExpiredLatestDomainQueryMsg

	msg.Xns.ExpiredLatestDomain.Label = label
	msg.Xns.ExpiredLatestDomain.TopDomain = topDomain

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarTokenIDByDomainQueryMsg struct {
	Xns struct {
		TokenIdByDomain struct {
			Label     string `json:"label"`
			TopDomain string `json:"top_domain"`
		} `json:"token_id_by_domain"`
	} `json:"xns"`
}

func NewNftRegistrarTokenIDByDomainQueryMsg(label, topDomain string) (string, error) {
	var msg nftRegistrarTokenIDByDomainQueryMsg

	msg.Xns.TokenIdByDomain.Label = label
	msg.Xns.TokenIdByDomain.TopDomain = topDomain

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarSubdomainQueryMsg struct {
	Xns struct {
		Subdomain struct {
			SubdomainLabel string `json:"subdomain_label"`
			Label          string `json:"label"`
			TopDomain      string `json:"top_domain"`
		} `json:"subdomain"`
	} `json:"xns"`
}

func NewNftRegistrarSubdomainQueryMsg(subdomainLabel, label, topDomain string) (string, error) {
	var msg nftRegistrarSubdomainQueryMsg

	msg.Xns.Subdomain.SubdomainLabel = subdomainLabel
	msg.Xns.Subdomain.Label = label
	msg.Xns.Subdomain.TopDomain = topDomain

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarSubdomainsQueryMsg struct {
	Xns struct {
		Subdomains struct {
			Label     string `json:"label"`
			TopDomain string `json:"top_domain"`
		} `json:"subdomains"`
	} `json:"xns"`
}

func NewNftRegistrarSubdomainsQueryMsg(label, topDomain string) (string, error) {
	var msg nftRegistrarSubdomainsQueryMsg

	msg.Xns.Subdomains.Label = label
	msg.Xns.Subdomains.TopDomain = topDomain

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarHashResultQueryMsg struct {
	Xns struct {
		Hash struct {
			Input string `json:"input"`
		} `json:"hash"`
	} `json:"xns"`
}

func NewNftRegistrarHashResultQueryMsg(input string) (string, error) {
	var msg nftRegistrarHashResultQueryMsg

	msg.Xns.Hash.Input = input

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarContractInfoQueryMsg struct {
	ContractInfo struct{} `json:"contract_info"`
}

func NewNftRegistrarContractInfoQueryMsg() (string, error) {
	var msg nftRegistrarContractInfoQueryMsg
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarNftInfoQueryMsg struct {
	NftInfo struct {
		TokenID string `json:"token_id"`
	} `json:"nft_info"`
}

func NewNftRegistrarNftInfoQueryMsg(tokenID string) (string, error) {
	var msg nftRegistrarNftInfoQueryMsg

	msg.NftInfo.TokenID = tokenID

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarOwnerOfQueryMsg struct {
	OwnerOf struct {
		TokenID        string `json:"token_id"`
		IncludeExpired bool   `json:"include_expired,omitempty"`
	} `json:"owner_of"`
}

func NewNftRegistrarOwnerOfQueryMsg(tokenID string, includeExpired interface{}) (string, error) {
	var msg nftRegistrarOwnerOfQueryMsg

	msg.OwnerOf.TokenID = tokenID

	if includeExpired != nil {
		includeExpiredAssertion, ok := includeExpired.(bool)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.OwnerOf.IncludeExpired = includeExpiredAssertion
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarAllNftInfoQueryMsg struct {
	AllNftInfo struct {
		TokenID        string `json:"token_id"`
		IncludeExpired bool   `json:"include_expired,omitempty"`
	} `json:"all_nft_info"`
}

func NewNftRegistrarAllNftInfoQueryMsg(tokenID string, includeExpired interface{}) (string, error) {
	var msg nftRegistrarAllNftInfoQueryMsg

	msg.AllNftInfo.TokenID = tokenID

	if includeExpired != nil {
		includeExpiredAssertion, ok := includeExpired.(bool)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.AllNftInfo.IncludeExpired = includeExpiredAssertion
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarAllOpteratorsQueryMsg struct {
	AllOperators struct {
		Owner          string `json:"owner"`
		IncludeExpired bool   `json:"include_expired,omitempty"`
		StartAfter     string `json:"start_after,omitempty"`
		Limit          uint32 `json:"limit,omitempty"`
	} `json:"all_operators"`
}

func NewNftRegistrarAllOperatorsQueryMsg(owner string, includeExpired, startAfter, limit interface{}) (string, error) {
	var msg nftRegistrarAllOpteratorsQueryMsg

	msg.AllOperators.Owner = owner

	if includeExpired != nil {
		includeExpiredAssertion, ok := includeExpired.(bool)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.AllOperators.IncludeExpired = includeExpiredAssertion
	}

	if startAfter != nil {
		startAfterAssertion, ok := startAfter.(string)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.AllOperators.StartAfter = startAfterAssertion
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
		msg.AllOperators.Limit = limitU32
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarNumTokensQueryMsg struct {
	NumTokens struct{} `json:"num_tokens"`
}

func NewNftRegistrarNumTokensQueryMsg() (string, error) {
	var msg nftRegistrarNumTokensQueryMsg
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarTokensQueryMsg struct {
	Tokens struct {
		Owner      string `json:"owner"`
		StartAfter string `json:"start_after,omitempty"`
		Limit      uint32 `json:"limit,omitempty"`
	} `json:"tokens"`
}

func NewNftRegistrarTokensQueryMsg(owner string, startAfter interface{}, limit interface{}) (string, error) {
	var msg nftRegistrarTokensQueryMsg

	msg.Tokens.Owner = owner

	if startAfter != nil {
		startAfterAssertion, ok := startAfter.(string)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.Tokens.StartAfter = startAfterAssertion
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
		msg.Tokens.Limit = limitU32
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarAllTokensQueryMsg struct {
	AllTokens struct {
		StartAfter string `json:"start_after,omitempty"`
		Limit      uint32 `json:"limit,omitempty"`
	} `json:"all_tokens"`
}

func NewNftRegistrarAlltokensQueryMsg(startAfter interface{}, limit interface{}) (string, error) {
	var msg nftRegistrarAllTokensQueryMsg
	if startAfter != nil {
		startAfterAssertion, ok := startAfter.(string)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.AllTokens.StartAfter = startAfterAssertion
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
		msg.AllTokens.Limit = limitU32
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarApprovalQueryMsg struct {
	Approval struct {
		TokenID        string `json:"token_id"`
		Spender        string `json:"spender"`
		IncludeExpired bool   `json:"include_expired,omitempty"`
	} `json:"approval"`
}

func NewNftRegistrarApprovalQueryMsg(tokenID, spender string, includeExpired interface{}) (string, error) {
	var msg nftRegistrarApprovalQueryMsg

	msg.Approval.TokenID = tokenID
	msg.Approval.Spender = spender

	if includeExpired != nil {
		includeExpiredAssertion, ok := includeExpired.(bool)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.Approval.IncludeExpired = includeExpiredAssertion
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarApprovalsQueryMsg struct {
	Approvals struct {
		TokenID        string `json:"token_id"`
		IncludeExpired bool   `json:"include_expired,omitempty"`
	} `json:"approvals"`
}

func NewNftRegistrarApprovalsQueryMsg(tokenID string, includeExpired interface{}) (string, error) {
	var msg nftRegistrarApprovalsQueryMsg

	msg.Approvals.TokenID = tokenID

	if includeExpired != nil {
		includeExpiredAssertion, ok := includeExpired.(bool)
		if !ok {
			return "", fmt.Errorf(types.ErrMsgInvalidDataType)
		}
		msg.Approvals.IncludeExpired = includeExpiredAssertion
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type nftRegistrarMinterQueryMsg struct {
	Minter struct{} `json:"minter"`
}

func NewNftRegistrarMinterQueryMsg() (string, error) {
	var msg nftRegistrarMinterQueryMsg
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
