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

type certificateService struct {
	certificateRepo repositories.ICertificateRepository
	ethClient      *ethclient.Client
}

var certificateSvc *certificateService

func InitCertificateService(
	certificateRepo repositories.ICertificateRepository,
	ethClient *ethclient.Client,
) {
	if certificateSvc == nil {
		certificateSvc = &certificateService{
			certificateRepo: certificateRepo,
			ethClient:      ethClient,
		}
	}
}

func CertificateServiceInstance() ICertificateService {
	return certificateSvc
}

func (s *certificateService) CreateCertificate(ctx *gin.Context, crudCreateCertificate entity.CRUDCreateCertificate) *e.DomainError {
	// TODO
	// go s.searchCertificateByTxHash(
	// 	new(gin.Context),
	// 	crudCreateCertificate,
	// )
	return nil
}

func (s *certificateService) searchCertificateByTxHash(ctx *gin.Context, crudCreateCertificate entity.CRUDCreateCertificate) {
	ctx.Set(constants.Prefix, uuid.MsgWithUUID("search-certificate-by-txhash-"+crudCreateCertificate.TxHash))

	blockTime := config.Instance().Blockchain.BscBlockTimeSecond
	txHash := common.HexToHash(crudCreateCertificate.TxHash)
	for i := 1; i <= 15; i++ {
		time.Sleep(time.Second * time.Duration(blockTime))
		tx, isPending, err := s.ethClient.TransactionByHash(context.Background(), txHash)

		if err != nil {
			log.Errorln(ctx, err)
			continue
		}
		if isPending {
			log.Debugln(ctx, crudCreateCertificate.TxHash+" is pending")
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
		log.Debugf(ctx, "New certificate address: %s\n", ethAddress)

		if err := s.saveNewCreatedCertificate(ctx, ethAddress); err != nil {
			log.Errorln(ctx, err)
			return
		}
		break
	}
}

func (s *certificateService) saveNewCreatedCertificate(ctx *gin.Context, ethAddress string) *e.DomainError {
	address := common.HexToAddress(ethAddress)
	instance, err := decert.NewDecert(address, s.ethClient)
	if err != nil {
		return e.NewDomainErrorUnknown([]string{}, err)
	}

	title, _ := instance.Name(nil)
	symbol, _ := instance.Symbol(nil)
	issuer, _ := instance.Issuer(nil)

	if err := s.certificateRepo.SaveNewCertificate(ctx, ethAddress, title, symbol, issuer.String()); err != nil {
		return transformers.DomainErrTransformerInstance().Transform(err)
	}

	return nil
}

func (s *certificateService) GetCertificates(
	ctx *gin.Context,
	crudGetCertificates entity.CRUDGetCertificates,
) ([]*entity.Certificate, *errors.DomainError) {
	ret, err := s.certificateRepo.GetCertificatesByCollectionId(ctx, crudGetCertificates.CollectionId)
	if err != nil {
		return nil, transformers.DomainErrTransformerInstance().Transform(err)
	}
	return ret, nil
}

func (s *certificateService) GetCertificateInfo(
	ctx *gin.Context, 
	crudGetCertificateInfo entity.CRUDGetCertificateInfo,
) (*entity.Certificate, *errors.DomainError) {
	ret, err := s.certificateRepo.GetCertificateById(ctx, crudGetCertificateInfo.ID)
	if err != nil {
		return nil, transformers.DomainErrTransformerInstance().Transform(err)
	}
	return ret, nil
}

func (s *certificateService) RevokeCertificate(
	ctx *gin.Context, 
	crudRevokeCertificate entity.CRUDRevokeCertificate,
) *errors.DomainError {
	// TODO 
	return nil
}