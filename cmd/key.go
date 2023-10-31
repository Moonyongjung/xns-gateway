package cmd

import (
	"fmt"
	"strings"

	"github.com/Moonyongjung/xns-gateway/cmd/param"
	"github.com/Moonyongjung/xns-gateway/key"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"
	"github.com/spf13/cobra"
)

func KeyCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "key",
		Aliases: []string{"k"},
		Short:   "Manage the private key for the XNS gateway",
	}

	cmd.AddCommand(
		recoverCmd(xnsContext),
	)

	return cmd
}

func recoverCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recover",
		Short: "Recover key by using mnemonic words. the XNS gateway can own only one key",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s key recover
$ %s k recover
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			xnsContext, err = checkRunInit(xnsContext)
			if err != nil {
				return err
			}

			addr, err := key.GenerateKey(xnsContext)
			if err != nil {
				return err
			}

			util.LogKV("gateway address", addr)
			util.LogSuccess(util.RecoverKeySuccessLog)

			return nil
		},
	}

	return cmd
}
