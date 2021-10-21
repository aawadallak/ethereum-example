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

	// Call a new TOKEN, must pass token contract address
	token := ethcli.NewTokenErc20("0xf8d596bcce8dd6bc447bf7c16ad238b9643cac07", client)

	balance, err := token.BalanceOf(account.FromAddres)

	if err != nil {
		log.Fatalln("An error ocurred due to %w", err)
	}

	// Setup the address which will recive the transaction
	account.ToAddress = "0xA2959D3F95eAe5dC7D70144Ce1b73b403b7EB6E0"

	log.Printf("Before Txn Balance: %v", balance)

	// Information needed to sign the transaction
	req := ethcli.TransferOpts{
		Mnemonic: account.Mnemonic,
		Path:     account.Path,
		Address:  account.ToAddress,
		Amount:   10,
	}

	// Call the transfer function
	res, err := token.Transfer(&req)

	if err != nil {
		log.Fatalln(err)
	}

	// The result must be the generated hash for transaction
	log.Println(res)
}
