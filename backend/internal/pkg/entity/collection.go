package entity

import "time"

type CRUDCreateCollection struct {
	TxHash   string
	Platform string
}

type CRUDGetCollections struct {
	Issuer string
	Limit  uint64
	Offset uint64
	Name   string
}

type Collection struct {
	ID           uint `gorm:"primaryKey"`
	Title        string
	Symbol       string
	Address      string
	Issuer       string
	TotalIssued  uint64
	TotalRevoked uint64
	CreatedAt    time.Time
}
