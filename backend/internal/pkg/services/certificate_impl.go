package services

import (
	"context"
	"time"
	"math/big"
	"strconv"
	"fmt"
	"strings"
	
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
	// collectionRepository := repositories.CollectionRepositoryInstance()
	// collection, _ := collectionRepository.GetCollectionById(ctx, crudCreateCertificate.CollectionId)

	// decertCaller, err := decert.NewDecertCaller(common.HexToAddress(collection.Issuer), s.ethClient)
	// decertTransactor, err := decert.NewDecertTransactor(common.HexToAddress(collection.Issuer), s.ethClient)

	// currentBatchSize, _ := decertCaller.BatchSize(bind.CallOpts{})
	// fmt.Println("Batch size", currentBatchSize)

	// if err != nil {
	// 	return e.NewDomainErrorUnknown([]string{}, err)
	// }
	// fmt.Println("decertCaller", decertCaller)
	// fmt.Println("decertTransactor", decertTransactor)
	// certData, err := decertCaller.CertData(nil, nftId)
	// if err != nil {
	// 	return e.NewDomainErrorUnknown([]string{}, err)
	// }
	// fmt.Println("Certdata", certData)
	
	// issuer := certData.Issuer
	// recipient := certData.Recipient
	// certHash := certData.CertHash
	// link := certData.Link
	// issuedAt := certData.IssuedAt
	// fmt.Println("Cert data from blockchain", issuer, recipient, certHash, link, issuedAt)

	nftIdInt64, _ := strconv.ParseInt(nftId.String(), 10, 64)


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
) (*errors.DomainError) {
	go s.RevokeCertificateById(
		new(gin.Context),
		crudRevokeCertificate,
	)
	return nil
}


func DecodeTransactionInputData(
	ctx *gin.Context,
	contractABI *abi.ABI, 
	data []byte,
	crudRevokeCertificate *entity.CRUDRevokeCertificate,
) {
	methodSigData := data[:4]
	inputsSigData := data[4:]
	method, err := contractABI.MethodById(methodSigData)
	if err != nil {
		log.Errorln(ctx, err)
		return
	}
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		log.Errorln(ctx, err)
	} else {
		fmt.Println(inputsMap)
	}
	fmt.Printf("Method Name: %s\n", method.Name)
	fmt.Printf("Method inputs: %v\n", inputsMap["_reason"])
	crudRevokeCertificate.RevokedAt = inputsMap["revokedAt"].(*big.Int).Int64()
	crudRevokeCertificate.RevokedReason = inputsMap["_reason"].(string)

	return
}

func (s *certificateService) RevokeCertificateById(
	ctx *gin.Context,
	crudRevokeCertificate entity.CRUDRevokeCertificate,
) {
	ctx.Set(constants.Prefix, uuid.MsgWithUUID("revoke certificate txhash="+crudRevokeCertificate.TxHash))

	dbCertificate, err := s.certificateRepo.GetCertificateById(ctx, crudRevokeCertificate.ID)
	if (err != nil || dbCertificate == nil) {
		log.Errorln(ctx, dbCertificate, err)
		return
	}
		
	collectionRepository := repositories.CollectionRepositoryInstance()
	dbCollection, err := collectionRepository.GetCollectionById(ctx, dbCertificate.CollectionId)
	if (err != nil || dbCollection == nil) {
		log.Errorln(ctx, err)
		return
	}

	blockTime := config.Instance().Blockchain.BscBlockTimeSecond
	txHash := common.HexToHash(crudRevokeCertificate.TxHash)
	for i := 1; i <= 15; i++ {
		time.Sleep(time.Second * time.Duration(blockTime))
		log.Debugln(ctx, txHash, blockTime)

		tx, isPending, err := s.ethClient.TransactionByHash(context.Background(), txHash)

		if err != nil {
			log.Errorln(ctx, err)
			continue
		}
		if isPending {
			log.Debugln(ctx, crudRevokeCertificate.TxHash+" is pending")
			continue
		}

		receipt, err := s.ethClient.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Errorln(ctx, err)
			continue
		}

		bcNftId := receipt.Logs[0].Topics[3].Big()
		bcCollectionAddress := common.HexToAddress(receipt.Logs[0].Address.String()).String()

		dbNftId := big.NewInt(int64(dbCertificate.CertNftId))
		dbCollectionAddress := common.HexToAddress(dbCollection.Address).String()

		inputData := tx.Data()
		tx.UnmarshalJSON(inputData)
		fmt.Println("DATA", (inputData))
		fmt.Println("DATA", string(inputData))
		
		var decertABI abi.ABI
		decertABI, err = abi.JSON(strings.NewReader(decert.DecertABI))
		if err != nil {
			log.Errorln(ctx, err)
			return
		}
		fmt.Println(decertABI)
		DecodeTransactionInputData(ctx, &decertABI, inputData, &crudRevokeCertificate)

		if (dbNftId.Cmp(bcNftId) != 0 || dbCollectionAddress != (bcCollectionAddress)) {
			log.Errorln(ctx, "collection or object id mismatch")
			fmt.Println("nftId db,bc", dbNftId, bcNftId, dbNftId.Cmp(bcNftId) != 0)
			fmt.Println("collectionAddress db,bc", dbCollectionAddress, bcCollectionAddress, dbCollectionAddress != bcCollectionAddress)
			fmt.Println(len(dbCollectionAddress), len(bcCollectionAddress))
			return
		}
		
		if err := s.certificateRepo.RevokeCertificate(ctx, crudRevokeCertificate); err != nil {
			log.Errorln(ctx, err)
			return
		}
		break
	}
}