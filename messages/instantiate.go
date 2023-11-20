package messages

import "github.com/Moonyongjung/xns-gateway/types"

const (
	DefaultNftRegistrarName   = "xns-nft-regitrar"
	DefaultNftRegistrarSymbol = "XNS"

	DefaultTopDomain = "xpla"

	DefaultTopPrice                      = "100000"
	DefaultMiddlePrice                   = "1000"
	DefaultLowPrice                      = "10"
	DefaultMinRegisterDuration           = "100"
	DefaultMaxFreeDomainRegisterDuration = "30000"
)

func NewNftRegistrarInstantiateMsg(app types.App, xnsBaseCodeID string) string {
	name := app.Config.ContractsOption.NftRegistrar.Name
	if name == "" {
		name = DefaultNftRegistrarName
	}

	symbol := app.Config.ContractsOption.NftRegistrar.Symbol
	if symbol == "" {
		symbol = DefaultNftRegistrarSymbol
	}

	return `{
		"base_code_id":` + xnsBaseCodeID + `,
		"name":"` + name + `",
		"symbol":"` + symbol + `"
	}`
}

func NewRegistrarInstantiateMsg(app types.App, nftRegistrarContractAddress string) string {
	return `{
		"nft_registrar_address":"` + nftRegistrarContractAddress + `"
	}`

}

func NewControllerInstantiateMsg(app types.App, registrarContractAddress string) string {
	topPrice := app.Config.ContractsOption.Controller.TopPrice
	if topPrice == "" {
		topPrice = DefaultTopPrice
	}

	middlePrice := app.Config.ContractsOption.Controller.MiddlePrice
	if middlePrice == "" {
		middlePrice = DefaultMiddlePrice
	}

	lowPrice := app.Config.ContractsOption.Controller.LowPrice
	if lowPrice == "" {
		lowPrice = DefaultLowPrice
	}

	minRegisterDuration := app.Config.ContractsOption.Controller.MinRegisterDuration
	if minRegisterDuration == "" {
		minRegisterDuration = DefaultMinRegisterDuration
	}

	maxFreeDomainRegisterDuration := app.Config.ContractsOption.Controller.MaxFreeDomainRegisterDuration
	if maxFreeDomainRegisterDuration == "" {
		maxFreeDomainRegisterDuration = DefaultMaxFreeDomainRegisterDuration
	}

	topDomain := app.Config.ContractsOption.Controller.TopDomain
	if topDomain == "" {
		topDomain = DefaultTopDomain
	}

	return `{
		"top_price":` + topPrice + `,
		"middle_price":` + middlePrice + `,
		"low_price":` + lowPrice + `,
		"min_register_duration":` + minRegisterDuration + `,
		"max_free_domain_register_duration":` + maxFreeDomainRegisterDuration + `,
		"registrar_address":"` + registrarContractAddress + `",
		"top_domain":"` + topDomain + `"
	}`
}

func NewResolverInstantiateMsg(app types.App, registrarAddress, nftRegistrarAddress string) string {
	return `{
		"registrar_address":"` + registrarAddress + `",
		"nft_registrar_address":"` + nftRegistrarAddress + `"
	}`
}
