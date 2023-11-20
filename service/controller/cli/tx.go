package cli

import (
	"fmt"
	"strings"

	"github.com/Moonyongjung/xns-gateway/cmd/flag"
	"github.com/Moonyongjung/xns-gateway/cmd/param"
	"github.com/Moonyongjung/xns-gateway/request"
	"github.com/Moonyongjung/xns-gateway/service/controller"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	"github.com/spf13/cobra"
	xtypes "github.com/xpladev/xpla.go/types"
	xutil "github.com/xpladev/xpla.go/util"
)

func ExecuteControllerCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "controller",
		Aliases: []string{"ct"},
		Short:   "execute to the controller",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if xnsContext.GetApp().Contracts.ControllerContract.ContractAddress == "" {
				return util.LogErr(types.ErrHaveNoAddress, "Controller")
			}

			return nil
		},
	}

	cmd.AddCommand(
		controllerSetConfigCmd(xnsContext),
		registerCmd(xnsContext),
		renewCmd(xnsContext),
		enrollSubdomainCmd(xnsContext),
		saveDomainData(xnsContext),
		withdrawAllCmd(xnsContext),
		withdrawCmd(xnsContext),
	)
	return cmd
}

func controllerSetConfigCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-config",
		Short: "set configuration of the controller contract",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute controller set-config
$ %s e ct set-config
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			topPrice, err := cmd.Flags().GetUint64(flag.FlagTopPrice)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			middlePrice, err := cmd.Flags().GetUint64(flag.FlagMiddlePrice)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			lowPrice, err := cmd.Flags().GetUint64(flag.FlagLowPrice)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			minRegisterDuration, err := cmd.Flags().GetUint64(flag.FlagMinRegisterDuration)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			registrarAddress, err := cmd.Flags().GetString(flag.FlagRegistrarAddress)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			msg, err := controller.NewControllerSetConfigExecuteMsg(
				flag.TopPriceI(topPrice),
				flag.MiddlePriceI(middlePrice),
				flag.LowPriceI(lowPrice),
				flag.MinRegisterDurationI(minRegisterDuration),
				flag.RegistrarAddressI(registrarAddress),
			)
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := executeController(xnsContext, msg, types.ZeroAmount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}
	cmd.Flags().Uint64(flag.FlagTopPrice, 0, "set config new top price")
	cmd.Flags().Uint64(flag.FlagMiddlePrice, 0, "set config new middle price")
	cmd.Flags().Uint64(flag.FlagLowPrice, 0, "set config new low price")
	cmd.Flags().Uint64(flag.FlagMinRegisterDuration, 0, "set config new min register duration")
	cmd.Flags().String(flag.FlagRegistrarAddress, "", "set config new registrar contract address")

	return cmd
}

func registerCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "register domain",
		Args:  param.WithUsage(cobra.ExactArgs(4)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute controller register [amount] [label] [expire-duration] [secret]
$ %s e ct register [amount] [label] [expire-duration] [secret]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			amount, label, duration, secret := args[0], args[1], args[2], args[3]

			durationU64, err := xutil.FromStringToUint64(duration)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			msg, err := controller.NewControllerDomainRegisterExecuteMsg(label, secret, durationU64)
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := executeController(xnsContext, msg, amount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}

	return cmd
}

func renewCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew",
		Short: "renew domain",
		Args:  param.WithUsage(cobra.ExactArgs(3)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute controller renew [amount] [label] [expire-duration]
$ %s e ct renew [amount] [label] [expire-duration]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			amount, label, duration := args[0], args[1], args[2]

			durationU64, err := xutil.FromStringToUint64(duration)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			msg, err := controller.NewControllerDomainRenewExecuteMsg(label, durationU64)
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := executeController(xnsContext, msg, amount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}

	return cmd
}

func enrollSubdomainCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll-subdomain",
		Short: "enroll subdomain of the domain",
		Args:  param.WithUsage(cobra.ExactArgs(2)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute controller enroll-subdomain [label] [subdomain-label]
$ %s e ct enroll-subdomain [label] [subdomain-label]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			domain, subdomain := args[0], args[1]

			textData, err := cmd.Flags().GetString(flag.FlagSubdomainTextData)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			accountData, err := cmd.Flags().GetString(flag.FlagSubdomainAccountData)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			publicKeyData, err := cmd.Flags().GetString(flag.FlagPublicKeyData)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			contentHashData, err := cmd.Flags().GetString(flag.FlagContentHashData)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			if textData == "" && accountData == "" && publicKeyData == "" && contentHashData == "" {
				return util.LogErr(types.ErrInvalidRequest, "select mapping data at least one (e.g. --account / --text)")
			}

			msg, err := controller.NewControllerEnrollSubdomainExecuteMsg(domain, subdomain, textData, accountData, publicKeyData, contentHashData)
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := executeController(xnsContext, msg, types.ZeroAmount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}
	cmd.Flags().String(flag.FlagSubdomainTextData, "", "mapping text data of the subdomain")
	cmd.Flags().String(flag.FlagSubdomainAccountData, "", "mapping account address data of the subdomain")
	cmd.Flags().String(flag.FlagPublicKeyData, "", "mapping public key data of the subdomain")
	cmd.Flags().String(flag.FlagContentHashData, "", "mapping content hash data of the subdomain")

	return cmd
}

func saveDomainData(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save-domain-data",
		Short: "save data in the domain",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute controller save-domain-data [label]
$ %s e ct save-domain-data [label]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			domain := args[0]

			textData, err := cmd.Flags().GetString(flag.FlagSubdomainTextData)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			accountData, err := cmd.Flags().GetString(flag.FlagSubdomainAccountData)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			publicKeyData, err := cmd.Flags().GetString(flag.FlagPublicKeyData)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			contentHashData, err := cmd.Flags().GetString(flag.FlagContentHashData)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			if textData == "" && accountData == "" && publicKeyData == "" && contentHashData == "" {
				return util.LogErr(types.ErrInvalidRequest, "select mapping data at least one (e.g. --account / --text)")
			}

			msg, err := controller.NewControllerSaveDomainDataExecuteMsg(domain, textData, accountData, publicKeyData, contentHashData)
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := executeController(xnsContext, msg, types.ZeroAmount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}
	cmd.Flags().String(flag.FlagSubdomainTextData, "", "mapping text data of the subdomain")
	cmd.Flags().String(flag.FlagSubdomainAccountData, "", "mapping account address data of the subdomain")
	cmd.Flags().String(flag.FlagPublicKeyData, "", "mapping public key data of the subdomain")
	cmd.Flags().String(flag.FlagContentHashData, "", "mapping content hash data of the subdomain")

	return cmd
}

func withdrawAllCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-all",
		Short: "withdraw all amount of the controller",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute controller withdraw-all
$ %s e ct withdraw-all
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := controller.NewControllerWithdrawAllExecuteMsg()
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := executeController(xnsContext, msg, types.ZeroAmount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}

	return cmd
}

func withdrawCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw",
		Short: "withdraw amount of the controller",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute controller withdraw [amount]
$ %s e ct withdraw [amount]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			amountBigInt, err := xutil.FromStringToBigInt(args[0])
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			msg, err := controller.NewControllerWithdrawExecuteMsg(amountBigInt)
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := executeController(xnsContext, msg, types.ZeroAmount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}

	return cmd
}

func executeController(xnsContext *types.XnsContext, msg, amount string) (*xtypes.TxRes, error) {
	address := xnsContext.GetApp().Contracts.ControllerContract.ContractAddress
	return request.ExecuteContract(xnsContext, address, msg, amount)
}
