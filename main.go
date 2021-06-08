package main

import (
	"fmt"
	"github.com/jamesfyp/spvwallet-demo/coinbase"
	"github.com/jamesfyp/spvwallet-demo/wallet"
	 "github.com/ethereum/go-ethereum/accounts"
	"log"
)

func main() {

}

func NewAccount() {
	mnemonic := wallet.NewMnemonic()
	if mnemonic == "" {
		panic("no mnemonic")
	}
	seed, err := wallet.NewSeedFromMnemonic(mnemonic, "")
	if err != nil {
		log.Fatal(err)
	}

	account, err := wallet.NewWallet(seed)
	if err != nil {
		log.Fatal(err)
	}

	path, err := accounts.ParseDerivationPath(fmt.Sprintf("m/44'/%d'/0'/0/0", coinbase.DOGE))


	pubKeyECDSA, err :=



}