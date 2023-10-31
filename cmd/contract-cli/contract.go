package cli

import (
	"github.com/Moonyongjung/xns-gateway/types"

	"github.com/spf13/cobra"
)

func ExecuteContractCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "contract",
		Aliases: []string{"c"},
		Short:   "execute to the XNS contract",
	}

	cmd.AddCommand(
		storeAllContractCmd(xnsContext),
		instantiateAllContractCmd(xnsContext),
		connectContractsCmd(xnsContext),
	)
	return cmd
}
