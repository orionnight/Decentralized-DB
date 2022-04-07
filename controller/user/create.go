package user

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"example.com/ece1770/logger"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type BlockChainUser struct {
	privateKeyHex string
	accountHex    string
}

var (
	bcuser *BlockChainUser
)

func CreateKeystore(c *gin.Context) {
	// new user
}

func Login(c *gin.Context) {
	jsonObj := getJson(c)

	if privateKeyHex, ok := jsonObj["privateKeyHex"].(string); ok {
		privateKey, err := crypto.HexToECDSA(privateKeyHex)
		if err != nil {
			logger.InternalLogger.WithField("component", "user-login").Error(err.Error())
		}
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			logger.InternalLogger.WithField("component", "user-login").Error(err.Error())
		}
		account := crypto.PubkeyToAddress(*publicKeyECDSA)

		quorumEndpoint := "http://localhost:8545"

		// connect to testnet
		client, err := ethclient.Dial(quorumEndpoint)
		if err != nil {
			logger.InternalLogger.WithField("component", "user-login").Error(err.Error())
		}
		balance, err := client.BalanceAt(context.Background(), account, nil)
		if err != nil {
			logger.InternalLogger.WithField("component", "user-login").Error(err.Error())
		}
		if balance.Cmp(big.NewInt(0)) <= 0 {
			logger.InternalLogger.WithField("component", "user-login").Error(balance)
		}
		bcuser = &BlockChainUser{privateKeyHex: privateKeyHex, accountHex: account.Hex()}
		fmt.Println("User login successful!")
		go transactionListener()
	} else {
		fmt.Println("privateKeyHex not provided!")
	}

	// start
}

func transactionListener() {
	client, err := ethclient.Dial("ws://localhost:8545")
	println("Inside TransactionListener()...")
	println("logged in account is", GetAccountHex())
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			// fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			// fmt.Println(block.Number().Uint64())   // 3477413
			// fmt.Println(block.Time())              // 1529525947
			// fmt.Println(block.Nonce())             // 130524141876765836
			// fmt.Println(len(block.Transactions())) // 7

			for _, tx := range block.Transactions() {
				fmt.Println(string(tx.Data()))
				// if tx.To().Hex() != GetAccountHex() {
				// 	// Based on data do sync
				// 	fmt.Println(string(tx.Data()))
				// }
			}
		}
	}
}

func GetPrivateKeyHex() string {
	if bcuser != nil {
		return bcuser.privateKeyHex
	}
	return ""
}

func GetAccountHex() string {
	if bcuser != nil {
		return bcuser.accountHex
	}
	return ""
}
