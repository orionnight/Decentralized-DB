package services

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func addressCheck() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := common.HexToAddress("0xf7157e022db032C63A0d557D46c205BC00cEDdfc")
	fmt.Printf("gas Price: %s\n", fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0)      // in wei (0.01 eth)
	gasLimit := uint64(5000000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x5Ed61D6D96432Cfa5F973ffFdC0cBb5A7247FD58")
	data := []byte("hello ")
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("gas Price: %d\n", gasPrice)
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	signedTx, err := ks.Wallets()[0].SignTxWithPassphrase(ks.Accounts()[0], "hongxianggeng", tx, chainID)
	if err != nil {
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
	rawTxBytes := new(bytes.Buffer)
	// ts.EncodeIndex(0, b)
	// rawTxBytes, err := hex.DecodeString(rawTx)
	ts.EncodeIndex(0, rawTxBytes)

	tx = new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes.Bytes(), &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
