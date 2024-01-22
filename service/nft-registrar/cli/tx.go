package cli

import (
	"fmt"
	"strings"

	"github.com/Moonyongjung/xns-gateway/cmd/flag"
	"github.com/Moonyongjung/xns-gateway/cmd/param"
	nftregistrar "github.com/Moonyongjung/xns-gateway/service/nft-registrar"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	"github.com/spf13/cobra"
)

func ExecuteNftRegistrarCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nft-registrar",
		Aliases: []string{"n"},
		Short:   "execute to the nft registrar",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if xnsContext.GetApp().Contracts.NftRegistrarContract.ContractAddress == "" {
				return util.LogErr(types.ErrHaveNoAddress, "NFT-Registrar")
			}

			return nil
		},
	}

	cmd.AddCommand(
		setRegistrarCmd(xnsContext),
		setResolverCmd(xnsContext),
		approveCmd(xnsContext),
		revokeCmd(xnsContext),
		approveAllCmd(xnsContext),
		revokeAllCmd(xnsContext),
		transferNftCmd(xnsContext),
	)
	return cmd
}

func setRegistrarCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-registrar",
		Short: "set nft-registrar with registrar contract",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute nft-registrar set-registrar
$ %s e n set-registrar
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nftregistrar.ExecSetRegistrar(xnsContext)
		},
	}

	return cmd
}

func setResolverCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-resolver",
		Short: "set nft-registrar with resolver contract",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute nft-registrar set-resolver
$ %s e n set-resolver
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nftregistrar.ExecSetResolver(xnsContext)
		},
	}

	return cmd
}

func approveCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve",
		Short: "approve spender of the token ID",
		Args:  param.WithUsage(cobra.ExactArgs(3)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute nft-registrar approve [granter] [spender] [token-id]
$ %s e n approve [granter] [spender] [token-id]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			granter, spender, tokenID := args[0], args[1], args[2]

			expireAtHeight, err := cmd.Flags().GetString(flag.FlagExpireAtHeight)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			expireAtTime, err := cmd.Flags().GetUint64(flag.FlagExpireAtTime)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			msg, err := nftregistrar.NewNftRegistrarApproveExecuteMsg(
				granter, spender, tokenID,
				flag.ExpireAtHeightI(expireAtHeight),
				flag.ExpireAtTimeI(expireAtTime),
			)
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := nftregistrar.ExecuteNftRegistrar(xnsContext, msg, types.ZeroAmount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}
	cmd.Flags().String(flag.FlagExpireAtHeight, "", "set expire at height")
	cmd.Flags().Uint64(flag.FlagExpireAtTime, 0, "set expire at time")

	return cmd
}

func revokeCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke",
		Short: "revoke spender of the token ID",
		Args:  param.WithUsage(cobra.ExactArgs(3)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute nft-registrar revoke [granter] [spender] [token-id]
$ %s e n revoke [granter] [spender] [token-id]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			granter, spender, tokenID := args[0], args[1], args[2]

			msg, err := nftregistrar.NewNftRegistrarRevokeExecuteMsg(granter, spender, tokenID)
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := nftregistrar.ExecuteNftRegistrar(xnsContext, msg, types.ZeroAmount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}

	return cmd
}

func approveAllCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-all",
		Short: "approve the operator of the granter",
		Args:  param.WithUsage(cobra.ExactArgs(2)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute nft-registrar approve-all [granter] [operator]
$ %s e n approve-all [granter] [operator]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			granter, operator := args[0], args[1]

			expireAtHeight, err := cmd.Flags().GetString(flag.FlagExpireAtHeight)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			expireAtTime, err := cmd.Flags().GetUint64(flag.FlagExpireAtTime)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			msg, err := nftregistrar.NewNftRegistrarApproveAllExecuteMsg(
				granter, operator,
				flag.ExpireAtHeightI(expireAtHeight),
				flag.ExpireAtTimeI(expireAtTime),
			)
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := nftregistrar.ExecuteNftRegistrar(xnsContext, msg, types.ZeroAmount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}
	cmd.Flags().String(flag.FlagExpireAtHeight, "", "set expire at height")
	cmd.Flags().Uint64(flag.FlagExpireAtTime, 0, "set expire at time")

	return cmd
}

func revokeAllCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-all",
		Short: "revoke the operator of the granter",
		Args:  param.WithUsage(cobra.ExactArgs(2)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute nft-registrar revoke-all [granter] [operator]
$ %s e n revoke-all [granter] [operator]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			granter, operator := args[0], args[1]

			msg, err := nftregistrar.NewNftRegistrarRevokeAllExecuteMsg(granter, operator)
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := nftregistrar.ExecuteNftRegistrar(xnsContext, msg, types.ZeroAmount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}

	return cmd
}

func transferNftCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-nft",
		Short: "transfer NFT token ID to the recipient",
		Args:  param.WithUsage(cobra.ExactArgs(3)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute nft-registrar transfer-nft [sender] [recipient] [token-id]
$ %s e n transfer-nft [sender] [recipient] [token-id]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			sender, recipient, tokenID := args[0], args[1], args[2]

			msg, err := nftregistrar.NewNftRegistrarTransferNftExecuteMsg(sender, recipient, tokenID)
			if err != nil {
				return util.LogErr(types.ErrNewMsg, err)
			}

			res, err := nftregistrar.ExecuteNftRegistrar(xnsContext, msg, types.ZeroAmount)
			if err != nil {
				return err
			}

			util.LogInfo(res.Response)

			return nil
		},
	}

	return cmd
}
