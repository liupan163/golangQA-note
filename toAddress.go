package main

import (
	"encoding/hex"
	"github.com/edunuzzi/go-bip44"
)

func  toAddress()  {
	/*entropy, _ := bip39.NewEntropy(160)
	mnemonic1, _ := bip39.NewMnemonic(entropy)
	seed := bip39.NewSeed(mnemonic1, "")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic1)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)*/

	mnemonic, _ := bip44.NewMnemonic(160)
	mnemonic.Value = "label situate argue grape purchase push cat stuff search health hand excess"
	println("mnemonic is =====>", mnemonic.Value)
	seedBytes, _ := mnemonic.NewSeed("")
	encodeStr := hex.EncodeToString(seedBytes)
	println("encodeStr is =====>", encodeStr)
	xKey, err := bip44.NewKeyFromSeedHex(encodeStr, bip44.MAINNET)
	if err != nil {
		println("err is =====>", err)
		return
	}
	accountKey, _ := xKey.BIP44AccountKey(bip44.BitcoinCoinType, 0, true)
	externalAddress, _ := accountKey.DeriveP2PKAddress(bip44.ExternalChangeType, 1, bip44.MAINNET)
	println("externalAddress is =====>", externalAddress)
}
