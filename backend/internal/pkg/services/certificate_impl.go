package services

import (
	"context"
	"fmt"
	"time"
	"math/big"
	"strconv"

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
	go s.searchCertificateByTxHash(
		new(gin.Context),
		crudCreateCertificate,
	)
	return nil
}

func (s *certificateService) searchCertificateByTxHash(ctx *gin.Context, crudCreateCertificate entity.CRUDCreateCertificate) {
	fmt.Println("Searching cert by txhash")
	ctx.Set(constants.Prefix, uuid.MsgWithUUID("search-certificate-by-txhash-"+crudCreateCertificate.TxHash))

	blockTime := config.Instance().Blockchain.BscBlockTimeSecond
	txHash := common.HexToHash(crudCreateCertificate.TxHash)
	for i := 1; i <= 15; i++ {
		time.Sleep(time.Second * time.Duration(blockTime))
		log.Debugln(ctx, txHash, blockTime)

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

		nftId := receipt.Logs[0].Topics[3].Big()
		if err != nil {
			log.Errorln(ctx, err)
			continue
		}
		log.Debugf(ctx, "New certificate id: %s\n", nftId)

		if err := s.saveNewCreatedCertificate(ctx, nftId, crudCreateCertificate); err != nil {
			log.Errorln(ctx, err)
			return
		}
		break
	}
}

func (s *certificateService) saveNewCreatedCertificate(ctx *gin.Context, nftId *big.Int, crudCreateCertificate entity.CRUDCreateCertificate) *e.DomainError {
	decertCaller, err := decert.NewDecertCaller(common.HexToAddress("0x81b959030Bf959e6397b4a3f9Ab13AF51915f27d"), s.ethClient)
	if err != nil {
		return e.NewDomainErrorUnknown([]string{}, err)
	}
	fmt.Println("decertCaller", decertCaller)
	certData, err := decertCaller.CertData(nil, nftId)
	if err != nil {
		return e.NewDomainErrorUnknown([]string{}, err)
	}
	fmt.Println("Certdata", certData)
	
	issuer := certData.Issuer
	recipient := certData.Recipient
	certHash := certData.CertHash
	link := certData.Link
	issuedAt := certData.IssuedAt
	fmt.Println("Cert data from blockchain", issuer, recipient, certHash, link, issuedAt)

	nftIdInt64, err := strconv.ParseInt(nftId.String(), 10, 64)


	if err := s.certificateRepo.SaveNewCertificate(
		ctx, 
		crudCreateCertificate,
		uint(nftIdInt64),
	); err != nil {
		return transformers.DomainErrTransformerInstance().Transform(err)
	}

	return nil
}

func (s *certificateService) GetCertificates(
	ctx *gin.Context,
	crudGetCertificates entity.CRUDGetCertificates,
) ([]*entity.Cert, *errors.DomainError) {
	ret, err := s.certificateRepo.GetCertificatesByCollectionId(
		ctx, 
		crudGetCertificates.CollectionId,
		crudGetCertificates.Limit,
		crudGetCertificates.Offset,
	)
	if err != nil {
		return nil, transformers.DomainErrTransformerInstance().Transform(err)
	}
	return ret, nil
}

func (s *certificateService) GetCertificateInfo(
	ctx *gin.Context, 
	crudGetCertificate entity.CRUDGetCertificate,
) (*entity.Cert, *errors.DomainError) {
	ret, err := s.certificateRepo.GetCertificateById(ctx, crudGetCertificate.ID)
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