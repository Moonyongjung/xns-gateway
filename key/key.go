package key

import (
	"os"
	"syscall"

	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/go-bip39"
	"github.com/xpladev/xpla.go/key"
	"golang.org/x/term"
)

func GenerateKey(xnsContext *types.XnsContext) (string, error) {
	_, appFilePath := types.AppFilePath()
	if _, err := os.Stat(appFilePath); os.IsNotExist(err) {
		return "", util.LogErr(types.ErrKey, "run init before gen key")
	}

	keyFileDir, keyFilePath := types.KeyFilePath()
	if _, err := os.Stat(keyFileDir); os.IsNotExist(err) {
		if err := os.Mkdir(keyFileDir, os.ModePerm); err != nil {
			return "", util.LogErr(types.ErrKey, err)
		}
	}

	util.LogInfo("input BIP39 mnemonic\n")
	mnemonicByte, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", util.LogErr(types.ErrKey, err)
	}

	mnemonic := string(mnemonicByte)

	if !bip39.IsMnemonicValid(mnemonic) {
		return "", util.LogErr(types.ErrKey, "invalid mnemonic")
	}

	var passphrase string
	if !xnsContext.GetApp().Config.XnsGateway.KeyTest {
		for {
			util.LogInfo("set passphrase\n")
			passphraseByte, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return "", util.LogErr(types.ErrKey, err)

			}

			util.LogInfo("re-input passphrase to confirm")
			confirm, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return "", util.LogErr(types.ErrKey, err)
			}

			if string(passphraseByte) == string(confirm) {
				passphrase = string(passphraseByte)
				break
			} else {
				util.LogWarning("not match passphrase")
			}
		}
	}

	privateKey, err := key.NewPrivKey(mnemonic)
	if err != nil {
		return "", util.LogErr(types.ErrKey, err)
	}

	addr, err := key.Bech32AddrString(privateKey)
	if err != nil {
		return "", util.LogErr(types.ErrKey, err)
	}

	armorKey := key.EncryptArmorPrivKey(privateKey, passphrase)

	f, err := os.Create(keyFilePath)
	if err != nil {
		return "", util.LogErr(types.ErrKey, err)
	}
	defer f.Close()

	if _, err = f.Write([]byte(armorKey)); err != nil {
		return "", util.LogErr(types.ErrKey, err)
	}

	return addr, nil
}

func ExtractKey(xnsContext *types.XnsContext, home string) (cryptotypes.PrivKey, error) {
	_, keyFilePath := types.KeyFilePath()
	if _, err := os.Stat(keyFilePath); os.IsNotExist(err) {
		return nil, util.LogErr(types.ErrKey, "run recovering key before execute")
	}

	keyFile, err := os.ReadFile(keyFilePath)
	if err != nil {
		return nil, util.LogErr(types.ErrKey, err)
	}

	passphraseByte := []byte("")
	if !xnsContext.GetApp().Config.XnsGateway.KeyTest {
		util.LogInfo("input passphrase")
		passphraseByte, err = term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return nil, util.LogErr(types.ErrKey, err)
		}
	}

	privateKey, _, err := key.UnarmorDecryptPrivKey(string(keyFile), string(passphraseByte))
	if err != nil {
		return nil, util.LogErr(types.ErrKey, err)
	}

	return privateKey, nil
}
