package address

import (
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type Account struct {
	Mnemonic   string
	FromAddres string
	ToAddress  string
	Path       string
	Amount     float64
}

const (
	mnemonic = "earn ripple small pyramid century crash poet print tired cat mix audit ticket sport twelve"
)

func GetAddress(path string) *Account {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	derivationPath := hdwallet.MustParseDerivationPath(path)
	account, err := wallet.Derive(derivationPath, true)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return &Account{
		Mnemonic:   mnemonic,
		FromAddres: account.Address.Hex(),
		Path:       "1",
		ToAddress:  "",
		Amount:     0,
	}
}
