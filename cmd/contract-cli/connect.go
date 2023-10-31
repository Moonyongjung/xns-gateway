package cli

import (
	"fmt"
	"strings"

	"github.com/Moonyongjung/xns-gateway/cmd/param"
	nftregistrar "github.com/Moonyongjung/xns-gateway/service/nft-registrar"
	"github.com/Moonyongjung/xns-gateway/service/registrar"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	"github.com/spf13/cobra"
)

func connectContractsCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "connect",
		Short:   "connect contracts by broadcasting setting tx",
		Args:    param.WithUsage(cobra.NoArgs),
		Aliases: []string{"c"},
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute contract connect
$ %s e c c
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			util.LogKV("[NFT-Registrar]", "Broadcast tx to set registrar contract")
			err := nftregistrar.ExecSetRegistrar(xnsContext)
			if err != nil {
				return err
			}
			util.LogKV("[NFT-Registrar]", "success - set registrar contract")
			xnsContext.WithXplaClient(xnsContext.GetXplaClient().WithSequence(""))

			util.LogKV("[NFT-Registrar]", "Broadcast tx to set resolver contract")
			err = nftregistrar.ExecSetResolver(xnsContext)
			if err != nil {
				return err
			}
			util.LogKV("[NFT-Registrar]", "success - set resolver contract")
			xnsContext.WithXplaClient(xnsContext.GetXplaClient().WithSequence(""))

			topDomain := xnsContext.GetApp().Config.ContractsOption.Controller.TopDomain
			util.LogKV("[Registrar]", "Broadcast tx to set controller contract")
			err = registrar.ExecSetController(xnsContext, topDomain)
			if err != nil {
				return err
			}
			util.LogKV("[Registrar]", "success - set controller contract")
			xnsContext.WithXplaClient(xnsContext.GetXplaClient().WithSequence(""))

			util.LogKV("[Registrar]", "Broadcast tx to set resolver contract")
			err = registrar.ExecSetResolver(xnsContext)
			if err != nil {
				return err
			}
			util.LogKV("[Registrar]", "success - set resolver contract")

			util.LogSuccess("Finished")

			return nil
		},
	}

	return cmd
}
