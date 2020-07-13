package main

import (
	"fmt"
	"github.com/edunuzzi/go-bip44"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	toAddress()
}
func toAddress() {
	entropy, _ := bip39.NewEntropy(160) // 15 words ||  word-->entropy
	mnemonic39, _ := bip39.NewMnemonic(entropy)
	fmt.Println("bip39 Mnemonic===> ", mnemonic39)
	//seed := bip39.NewSeed(mnemonic39, "")
	seed := bip39.NewSeed("coast tomorrow aerobic merit cause brain castle minute change impose column anger ketchup fantasy brush", "")
	/*拿到bip39生成的同一个 mnemonic */
	//bip32
	bip39analyse(seed)
	mnemonic39 = "coast tomorrow aerobic merit cause brain castle minute change impose column anger ketchup fantasy brush"
	//bip44analyse(mnemonic39)
	edunuzzi()
}
func bip39analyse(seed []byte) {
	masterKey, _ := bip32.NewMasterKey(seed)
	fmt.Println("masterKey ===>: ", masterKey, "/n masterKey.Depth: ", masterKey.Depth, " /nmasterKey.ChildNumber: ",
		masterKey.ChildNumber, " /n masterKey.Version: ", masterKey.Version, "/n masterKey.IsPrivate: ", masterKey.IsPrivate,
		"masterKey.ChainCode: ", masterKey.ChainCode)
	fmt.Println("masterKey.Key: ", masterKey.Key)
	fmt.Println("masterKey =================================>>>")
	publicKey := masterKey.PublicKey()
	fmt.Println("masterKey.PublicKey:==> ", publicKey)

	keysMap := map[string]*bip32.Key{}
	keysMap["root_0"], _ = masterKey.NewChildKey(0)
	fmt.Println("root_0key=>: ", keysMap["root_0"])
	fmt.Println("root_0publickey=>: ", keysMap["root_0"].PublicKey()) //right
	root_0_0key, _ := keysMap["root_0"].NewChildKey(0)
	fmt.Println("root_0_0key=>: ", root_0_0key)
	fmt.Println("root_0_0key.Depth=>: ", root_0_0key.Depth)
	fmt.Println("root_0_0key.ChildNumber=>: ", root_0_0key.ChildNumber)
	root_0_0keypubkey := root_0_0key.PublicKey()
	fmt.Println("root_0_0key=>: ", root_0_0key)
	fmt.Println("root_0_0keypubkey=>: ", root_0_0keypubkey)

	keysMap["root_1"], _ = masterKey.NewChildKey(1)
	departmentAuditKeys := map[string]*bip32.Key{}
	departmentAuditKeys["root_0pub"] = keysMap["root_0"].PublicKey()
	for department, pubKey := range departmentAuditKeys {
		fmt.Println(department+"===>", keysMap[department], "===>", pubKey)
		//fmt.Println("===>", keysMap[department].Depth, "===>", keysMap[department].ChildNumber)
	}
}
func bip44analyse(mnemonic39 string) {
	mnemonic, _ := bip44.NewMnemonic(160)
/*	mnemonic.Value = mnemonic39
	println("mnemonic39 is =====>", mnemonic.Value)
	seedBytes, _ := mnemonic.NewSeed("")
	println("seedBytes is =====>", seedBytes)*/

	seedBytes, _ := mnemonic.NewSeed("")
	//encodeStr := hex.EncodeToString(seedBytes)
	//xKey, err := bip44.NewKeyFromSeedHex(encodeStr, bip44.MAINNET)

	xKey, _ := bip44.NewKeyFromSeedBytes(seedBytes, bip44.MAINNET)
	println("xKey is =====>", xKey)
	accountKey, _ := xKey.BIP44AccountKey(bip44.BitcoinCoinType, 0, true)
	externalAddress, _ := accountKey.DeriveP2PKAddress(bip44.ExternalChangeType, 1, bip44.MAINNET)
	println("externalAddress is =====>", externalAddress)

}
func edunuzzi()  {
	bitSize := 256
	mnemonic, _ := bip44.NewMnemonic(bitSize)
	seedBytes,_ := mnemonic.NewSeed("")
	xKey, _ := bip44.NewKeyFromSeedBytes(seedBytes, bip44.MAINNET)
	accountKey, _ := xKey.BIP44AccountKey(bip44.BitcoinCoinType, 0, true)

	externalAddress, _ := accountKey.DeriveP2PKAddress(bip44.ExternalChangeType, 0, bip44.MAINNET)
	println("externalAddress is =====>", externalAddress)
}
