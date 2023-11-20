package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Moonyongjung/xns-gateway/cmd/param"
	"github.com/Moonyongjung/xns-gateway/service/controller"
	"github.com/Moonyongjung/xns-gateway/types"

	"github.com/Moonyongjung/xns-gateway/util"
	"github.com/spf13/cobra"
)

func QueryControllerCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "controller",
		Aliases: []string{"ct"},
		Short:   "query to the controller",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if xnsContext.GetApp().Contracts.ControllerContract.ContractAddress == "" {
				return util.LogErr(types.ErrHaveNoAddress, "Controller")
			}

			return nil
		},
	}

	cmd.AddCommand(
		controllerBalanceCmd(xnsContext),
		getControllerConfigCmd(xnsContext),
		calcRegisterCostCmd(xnsContext),
		topDomainCmd(xnsContext),
	)
	return cmd
}

func controllerBalanceCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balance",
		Short: "balance of the controller",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query controller balance
$ %s q ct balance
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := controller.QueryBalance(xnsContext)
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

func getControllerConfigCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "get config of the controller",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query controller config
$ %s q ct config
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := controller.QueryConfig(xnsContext)
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

func calcRegisterCostCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-cost",
		Short: "calculate register amount of the domain",
		Args:  param.WithUsage(cobra.ExactArgs(2)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query controller register-cost [label] [expire-duration]
$ %s q ct register-cost [label] [expire-duration]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			label, duration := args[0], args[1]

			res, err := controller.QueryRegisterCost(xnsContext, label, duration)
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

func topDomainCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "top-domain",
		Short: "get top domain of the controller",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query controller top-domain
$ %s q ct top-domain
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := controller.QueryTopDomain(xnsContext)
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
