package cli

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Moonyongjung/xns-gateway/cmd/param"
	"github.com/Moonyongjung/xns-gateway/messages"
	"github.com/Moonyongjung/xns-gateway/types"
	contracttypes "github.com/Moonyongjung/xns-gateway/types/contract"
	"github.com/Moonyongjung/xns-gateway/util"
	"github.com/spf13/cobra"
	"github.com/xpladev/xpla.go/key"
	xtypes "github.com/xpladev/xpla.go/types"
	"golang.org/x/exp/slices"
)

func instantiateAllContractCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "instantiate-all",
		Short:   "Instantiate all xns wasm contracts",
		Args:    param.WithUsage(cobra.NoArgs),
		Aliases: []string{"i"},
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s execute contract instantiate-all
$ %s e c i
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			xplac := xnsContext.GetXplaClient()
			// app := types.XNSApp()
			home := xnsContext.GetApp().Home

			addr, err := key.Bech32AddrString(xplac.GetPrivateKey())
			if err != nil {
				return util.LogErr(types.ErrKey, err)
			}

			codeIDs := []string{
				xnsContext.GetApp().Contracts.NftRegistrarContract.CodeID,
				xnsContext.GetApp().Contracts.RegistrarContract.CodeID,
				xnsContext.GetApp().Contracts.ControllerContract.CodeID,
				xnsContext.GetApp().Contracts.ResolverContract.CodeID,
			}

			if slices.Contains(codeIDs, "") {
				return util.LogErr(types.ErrParseApp, "code ID of contract is empty")
			}

			xnsBaseContractCodeID := xnsContext.GetApp().Contracts.BaseContract.CodeID

			var txResponses []contracttypes.InstantiateResponseInfo
			var dependentNftRegistrarAddress string
			var dependentRegistrarAddress string
			for i, codeID := range codeIDs {
				var instMsg string
				var contract string

				switch {
				// NFT registrar contract
				case i == 0:
					contract = types.NftRegistrarContract
					instMsg = messages.NewNftRegistrarInstantiateMsg(xnsContext.GetApp(), xnsBaseContractCodeID)

				// Registrar contract
				case i == 1:
					contract = types.RegistrarContract
					instMsg = messages.NewRegistrarInstantiateMsg(xnsContext.GetApp(), dependentNftRegistrarAddress)

				// Controller contract
				case i == 2:
					contract = types.ControllerContract
					instMsg = messages.NewControllerInstantiateMsg(xnsContext.GetApp(), dependentRegistrarAddress)

				// Resolver contract
				case i == 3:
					contract = types.ResolverContract
					instMsg = messages.NewResolverInstantiateMsg(xnsContext.GetApp(), dependentRegistrarAddress, dependentNftRegistrarAddress)
				}

				instantiateMsg := xtypes.InstantiateMsg{
					CodeId:  codeID,
					Amount:  "0",
					Label:   "instantiate contract",
					InitMsg: instMsg,
					Admin:   addr,
				}

				txBytes, err := xplac.InstantiateContract(instantiateMsg).CreateAndSignTx()
				if err != nil {
					return util.LogErr(types.ErrTx, err)
				}

				util.LogWait("send tx to instantiate contract,", strings.Split(contract, "-")[0])
				util.LogKV("code ID", codeID)

				res, err := xplac.BroadcastBlock(txBytes)
				if err != nil {
					return util.LogErr(types.ErrBroadcast, err)
				}

				addr := res.Response.Logs[0].Events[0].Attributes[0].Value
				switch {
				case i == 0:
					dependentNftRegistrarAddress = addr

				case i == 1:
					dependentRegistrarAddress = addr
				}

				txResponses = append(txResponses, contracttypes.NewInstantiateResponseInfo(contract, res))
				xplac.WithSequence("")
			}

			err = contracttypes.SetContractAddress(filepath.Join(home, types.DefaultAppPath, types.DefaultAppFilePath), txResponses)
			if err != nil {
				return util.LogErr(types.ErrContract, err)
			}

			util.LogSuccess(util.InstantiatecontractsSuccessLog)

			return nil
		},
	}

	return cmd
}
