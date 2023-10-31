package cli

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Moonyongjung/xns-gateway/cmd/param"
	"github.com/Moonyongjung/xns-gateway/types"
	contracttypes "github.com/Moonyongjung/xns-gateway/types/contract"
	"github.com/Moonyongjung/xns-gateway/util"
	"github.com/spf13/cobra"
	xtypes "github.com/xpladev/xpla.go/types"
)

func storeAllContractCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "store-all",
		Short:   "Store all xns contracts",
		Args:    param.WithUsage(cobra.NoArgs),
		Aliases: []string{"s"},
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute contract store-all
$ %s e c s
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			xplac := xnsContext.GetXplaClient()
			home := xnsContext.GetApp().Home

			var txResponses []contracttypes.StoreResponseInfo
			for _, contract := range types.XnsContracts {
				contractPath := filepath.Join(types.DefaultContractWasmDir, contract)
				storeMsg := xtypes.StoreMsg{
					FilePath: contractPath,
				}

				txBytes, err := xplac.StoreCode(storeMsg).CreateAndSignTx()
				if err != nil {
					return util.LogErr(types.ErrTx, err)
				}

				util.LogWait("send tx to store contract,", contract)

				res, err := xplac.BroadcastBlock(txBytes)
				if err != nil {
					return util.LogErr(types.ErrBroadcast, err)
				}
				util.LogSuccess("OK")

				txResponses = append(txResponses, contracttypes.NewStoreResponseInfo(contract, res))
				xplac.WithSequence("")
			}

			// Save the code ID of the anchor contract in the app.yaml
			err := contracttypes.SetCodeID(filepath.Join(home, types.DefaultAppPath, types.DefaultAppFilePath), txResponses)
			if err != nil {
				return util.LogErr(types.ErrContract, err)
			}

			util.LogSuccess(util.StoreContractsSuccessLog)

			return nil
		},
	}

	return cmd
}
