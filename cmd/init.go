package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Moonyongjung/xns-gateway/cmd/flag"
	"github.com/Moonyongjung/xns-gateway/cmd/param"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func InitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		Short:   "Initialize the anchor",
		Args:    param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s init 
$ %s i
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFile, err := cmd.Flags().GetString(flag.FlagConfigFile)
			if err != nil {
				return util.LogErr(types.ErrInit, err)
			}

			appFileDir := path.Join(types.DefaultHome, types.DefaultAppPath)
			appFilePath := path.Join(appFileDir, types.DefaultAppFilePath)

			if _, err := os.Stat(appFilePath); os.IsNotExist(err) {
				if _, err := os.Stat(appFileDir); os.IsNotExist(err) {
					if _, err := os.Stat(types.DefaultHome); os.IsNotExist(err) {
						if err = os.Mkdir(types.DefaultHome, os.ModePerm); err != nil {
							return util.LogErr(types.ErrParseConfig, err)
						}
					}
					if err = os.Mkdir(appFileDir, os.ModePerm); err != nil {
						return util.LogErr(types.ErrParseConfig, err)
					}
				}

				yamlFile, err := os.ReadFile(configFile)
				if err != nil {
					return util.LogErr(types.ErrParseConfig, err)
				}

				var configType types.Config
				err = yaml.Unmarshal(yamlFile, &configType)
				if err != nil {
					return util.LogErr(types.ErrParseConfig, err)
				}

				// Generate default app file in the home directory.
				err = types.NewDefaultApp(types.DefaultHome, appFilePath, configType)
				if err != nil {
					return util.LogErr(types.ErrParseApp, err)
				}

				util.LogSuccess(util.InitSuccessLog)
				return nil
			}

			return util.LogErr(types.ErrParseApp, "app file already exists")
		},
	}
	cmd.Flags().String(flag.FlagConfigFile, types.DefaultConfigPath, "configuration file path")

	return cmd
}
