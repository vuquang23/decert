package dto

import (
	"time"
	"strings"
	"math/big"
)

type Datetime struct {
    time.Time
}

func (t *Datetime) UnmarshalJSON(input []byte) error {
    strInput := strings.Trim(string(input), `"`)
    newTime, err := time.Parse("2006-01-02T15:04:05-0700", strInput)
    if err != nil {
        return err
    }

    t.Time = newTime
    return nil
}

type CertDataTypeReceiver struct {
	Name		string
	Wallet		string
	DateOfBirth	string
}

type CertDataTypeIssuer struct {
	Name		string
	Wallet		string
	Position	string
}

type CertDataTypeRequest struct {
	CertTitle	string	`json:"certTitle"`
	Issuer		CertDataTypeIssuer	`json:"issuer"`
	Receiver	CertDataTypeReceiver	`json:"receiver"`
	Description	string	`json:"description"`
	IssuedAt	int64	`json:"issuedAt"`
	ExpiredAt	int64	`json:"expiredAt"`
	CertImage	string		`json:"certImage"`
	Platform	string		`json:"platform"`
}

type CreateCertificateRequest struct {
	CollectionId	uint `json:"collectionId"`
	TxHash   	string `json:"txHash"`
	Platform 	string `json:"platform"`
	CertData	CertDataTypeRequest `json:"certData"`
}

type GetCertificatesRequest struct {
	CollectionId	uint `json:"collectionId"`
	Offset	uint `json:"offset"`
	Limit	uint `json:"limit"`
}

type GetCertificateRequest struct {
	CertId	uint `json:"certId"`
}

type CertificateResponse struct {
	ID              uint					`json:"id"`
	Description     string					`json:"description"`
	IssuedAt       	big.Int					`json:"issuedAt"`
	ExpiredAt       big.Int					`json:"expiredAt"`
	CollectionId   	uint					`json:"collectionId"`
	CertNftId       uint					`json:"certNftId"`
	Data  			string					`json:"data"`
	RevokedAt 		big.Int					`json:"revokedAt"`
	RevokedReason	string					`json:"revokedReason"`
	Receiver		CertDataTypeReceiver	`json:"receiver"`
}
