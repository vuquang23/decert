package config

type BlockchainConfig struct {
	BscBlockTimeSecond int64  `default:"3" env:"BSCBLOCKTIMESECOND"`
	BscTestnetRpc      string `default:"https://data-seed-prebsc-1-s1.binance.org:8545" env:"BSCTESTNETRPC"`
}
