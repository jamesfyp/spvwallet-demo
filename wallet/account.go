package wallet

import (
	"errors"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tyler-smith/go-bip39"
	"os"
	"sync"
)

// Wallet is the underlying wallet struct.
type Wallet struct {
	mnemonic    string
	masterKey   *hdkeychain.ExtendedKey
	seed        []byte
	url         accounts.URL
	paths       map[common.Address]accounts.DerivationPath
	accounts    []accounts.Account
	stateLock   sync.RWMutex
	fixIssue172 bool
}

const issue179FixEnvar = "GO_ETHEREUM_HDWALLET_FIX_ISSUE_179"

type Account struct {
	MasterKey *hdkeychain.ExtendedKey
	Mnemonic  string
}

// NewMnemonic 助记词
func NewMnemonic() string {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return mnemonic
}

// NewSeedFromMnemonic 种子
func NewSeedFromMnemonic(mnemonic string, password string) ([]byte, error) {
	if mnemonic == "" {
		return nil, errors.New("mnemonic is required")
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("mnemonic is invalid")
	}

	return bip39.NewSeedWithErrorChecking(mnemonic, password)

}

func NewWallet(seed []byte) (*Wallet, error) {
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		masterKey:   masterKey,
		seed:        seed,
		accounts:    []accounts.Account{},
		paths:       map[common.Address]accounts.DerivationPath{},
		fixIssue172: false || len(os.Getenv(issue179FixEnvar)) > 0,
	}, nil
}
