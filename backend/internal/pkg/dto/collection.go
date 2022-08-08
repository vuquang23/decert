package dto

type CreateCollectionRequest struct {
	TxHash   string `json:"txHash"`
	Platform string `json:"platform"`
}

type GetCollectionRequest struct {
	ID		uint   `json:"id"`
}

type GetCollectionsRequest struct {
	Issuer string `form:"issuer,required"`
	Limit  uint64 `form:"limit,required"`
	Offset uint64 `form:"offset"`
	Name   string `form:"name"`
}

type CollectionResponse struct {
	ID                uint   `json:"id"`
	CollectionName    string `json:"collectionName"`
	CollectionSymbol  string `json:"collectionSymbol"`
	CollectionAddress string `json:"collectionAddress"`
	Issuer            string `json:"issuer"`
	TotalIssued       uint64 `json:"totalIssued"`
	TotalRevoked      uint64 `json:"totalRevoked"`
}
