package entity

import (
	"time"
)

type CertDataType struct {
	CertTitle	string
	Issuer		string
	Receiver	string
	Description	string
	IssuedAt	time.Time
	ExpiredAt	time.Time
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
	CollectionId uint
}

type CRUDGetCertificateInfo struct {
	ID	uint
}

type CRUDRevokeCertificate struct {
	ID	uint
}

type Certificate struct {
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
