package dto

import (
	"time"
	"strings"
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

type CertDataTypeRequest struct {
	CertTitle	string	`json:"certTitle"`
	Issuer		string	`json:"issuer"`
	Receiver	string	`json:"receiver"`
	Description	string	`json:"description"`
	IssuedAt	Datetime	`json:"issuedAt"`
	ExpiredAt	Datetime	`json:"expiredAt"`
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
	// TODO
	// Issuer string `form:"issuer,required"`
	// Limit  uint `form:"limit,required"`
	// Offset uint `form:"offset"`
}

type CertificateResponse struct {
	ID                uint   `json:"id"`
	// CertificateName    string `json:"CertificateName"`
	// CertificateSymbol  string `json:"CertificateSymbol"`
	// CertificateAddress string `json:"CertificateAddress"`
	// Issuer            string `json:"issuer"`
	// TotalIssued       uint `json:"totalIssued"`
	// TotalRevoked      uint `json:"totalRevoked"`
}
