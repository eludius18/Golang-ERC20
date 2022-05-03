package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	contract "unlocked/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ganacheUrl     = "http://localhost:7545"
	addr2          = ""
	rinkebyAddress = ""
	infuraUrl      = ""
	pKey           = ""
)

func main() {

	// Private Key
	privateKey, err := crypto.HexToECDSA(pKey)
	if err != nil {
		log.Fatal(err)
	}

	// Client
	client, err := ethclient.Dial(infuraUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	add := common.HexToAddress(rinkebyAddress)

	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(8000000)
	auth.Nonce = big.NewInt(int64(nonce))

	a, tx, _, err := contract.DeployUnlocked(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------------------------")
	fmt.Println("CONTRACT ADDRESS:", a.Hex())
	fmt.Println("TRANSACTION HASH:", tx.Hash().Hex())
	fmt.Println("-----------------------------------")
}