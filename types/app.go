package types

type App struct {
	Home      string
	Config    Config
	Contracts Contracts
}

type Config struct {
	XnsGateway      XnsGateway      `yaml:"xns_gateway" json:"xns_gateway"`
	TargetChain     TargetChain     `yaml:"target_chain" json:"target_chain"`
	ContractsOption ContractsOption `yaml:"contracts_option" json:"contracts_option"`
	SwaggerOption   SwaggerOption   `yaml:"swagger_option" json:"swagger_option"`
}

type XnsGateway struct {
	TcpPort string `yaml:"tcp_port" json:"tcp_port"`
	KeyTest bool   `yaml:"key_test" json:"key_test"`
}

type TargetChain struct {
	ChainID       string `yaml:"chain_id" json:"chain_id"`
	LcdURL        string `yaml:"lcd_url" json:"lcd_url"`
	GasAdj        string `yaml:"gas_adj" json:"gas_adj"`
	GasLimit      string `yaml:"gas_limit" json:"gas_limit"`
	BroadcastMode string `yaml:"broadcast_mode" json:"broadcast_mode"`
}

type Contracts struct {
	BaseContract struct {
		CodeID string `yaml:"code_id" json:"code_id"`
	} `yaml:"base_contract,omitempty" json:"base_contract,omitempty"`
	NftRegistrarContract ContractInfo `yaml:"nft_registrar_contract,omitempty" json:"nft_registrar_contract,omitempty"`
	RegistrarContract    ContractInfo `yaml:"registrar_contract,omitempty" json:"registrar_contract,omitempty"`
	ControllerContract   ContractInfo `yaml:"controller_contract,omitempty" json:"controller_contract,omitempty"`
	ResolverContract     ContractInfo `yaml:"resolver_contract,omitempty" json:"resolver_contract,omitempty"`
}

type ContractInfo struct {
	CodeID          string `yaml:"code_id" json:"code_id"`
	ContractAddress string `yaml:"contract_address" json:"contract_address"`
}

type ContractsOption struct {
	NftRegistrar struct {
		Name   string `yaml:"name" json:"name"`
		Symbol string `yaml:"symbol" json:"symbol"`
	} `yaml:"nft_registrar" json:"nft_registrar"`

	Controller struct {
		TopPrice                      string `yaml:"top_price" json:"top_price"`
		MiddlePrice                   string `yaml:"middle_price" json:"middle_price"`
		LowPrice                      string `yaml:"low_price" json:"low_price"`
		MinRegisterDuration           string `yaml:"min_register_duration" json:"min_register_duration"`
		MaxFreeDomainRegisterDuration string `yaml:"max_free_domain_register_duration" json:"max_free_domain_register_duration"`
		TopDomain                     string `yaml:"top_domain" json:"top_domain"`
	} `yaml:"controller" json:"controller"`
}

type SwaggerOption struct {
	ServerHostIp string `yaml:"server_host_ip"`
}

func NewDefaultApp(home string, appFilePath string, config Config) error {
	appType := &App{
		Home:   home,
		Config: config,
	}

	err := SaveAppFile(appFilePath, appType)
	if err != nil {
		return err
	}

	return nil
}
