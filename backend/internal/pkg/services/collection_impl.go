package services

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"decert/internal/pkg/blockchain/decert"
	"decert/internal/pkg/config"
	"decert/internal/pkg/constants"
	"decert/internal/pkg/entity"
	"decert/internal/pkg/errors"
	e "decert/internal/pkg/errors"
	"decert/internal/pkg/repositories"
	"decert/internal/pkg/transformers"
	"decert/internal/pkg/utils/log"
	"decert/internal/pkg/utils/uuid"
)

type collectionService struct {
	collectionRepo repositories.ICollectionRepository
	ethClient      *ethclient.Client
}

var collectionSvc *collectionService

func InitCollectionService(
	collectionRepo repositories.ICollectionRepository,
	ethClient *ethclient.Client,
) {
	if collectionSvc == nil {
		collectionSvc = &collectionService{
			collectionRepo: collectionRepo,
			ethClient:      ethClient,
		}
	}
}

func CollectionServiceInstance() ICollectionService {
	return collectionSvc
}

func (s *collectionService) CreateCollection(ctx *gin.Context, crudCreateCollection entity.CRUDCreateCollection) *e.DomainError {
	go s.searchCollectionByTxHash(
		new(gin.Context),
		crudCreateCollection,
	)
	return nil
}

func (s *collectionService) searchCollectionByTxHash(ctx *gin.Context, crudCreateCollection entity.CRUDCreateCollection) {
	ctx.Set(constants.Prefix, uuid.MsgWithUUID("search-collection-by-txhash-"+crudCreateCollection.TxHash))

	blockTime := config.Instance().Blockchain.BscBlockTimeSecond
	txHash := common.HexToHash(crudCreateCollection.TxHash)
	for i := 1; i <= 15; i++ {
		time.Sleep(time.Second * time.Duration(blockTime))
		tx, isPending, err := s.ethClient.TransactionByHash(context.Background(), txHash)

		if err != nil {
			log.Errorln(ctx, err)
			continue
		}
		if isPending {
			log.Debugln(ctx, crudCreateCollection.TxHash+" is pending")
			continue
		}

		receipt, err := s.ethClient.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Errorln(ctx, err)
			continue
		}

		address := receipt.Logs[0].Topics[1].String()
		startIndex := len(address) - 40
		endIndex := len(address)
		ethAddress := "0x" + address[startIndex:endIndex]
		log.Debugf(ctx, "New collection address: %s\n", ethAddress)

		if err := s.saveNewCreatedCollection(ctx, ethAddress); err != nil {
			log.Errorln(ctx, err)
			return
		}
		break
	}
}

func (s *collectionService) saveNewCreatedCollection(ctx *gin.Context, ethAddress string) *e.DomainError {
	address := common.HexToAddress(ethAddress)
	instance, err := decert.NewDecert(address, s.ethClient)
	if err != nil {
		return e.NewDomainErrorUnknown([]string{}, err)
	}

	title, _ := instance.Name(nil)
	symbol, _ := instance.Symbol(nil)
	issuer, _ := instance.Issuer(nil)

	if err := s.collectionRepo.SaveNewCollection(ctx, ethAddress, title, symbol, issuer.String()); err != nil {
		return transformers.DomainErrTransformerInstance().Transform(err)
	}

	return nil
}

func (s *collectionService) GetCollections(
	ctx *gin.Context,
	crudGetCollections entity.CRUDGetCollections,
) ([]*entity.Collection, *errors.DomainError) {
	ret, err := s.collectionRepo.GetCollectionsByIssuerAddress(ctx, crudGetCollections.Issuer,
		crudGetCollections.Limit, crudGetCollections.Offset)
	if err != nil {
		return nil, transformers.DomainErrTransformerInstance().Transform(err)
	}
	return ret, nil
}
