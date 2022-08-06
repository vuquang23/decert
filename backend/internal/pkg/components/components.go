package components

import (
	"decert/internal/pkg/blockchain"
	"decert/internal/pkg/db/mysql"
	"decert/internal/pkg/repositories"
	"decert/internal/pkg/services"
)

func Init() error {
	// infra
	if err := mysql.Init(); err != nil {
		return err
	}

	if err := blockchain.InitEthClient(); err != nil {
		return err
	}

	// repo
	repositories.InitCollectionRepository(mysql.Instance())
	repositories.InitCertificateRepository(mysql.Instance())

	// service
	services.InitCollectionService(
		repositories.CollectionRepositoryInstance(),
		blockchain.BlockchainClientInstance(),
	)
	services.InitCertificateService(
		repositories.CertificateRepositoryInstance(),
		blockchain.BlockchainClientInstance(),
	)

	return nil
}
