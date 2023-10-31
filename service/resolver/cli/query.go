package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Moonyongjung/xns-gateway/cmd/param"
	"github.com/Moonyongjung/xns-gateway/service/resolver"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	"github.com/spf13/cobra"
)

func QueryResolverCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "resolver",
		Aliases: []string{"rs"},
		Short:   "query to the resolver",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if xnsContext.GetApp().Contracts.ResolverContract.ContractAddress == "" {
				return util.LogErr(types.ErrHaveNoAddress, "Resolver")
			}

			return nil
		},
	}

	cmd.AddCommand(
		getConfigCmd(xnsContext),
	)
	return cmd
}

func getConfigCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "get config of the resolver",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query resolver config
$ %s q rs config
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := resolver.QueryConfig(xnsContext)
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
