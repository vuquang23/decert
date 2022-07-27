package repositories

import "gorm.io/gorm"

type collectionRepository struct {
	db *gorm.DB
}

var collectionRepo *collectionRepository

func InitCollectionRepository(db *gorm.DB) {
	if collectionRepo == nil {
		collectionRepo = &collectionRepository{
			db: db,
		}
	}
}

func CollectionRepositoryInstance() ICollectionRepository {
	return collectionRepo
}
