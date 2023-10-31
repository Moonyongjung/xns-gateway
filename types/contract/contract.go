package contract

import (
	"fmt"

	"github.com/Moonyongjung/xns-gateway/types"
)

func SetCodeID(appFilePath string, txResponses []StoreResponseInfo) error {
	xnsApp, err := types.ReadAppFile(appFilePath)
	if err != nil {
		return err
	}

	for _, res := range txResponses {
		switch {
		case res.Contract == types.XnsBaseContract:
			xnsApp.Contracts.BaseContract.CodeID = res.CodeID

		case res.Contract == types.NftRegistrarContract:
			xnsApp.Contracts.NftRegistrarContract.CodeID = res.CodeID

		case res.Contract == types.RegistrarContract:
			xnsApp.Contracts.RegistrarContract.CodeID = res.CodeID

		case res.Contract == types.ControllerContract:
			xnsApp.Contracts.ControllerContract.CodeID = res.CodeID

		case res.Contract == types.ResolverContract:
			xnsApp.Contracts.ResolverContract.CodeID = res.CodeID

		default:
			return fmt.Errorf("invalid xns contract type")
		}
	}

	err = types.SaveAppFile(appFilePath, xnsApp)
	if err != nil {
		return err
	}

	return nil
}

func SetContractAddress(appFilePath string, txResponses []InstantiateResponseInfo) error {
	xnsApp, err := types.ReadAppFile(appFilePath)
	if err != nil {
		return err
	}

	for _, res := range txResponses {
		switch {
		case res.Contract == types.NftRegistrarContract:
			xnsApp.Contracts.NftRegistrarContract.ContractAddress = res.Address

		case res.Contract == types.RegistrarContract:
			xnsApp.Contracts.RegistrarContract.ContractAddress = res.Address

		case res.Contract == types.ControllerContract:
			xnsApp.Contracts.ControllerContract.ContractAddress = res.Address

		case res.Contract == types.ResolverContract:
			xnsApp.Contracts.ResolverContract.ContractAddress = res.Address

		default:
			return fmt.Errorf("invalid xns contract type")
		}
	}

	err = types.SaveAppFile(appFilePath, xnsApp)
	if err != nil {
		return err
	}

	return nil
}
