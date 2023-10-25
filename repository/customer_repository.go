package repository

import (
	"errors"

	"github.com/Axrous/mnc/model/domain"
	"github.com/sonyarouje/simdb"
)

type CustomerRepository interface {
	Save(customer domain.Customer)
	FindById(customer domain.Customer) (domain.Customer, error)
	FindByUsername(customer domain.Customer) (domain.Customer, error)
}

type customerRepositoryImpl struct {
	db *simdb.Driver
}

// FindByUsername implements CustomerRepository.
func (repo *customerRepositoryImpl) FindByUsername(customer domain.Customer) (domain.Customer, error) {
	var result domain.Customer
	err := repo.db.Open(&domain.Customer{}).Where("username", "=", customer.Username).First().AsEntity(&result)
	if err != nil {
		return domain.Customer{}, errors.New("customer not found")
	}
	return result, nil
}

// FindById implements CustomerRepository.
func (repo *customerRepositoryImpl) FindById(customer domain.Customer) (domain.Customer, error) {
	var result domain.Customer
	err := repo.db.Open(&domain.Customer{}).Where("id", "=", customer.Id).First().AsEntity(&result)
	if err != nil {
		return domain.Customer{}, errors.New("customer not found")
	}
	return result, nil
}

// Save implements CustomerRepository.
func (repo *customerRepositoryImpl) Save(customer domain.Customer) {
	repo.db.Insert(&customer)
}

func NewCustomerRepository(db *simdb.Driver) CustomerRepository {
	return &customerRepositoryImpl{
		db: db,
	}
}
