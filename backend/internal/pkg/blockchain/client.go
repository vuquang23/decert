package blockchain

import (
	"decert/internal/pkg/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

var ethClient *ethclient.Client

func InitEthClient() error {
	if ethClient == nil {
		var err error
		ethClient, err = ethclient.Dial(config.Instance().Blockchain.BscTestnetRpc)
		if err != nil {
			return err
		}
	}
	return nil
}

func BlockchainClientInstance() *ethclient.Client {
	return ethClient
}
