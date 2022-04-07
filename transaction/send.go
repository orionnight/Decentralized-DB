package transaction

import (
	"context"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func SendRawTransaction(txMessage string, privatekeyHex string) string {
	// op = "CREATE" or "DELETE" or "UPDATE"
	quorumEndpoint := "http://localhost:8545"

	// connect to testnet
	client, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-sender").Error(err.Error())
	}
	// create raw transaction
	rawTxBytes := CreateRawTransaction(client, txMessage, privatekeyHex)

	tx := new(types.Transaction)
	// fill tx with rawTxBytes
	rlp.DecodeBytes(rawTxBytes, &tx)
	// broadcast the transaction
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-sender").Error(err.Error())
	}

	return tx.Hash().Hex()
}
