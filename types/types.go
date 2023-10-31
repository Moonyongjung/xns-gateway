package types

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/xpladev/xpla.go/provider"
	"gopkg.in/yaml.v3"
)

var (
	DefaultAppName         = "xns"
	DefaultHome            = filepath.Join(os.Getenv("HOME"), ".xns")
	DefaultAppPath         = "config"
	DefaultAppFilePath     = "app.yaml"
	DefaultKeyPath         = "key"
	DefaultKeyFilePath     = "xns_gateway_key.pem"
	DefaultConfigPath      = "./config.yaml"
	DefaultContractWasmDir = "./contract/artifacts"

	XnsBaseContract      = "xns_base-aarch64.wasm"
	NftRegistrarContract = "nft_registrar-aarch64.wasm"
	RegistrarContract    = "registrar-aarch64.wasm"
	ControllerContract   = "controller-aarch64.wasm"
	ResolverContract     = "resolver-aarch64.wasm"

	ZeroAmount = "0"
)

var (
	XnsContracts = []string{
		XnsBaseContract,
		NftRegistrarContract,
		RegistrarContract,
		ControllerContract,
		ResolverContract,
	}
)

type XnsContext struct {
	viper      *viper.Viper
	xplaClient provider.XplaClient
	channels   Channels
	app        App
}

func NewXnsContext() *XnsContext {
	return &XnsContext{}
}

func (x *XnsContext) WithViper(viper *viper.Viper) *XnsContext {
	x.viper = viper
	return x
}

func (x *XnsContext) WithXplaClient(xplaClient provider.XplaClient) *XnsContext {
	x.xplaClient = xplaClient
	return x
}

func (x *XnsContext) WithChannels() *XnsContext {
	x.channels = makeChannels()
	return x
}

func (x *XnsContext) WithApp(filePath string) error {
	appType, err := ReadAppFile(filePath)
	if err != nil {
		return err
	}
	x.app = *appType

	return nil
}

func (x *XnsContext) GetViper() *viper.Viper             { return x.viper }
func (x *XnsContext) GetXplaClient() provider.XplaClient { return x.xplaClient }
func (x *XnsContext) GetChannels() Channels              { return x.channels }
func (x *XnsContext) GetApp() App                        { return x.app }

func ReadAppFile(filePath string) (*App, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var appType App

	err = yaml.Unmarshal(yamlFile, &appType)
	if err != nil {
		return nil, err
	}

	return &appType, nil
}

func SaveAppFile(appFilePath string, appType *App) error {
	bytes, err := yaml.Marshal(appType)
	if err != nil {
		return err
	}

	f, err := os.Create(appFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(bytes); err != nil {
		return err
	}

	return nil
}
