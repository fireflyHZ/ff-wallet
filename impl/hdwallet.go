package impl

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/tyler-smith/go-bip39"
)

//var TronBytePrefix = byte(0x41)

//func createAccount(mnemonic string, pathStr string) (accounts.Account, error) {
//	//var account accounts.Account
//	path := parseDerivePath(pathStr)
//	masterKey, err := newFromMnemonic(mnemonic)
//	if err != nil {
//		return accounts.Account{}, err
//	}
//	return derive(masterKey, path)
//}

//func derive(key *hdkeychain.ExtendedKey, path accounts.DerivationPath) (accounts.Account, error) {
//	address, err := deriveAddr(key, path)
//	if err != nil {
//		return accounts.Account{}, err
//	}
//	account := accounts.Account{
//		Address: address,
//		URL: accounts.URL{
//			Scheme: "",
//			Path:   path.String(),
//		},
//	}
//	return account, nil
//}

//func deriveAddr(key *hdkeychain.ExtendedKey, path accounts.DerivationPath) (common.Address, error) {
//	pubkey, err := derivePubkey(key, path)
//	if err != nil {
//		return common.Address{}, err
//	}
//	address := crypto.PubkeyToAddress(*pubkey)
//	return address, nil
//}

//func derivePubkey(key *hdkeychain.ExtendedKey, path accounts.DerivationPath) (*ecdsa.PublicKey, error) {
//	priKey, err := derivePrikey(key, path)
//	if err != nil {
//		return nil, err
//	}
//
//	publicKey := priKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		return nil, errors.New("failed to get public key")
//	}
//
//	return publicKeyECDSA, nil
//}

func derivePrikey(key *hdkeychain.ExtendedKey, path accounts.DerivationPath) (*ecdsa.PrivateKey, error) {
	var err error
	for _, n := range path {
		//key, err = key.Child(n)
		key, err = key.Derive(n)
		if err != nil {
			return nil, err
		}
	}

	privateKey, err := key.ECPrivKey()
	if err != nil {
		return nil, err
	}
	privateKeyECDSA := privateKey.ToECDSA()
	fmt.Println()
	return privateKeyECDSA, nil
}

func derivePrikeyBytes(key *hdkeychain.ExtendedKey, path accounts.DerivationPath) ([]byte, error) {
	var err error
	for _, n := range path {
		key, err = key.Derive(n)
		if err != nil {
			return []byte{}, err
		}
	}

	privateKey, err := key.ECPrivKey()
	if err != nil {
		return []byte{}, err
	}
	//privateKeyECDSA := privateKey.ToECDSA()
	return privateKey.Serialize(), nil
}

func newFromMnemonic(mnemonic string) (*hdkeychain.ExtendedKey, error) {
	if mnemonic == "" {
		return nil, errors.New("mnemonic is required")
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("mnemonic is invalid")
	}
	seed := bip39.NewSeed(mnemonic, "")
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)

	if err != nil {
		return nil, err
	}
	return masterKey, nil
}

func getPrivateKey(mnemonic string, pathStr string) (*ecdsa.PrivateKey, error) {
	path := parseDerivePath(pathStr)
	masterKey, err := newFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	return derivePrikey(masterKey, path)
}
func getPrivateKeyBytes(mnemonic string, pathStr string) ([]byte, error) {
	path := parseDerivePath(pathStr)
	masterKey, err := newFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	return derivePrikeyBytes(masterKey, path)
}
func parseDerivePath(path string) accounts.DerivationPath {
	parsed, err := accounts.ParseDerivationPath(path)
	if err != nil {
		panic(err)
	}

	return parsed
}
