package main

import (
	"example/address"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/medeirosfalante/ethcli"
)

func main() {

	// Setup Address based on bip 39.
	account := address.GetAddress("1")

	// Initializiate the rpc connection with the network
	// In this example was used BSC Test net
	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
	if err != nil {
		log.Printf("err %s", err.Error())
		return
	}

	// Call a native client, for use token, look ours token section instead.
	native := ethcli.NewNative(client)

	balance, err := native.BalanceOf(account.FromAddres)

	if err != nil {
		log.Fatalln("An error ocurred due to %w", err)
	}

	// Setup the address which will recive the transaction
	account.ToAddress = "0x652F773E442264e0A9aC299264bD976eF8Fe6725"

	log.Printf("Before Txn BNB Balance: %v", balance)

	// Information needed to sign the transaction
	req := ethcli.TransferOpts{
		Mnemonic: account.Mnemonic,
		Path:     account.Path,
		Address:  account.ToAddress,
		Amount:   0.05,
	}

	res, err := native.Transfer(&req)

	if err != nil {
		log.Fatalln(err)
	}

	// The result must be the generated hash for transaction
	log.Println(res)
}
