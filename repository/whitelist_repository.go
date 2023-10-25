package repository

import (
	"github.com/Axrous/mnc/helper"
	"github.com/Axrous/mnc/model/domain"
	"github.com/sonyarouje/simdb"
)

type WhitelistRepository interface {
	Save(whitelist domain.Whitelist)
	FindById(whitelist domain.Whitelist) (domain.Whitelist, error)
	Delete(whitelist domain.Whitelist) error
}

type whitelistRepositoryImpl struct {
	db *simdb.Driver
}

// Delete implements WhitelistRepository.
func (repo *whitelistRepositoryImpl) Delete(whitelist domain.Whitelist) error {
	return repo.db.Delete(&whitelist)

}

// FindById implements WhitelistRepository.
func (repo *whitelistRepositoryImpl) FindById(whitelist domain.Whitelist) (domain.Whitelist, error) {
	var result domain.Whitelist
	err := repo.db.Open(&domain.Whitelist{}).Where("id", "=", whitelist.Id).First().AsEntity(&result)
	if err != nil {
		return domain.Whitelist{}, err
	}
	return result, nil
}

// Save implements WhitelistRepository.
func (repo *whitelistRepositoryImpl) Save(whitelist domain.Whitelist) {
	err := repo.db.Insert(&whitelist)
	helper.PanicIfError(err)
}

func NewWhiteListRepository(db *simdb.Driver) WhitelistRepository {
	return &whitelistRepositoryImpl{
		db: db,
	}
}
