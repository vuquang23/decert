package entity

import (
	"time"
)

type CertDataTypeIssuer struct {
	Name		string
	Wallet		string
	Position	string
}

type CertDataTypeReceiver struct {
	Name		string
	Wallet		string
	DateOfBirth	string
}

type CertDataType struct {
	CertTitle	string
	Issuer		CertDataTypeIssuer
	Receiver	CertDataTypeReceiver
	Description	string
	IssuedAt	int64
	ExpiredAt	int64
	CertImage	string
	Platform	string
}

type CRUDCreateCertificate struct {
	TxHash   	string
	Platform 	string
	CollectionId	uint
	CertData	CertDataType
}

type CRUDGetCertificates struct {
	CollectionId 	uint
	Limit			uint
	Offset			uint
}

type CRUDGetCertificate struct {
	ID	uint
}

type CRUDRevokeCertificate struct {
	ID	uint
	TxHash   	string
	Platform 	string
	RevokedAt 		int64
	RevokedReason	string
}

type Cert struct {
	ID           	uint `gorm:"primaryKey"`
	Description     string
	IssuedAt       	time.Time
	ExpiredAt       time.Time
	CollectionId   	uint
	CertNftId       uint
	Data  			string
	RevokedAt 		time.Time
	RevokedReason	string
	Receiver		string
}
