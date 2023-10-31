package cmd

import (
	"context"
	"os"
	"path"

	"github.com/Moonyongjung/xns-gateway/cmd/client"
	contractcli "github.com/Moonyongjung/xns-gateway/cmd/contract-cli"
	controllercli "github.com/Moonyongjung/xns-gateway/service/controller/cli"
	nftregistrarcli "github.com/Moonyongjung/xns-gateway/service/nft-registrar/cli"
	registrarcli "github.com/Moonyongjung/xns-gateway/service/registrar/cli"
	resolvercli "github.com/Moonyongjung/xns-gateway/service/resolver/cli"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Start() {
	cobra.EnableCommandSorting = false

	xnsContext := types.NewXnsContext().
		WithViper(viper.New()).
		WithChannels()

	rootCmd := NewRootCmd(xnsContext)
	rootCmd.SilenceUsage = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func NewRootCmd(xnsContext *types.XnsContext) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   types.DefaultAppName,
		Short: "The gateway for XNS(XPLA Name Serivce)",
		Long: util.ToString(`the XNS gateway can provide functions as below

1. Initialize the XNS gateway by reading config.yaml file which an user set
2. Recover the key to send transactions to the XPLA
3. Easily execute and query to the XNS contract by using CLI
4. Run the XNS gateway server
		`, ""),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	rootCmd.AddCommand(
		InitCmd(),
		KeyCmd(xnsContext),
		ExecuteCmd(xnsContext),
		QueryCmd(xnsContext),
		StartCmd(xnsContext),
	)

	return rootCmd
}

func ExecuteCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "execute",
		Aliases: []string{"e"},
		Short:   "Execute XNS",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			xnsContext, err = checkRunInit(xnsContext)
			if err != nil {
				return err
			}

			xplac, err := client.NewXplaClient(xnsContext, true)
			if err != nil {
				return err
			}
			xnsContext.WithXplaClient(xplac)

			return nil
		},
	}

	cmd.AddCommand(
		contractcli.ExecuteContractCmd(xnsContext),
		nftregistrarcli.ExecuteNftRegistrarCmd(xnsContext),
		registrarcli.ExecuteRegistrarCmd(xnsContext),
		controllercli.ExecuteControllerCmd(xnsContext),
	)
	return cmd
}

func QueryCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Query XNS",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			xnsContext, err = checkRunInit(xnsContext)
			if err != nil {
				return err
			}

			xplac, err := client.NewXplaClient(xnsContext, false)
			if err != nil {
				return err
			}
			xnsContext.WithXplaClient(xplac)

			return nil
		},
	}

	cmd.AddCommand(
		nftregistrarcli.QueryNftRegistrarCmd(xnsContext),
		registrarcli.QueryRegistrarCmd(xnsContext),
		controllercli.QueryControllerCmd(xnsContext),
		resolvercli.QueryResolverCmd(xnsContext),
		XnsContractsInfoCmd(xnsContext),
	)
	return cmd
}

func checkRunInit(xnsContext *types.XnsContext) (*types.XnsContext, error) {
	appFileDir := path.Join(types.DefaultHome, types.DefaultAppPath)
	appFilePath := path.Join(appFileDir, types.DefaultAppFilePath)
	if _, err := os.Stat(appFilePath); os.IsNotExist(err) {
		return nil, util.LogErr(types.ErrParseApp, "run init first")
	}

	if err := xnsContext.WithApp(appFilePath); err != nil {
		return nil, util.LogErr(types.ErrParseApp, err)
	}

	return xnsContext, nil
}
