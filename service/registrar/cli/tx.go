package cli

import (
	"fmt"
	"strings"

	"github.com/Moonyongjung/xns-gateway/cmd/param"
	"github.com/Moonyongjung/xns-gateway/service/registrar"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	"github.com/spf13/cobra"
)

func ExecuteRegistrarCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "registrar",
		Aliases: []string{"r"},
		Short:   "execute to the registrar",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if xnsContext.GetApp().Contracts.RegistrarContract.ContractAddress == "" {
				return util.LogErr(types.ErrHaveNoAddress, "Registrar")
			}

			return nil
		},
	}

	cmd.AddCommand(
		setControllerCmd(xnsContext),
		removeControllerCmd(xnsContext),
		setResolverCmd(xnsContext),
	)
	return cmd
}

func setControllerCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-controller",
		Short: "set controller on the registrar contract",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute registrar set-controller [top_domain]
$ %s e r set-controller [top_domain]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			topDomain := args[0]
			return registrar.ExecSetController(xnsContext, topDomain)
		},
	}

	return cmd
}

func removeControllerCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-controller",
		Short: "remove controller on the registrar contract",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute registrar remove-controller [controller-address]
$ %s e r remove-controller [controller-address]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := registrar.NewRegistrarRemoveControllerExecuteMsg(args[0])
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := registrar.ExecuteRegistrar(xnsContext, msg, types.ZeroAmount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}

	return cmd
}

func setResolverCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-resolver",
		Short: "set resolver on the registrar contract",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute registrar set-resolver
$ %s e r set-resolver
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return registrar.ExecSetResolver(xnsContext)
		},
	}

	return cmd
}
