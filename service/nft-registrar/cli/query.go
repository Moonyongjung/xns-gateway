package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Moonyongjung/xns-gateway/cmd/flag"
	"github.com/Moonyongjung/xns-gateway/cmd/param"
	nftregistrar "github.com/Moonyongjung/xns-gateway/service/nft-registrar"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"
	"github.com/spf13/cobra"
)

func QueryNftRegistrarCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nft-registrar",
		Aliases: []string{"n"},
		Short:   "query to the nft registrar",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if xnsContext.GetApp().Contracts.NftRegistrarContract.ContractAddress == "" {
				return util.LogErr(types.ErrHaveNoAddress, "NFT-Registrar")
			}

			return nil
		},
	}

	cmd.AddCommand(
		getRegistrarCmd(xnsContext),
		getResolverCmd(xnsContext),
		accountByDomainCmd(xnsContext),
		domainByAccountCmd(xnsContext),
		domainHistoryByAccountCmd(xnsContext),
		expiredLatestDomainCmd(xnsContext),
		tokenIDByDomainCmd(xnsContext),
		subdomainCmd(xnsContext),
		subdomainsCmd(xnsContext),
		primaryDomainCmd(xnsContext),
		hashResultCmd(xnsContext),
		contractInfoCmd(xnsContext),
		nftInfoCmd(xnsContext),
		ownerOfCmd(xnsContext),
		allNftInfoCmd(xnsContext),
		allOperatorsCmd(xnsContext),
		numTokensCmd(xnsContext),
		tokensCmd(xnsContext),
		allTokensCmd(xnsContext),
		approvalCmd(xnsContext),
		approvalsCmd(xnsContext),
		minterCmd(xnsContext),
	)
	return cmd
}

func getRegistrarCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registrar",
		Short: "get connected registrar contract of nft-registrar",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar registrar
$ %s q n registrar
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := nftregistrar.QueryRegistrar(xnsContext)
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

func getResolverCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolver",
		Short: "get connected resolver contract of nft-registrar",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar resolver
$ %s q n resolver
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := nftregistrar.QueryResolver(xnsContext)
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

func accountByDomainCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account-by-domain",
		Short: "get the account address by using the domain",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar account-by-domain [domain]
$ %s q n account-by-domain [domain]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			domain := args[0]

			second, top, err := splitedDomain(domain)
			if err != nil {
				return err
			}

			res, err := nftregistrar.QueryAccountByDomain(xnsContext, second, top)
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

func domainByAccountCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "domain-by-account",
		Short: "get domains by using the account address",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar domain-by-account [account-address]
$ %s q n domain-by-account [account-address]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			account := args[0]
			res, err := nftregistrar.QueryDomainByAccount(xnsContext, account)
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

func domainHistoryByAccountCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "domain-history-by-account",
		Short: "get domain history by using the account address",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar domain-history-by-account [account-address]
$ %s q n domain-history-by-account [account-address]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			account := args[0]
			res, err := nftregistrar.QueryDomainHistoryByAccount(xnsContext, account)
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

func expiredLatestDomainCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "expired-latest-domain",
		Short: "get expired latest domain",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar expired-latest-domain [domain]
$ %s q n expired-latest-domain [domain]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			domain := args[0]
			second, top, err := splitedDomain(domain)
			if err != nil {
				return err
			}

			res, err := nftregistrar.QueryExpiredLatestDomain(xnsContext, second, top)
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

func tokenIDByDomainCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokenid-by-domain",
		Short: "get token ID by using domain",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar tokenid-by-domain [domain]
$ %s q n tokenid-by-domain [domain]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			domain := args[0]
			second, top, err := splitedDomain(domain)
			if err != nil {
				return err
			}

			res, err := nftregistrar.QueryTokenIDByDomain(xnsContext, second, top)
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

func subdomainCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subdomain",
		Short: "get subdomain mapped data",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar subdomain [subdomain]
$ %s q n subdomain [subdomain]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			subdomain := args[0]
			sub, second, top, err := splitedSubdomain(subdomain)
			if err != nil {
				return err
			}

			res, err := nftregistrar.QuerySubdomain(xnsContext, sub, second, top)
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

func subdomainsCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subdomains",
		Short: "get subdomain list of the domain",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar subdomains [domain]
$ %s q n subdomains [domain]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			domain := args[0]
			second, top, err := splitedDomain(domain)
			if err != nil {
				return err
			}

			res, err := nftregistrar.QuerySubdomains(xnsContext, second, top)
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

func primaryDomainCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "primary-domain",
		Short: "get primary domain of the account",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar primary-domain [account-address]
$ %s q n primary-domain [account-address]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			account := args[0]

			res, err := nftregistrar.QueryPrimaryDomain(xnsContext, account)
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

func hashResultCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hash",
		Short: "get hash result for input",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar hash [input]
$ %s q n hash [input]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			input := args[0]

			res, err := nftregistrar.QueryHashResult(xnsContext, input)
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

func contractInfoCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contract-info",
		Short: "get contract info",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar contract-info
