package transaction

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func CreateRawTransaction(client *ethclient.Client, txMessage string, privatekeyHex string) []byte {
	// get caller's private key
	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}
	// get caller's address from public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// get pending nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}
	// set up transaction configuration
	value := big.NewInt(0)
	gasLimit := uint64(5000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}
	// add sql/nosql command to new transaction
	data := []byte(txMessage)
	tx := types.NewTransaction(nonce, fromAddress, value, gasLimit, gasPrice, data)
	// get chain id
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}
	// sign the newly generated transaction with private key
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}
	// get hex form of signed transaction
	rawTxBytes, _ := rlp.EncodeToBytes(signedTx)
	return rawTxBytes
}
