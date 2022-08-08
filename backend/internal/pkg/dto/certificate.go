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

type CertDataReceiver struct {
	Name		string		`json:"name"`
	Wallet		string		`json:"wallet"`
	DateOfBirth	string		`json:"dateOfBirth"`
}

type CertDataIssuer struct {
	Name		string		`json:"name"`
	Wallet		string		`json:"wallet"`
	Position	string		`json:"position"`
}

type CertDataType struct {
	CertTitle	string				`json:"certTitle"`
	Issuer		CertDataIssuer		`json:"issuer"`
	Receiver	CertDataReceiver	`json:"receiver"`
	Description	string				`json:"description"`
	IssuedAt	string				`json:"issuedAt"`
	ExpiredAt	string				`json:"expiredAt"`
	CertImage	string				`json:"certImage"`
	Platform	string				`json:"platform"`
}

type CertDataTypeFromDB struct {
	CertTitle	string				`json:"certTitle"`
	Issuer		CertDataIssuer		`json:"issuer"`
	Receiver	CertDataReceiver	`json:"receiver"`
	Description	string				`json:"description"`
	IssuedAt	int64				`json:"issuedAt"`
	ExpiredAt	int64				`json:"expiredAt"`
	CertImage	string				`json:"certImage"`
	Platform	string				`json:"platform"`
}

type CreateCertificateRequest struct {
	CollectionId	uint 			`json:"collectionId"`
	TxHash   		string 			`json:"txHash"`
	Platform 		string 			`json:"platform"`
	CertData		CertDataType 	`json:"certData"`
}

type GetCertificatesRequest struct {
	CollectionId	uint 	`json:"collectionId"`
	ReceiverAddress	string	`json:"receiverAddress"`
	Offset			uint 	`json:"offset"`
	Limit			uint 	`json:"limit"`
}

type GetCertificateRequest struct {
	CertId	uint `json:"certId"`
}

type RevokeCertificateRequest struct {
	CertId	uint `json:"certId"`
	TxHash   	string `json:"txHash"`
	Platform 	string `json:"platform"`
}

type CertificateResponse struct {
	ID              uint				`json:"id"`
	Description     string				`json:"description"`
	IssuedAt       	string				`json:"issuedAt"`
	ExpiredAt       string				`json:"expiredAt"`
	CollectionId   	uint				`json:"collectionId"`
	CertNftId       uint				`json:"certNftId"`
	Data  			CertDataType		`json:"certData"`
	RevokedAt 		string				`json:"revokedAt"`
	RevokedReason	string				`json:"revokedReason"`
	Receiver		CertDataReceiver	`json:"receiver"`
	ReceiverAddress string				`json:"receiverAddress"`
}