$ %s q n contract-info
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := nftregistrar.QueryContractInfo(xnsContext)
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

func nftInfoCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft-info",
		Short: "get nft info by token ID",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar nft-info [token-id]
$ %s q n nft-info [token-id]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			tokenID := args[0]
			res, err := nftregistrar.QueryNftInfo(xnsContext, tokenID)
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

func ownerOfCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "owner-of",
		Short: "get owner of token ID",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar owner-of [token-id]
$ %s q n owner-of [token-id]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			tokenID := args[0]
			includeExpired, err := cmd.Flags().GetBool(flag.FlagIncludeExpired)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			res, err := nftregistrar.QueryOwnerOf(xnsContext, tokenID, includeExpired)
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
	cmd.Flags().Bool(flag.FlagIncludeExpired, false, "expired option")

	return cmd
}

func allNftInfoCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-nft-info",
		Short: "get all nft info of the token ID",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar all-nft-info [token-id]
$ %s q n all-nft-info [token-id]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			tokenID := args[0]
			includeExpired, err := cmd.Flags().GetBool(flag.FlagIncludeExpired)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			res, err := nftregistrar.QueryAllNftInfo(xnsContext, tokenID, includeExpired)
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
	cmd.Flags().Bool(flag.FlagIncludeExpired, false, "expired option")

	return cmd
}

func allOperatorsCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-operators",
		Short: "get all operators of the owner",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar all-operators [owner]
$ %s q n all-operators [owner]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]

			includeExpired, err := cmd.Flags().GetBool(flag.FlagIncludeExpired)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			startAfter, err := cmd.Flags().GetString(flag.FlagStarAfter)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			limit, err := cmd.Flags().GetString(flag.FlagLimit)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			res, err := nftregistrar.QueryAllOperators(xnsContext, owner, includeExpired, startAfter, limit)
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
	cmd.Flags().Bool(flag.FlagIncludeExpired, false, "include expired option")
	cmd.Flags().String(flag.FlagStarAfter, "", "query range (start after)")
	cmd.Flags().String(flag.FlagLimit, "", "query range (limit)")

	return cmd
}

func numTokensCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "num-tokens",
		Short: "get the number of tokens",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar num-tokens
$ %s q n num-tokens
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := nftregistrar.QueryNumTokens(xnsContext)
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

func tokensCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokens",
		Short: "get token list of the owner",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar tokens [owner]
$ %s q n tokens [owner]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]

			startAfter, err := cmd.Flags().GetString(flag.FlagStarAfter)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			limit, err := cmd.Flags().GetString(flag.FlagLimit)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			res, err := nftregistrar.QueryTokens(xnsContext, owner, startAfter, limit)
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

func allTokensCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-tokens",
		Short: "get all tokens of the contract",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar all-tokens
$ %s q n all-tokens
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

			res, err := nftregistrar.QueryAllTokens(xnsContext, startAfter, limit)
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

func approvalCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approval",
		Short: "get approval of the token ID",
		Args:  param.WithUsage(cobra.ExactArgs(2)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar approval [token-id] [spender]
$ %s q n approval [token-id] [spender]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			tokenID, spender := args[0], args[1]

			includeExpired, err := cmd.Flags().GetBool(flag.FlagIncludeExpired)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			res, err := nftregistrar.QueryApproval(xnsContext, tokenID, spender, includeExpired)
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
	cmd.Flags().Bool(flag.FlagIncludeExpired, false, "include expired option")

	return cmd
}

func approvalsCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approvals",
		Short: "get approvals of the token ID",
		Args:  param.WithUsage(cobra.ExactArgs(1)),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar approvals [token-id]
$ %s q n approvals [token-id]
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			tokenID := args[0]

			includeExpired, err := cmd.Flags().GetBool(flag.FlagIncludeExpired)
			if err != nil {
				return util.LogErr(types.ErrParseData, err)
			}

			res, err := nftregistrar.QueryApprovals(xnsContext, tokenID, includeExpired)
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
	cmd.Flags().Bool(flag.FlagIncludeExpired, false, "include expired option")

	return cmd
}

func minterCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "minter",
		Short: "get minter of the contract",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query nft-registrar minter
$ %s q n minter
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := nftregistrar.QueryMinter(xnsContext)
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

func splitedDomain(domain string) (string, string, error) {
	if !strings.Contains(domain, ".") {
		return "", "", util.LogErr(types.ErrParseData, "input domain (e.g. abc.xpla)")
	}

	splited := strings.Split(domain, ".")
	second, top := splited[0], splited[1]

	return second, top, nil
}

func splitedSubdomain(subdomain string) (string, string, string, error) {
	if !strings.Contains(subdomain, ".") {
		return "", "", "", util.LogErr(types.ErrParseData, "input subdomain (e.g. abc.def.xpla)")
	}

	splited := strings.Split(subdomain, ".")
	if len(splited) != 3 {
		return "", "", "", util.LogErr(types.ErrParseData, "input subdomain (e.g. abc.def.xpla)")
	}

	sub, second, top := splited[0], splited[1], splited[2]

	return sub, second, top, nil
}
