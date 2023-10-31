package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Moonyongjung/xns-gateway/cmd/flag"
	"github.com/Moonyongjung/xns-gateway/cmd/param"
	"github.com/Moonyongjung/xns-gateway/service/registrar"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	"github.com/spf13/cobra"
)

func QueryRegistrarCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "registrar",
		Aliases: []string{"r"},
		Short:   "query to the registrar",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if xnsContext.GetApp().Contracts.RegistrarContract.ContractAddress == "" {
				return util.LogErr(types.ErrHaveNoAddress, "Registrar")
			}

			return nil
		},
	}

	cmd.AddCommand(
		getConfigCmd(xnsContext),
		nftRegistrarCmd(xnsContext),
		controllersCmd(xnsContext),
	)
	return cmd
}

func getConfigCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "get config of the registrar",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query registrar config
$ %s q r config
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := registrar.QueryConfig(xnsContext)
			if err != nil {
				return err
			}

			resBytes, err := json.Marshal(*res)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			util.LogInfo(string(resBytes))
			return nil
		},
	}

	return cmd
}

func nftRegistrarCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft-registrar",
		Short: "get nft-registrar mapped with the registrar",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query registrar nft-registrar
$ %s q r nft-registrar
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := registrar.QueryNftRegistrar(xnsContext)
			if err != nil {
				return err
			}

			resBytes, err := json.Marshal(*res)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			util.LogInfo(string(resBytes))
			return nil
		},
	}

	return cmd
}

func controllersCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "controllers",
		Short: "get controllers in registrar",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query registrar controllers
$ %s q r controllers
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			startAfter, err := cmd.Flags().GetString(flag.FlagStarAfter)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			limit, err := cmd.Flags().GetString(flag.FlagLimit)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			res, err := registrar.QueryControllers(xnsContext, startAfter, limit)
			if err != nil {
				return err
			}

			resBytes, err := json.Marshal(*res)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			util.LogInfo(string(resBytes))
			return nil
		},
	}
	cmd.Flags().String(flag.FlagStarAfter, "", "query range (start after)")
	cmd.Flags().String(flag.FlagLimit, "", "query range (limit)")

	return cmd
}
