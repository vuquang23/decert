package services

import "decert/internal/pkg/repositories"

type collectionService struct {
	collectionRepo repositories.ICollectionRepository
}

var collectionSvc *collectionService

func InitCollectionService(collectionRepo repositories.ICollectionRepository) {
	if collectionSvc == nil {
		collectionSvc = &collectionService{
			collectionRepo: collectionRepo,
		}
	}
}

func CollectionServiceInstance() ICollectionService {
	return collectionSvc
}
