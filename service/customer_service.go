package service

import (
	"context"

	"github.com/Axrous/mnc/exception"
	"github.com/Axrous/mnc/helper"
	"github.com/Axrous/mnc/model/domain"
	"github.com/Axrous/mnc/model/web"
	"github.com/Axrous/mnc/repository"
)

type CustomerService interface {
	Save(ctx context.Context, request web.CustomerCreateRequest)
	FindById(ctx context.Context, id string) web.CustomerResponse
	Login(ctx context.Context, request web.CustomerLoginRequest) string
	Logout(ctx context.Context)
}

type customerServiceImpl struct {
	repo          repository.CustomerRepository
	whiteListRepo repository.WhitelistRepository
}

// Logout implements CustomerService.
func (service*customerServiceImpl) Logout(ctx context.Context) {
	err := service.whiteListRepo.Delete(domain.Whitelist{Id: ctx.Value("id").(string)})
	if err != nil {
		panic(exception.NewUnauthorizeError("login first"))
	}
}

// Login implements CustomerService.
func (service *customerServiceImpl) Login(ctx context.Context, request web.CustomerLoginRequest) string {

	customer, err := service.repo.FindByUsername(domain.Customer{Username: request.Username})
	if err != nil {
		panic(exception.NewUnauthorizeError("wrong username or password"))
	}

	if customer.Password != request.Password {
		panic(exception.NewUnauthorizeError("wrong username or password"))
	}

	whitelist, _ := service.whiteListRepo.FindById(domain.Whitelist{Id: customer.Id})

	if whitelist == (domain.Whitelist{}) {
		service.whiteListRepo.Save(domain.Whitelist{Id: customer.Id})
	}
	return helper.CreateToken(customer)
}

// FindById implements CustomerService.
func (service *customerServiceImpl) FindById(ctx context.Context, id string) web.CustomerResponse {
	customer, err := service.repo.FindById(domain.Customer{Id: id})
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.CustomerResponse{
		Id:       customer.Id,
		Name:     customer.Name,
		Username: customer.Username,
	}
}

// Save implements CustomerService.
func (service *customerServiceImpl) Save(ctx context.Context, request web.CustomerCreateRequest) {

	customer, _ := service.repo.FindByUsername(domain.Customer{Username: request.Username})

	if customer != (domain.Customer{}) {
		panic(exception.NewUnauthorizeError("username has been taken"))
	}

	service.repo.Save(domain.Customer{
		Id:       helper.GenerateUUID(),
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
	})
}

func NewCustomerService(repo repository.CustomerRepository, whitelistRepo repository.WhitelistRepository) CustomerService {
	return &customerServiceImpl{
		repo:          repo,
		whiteListRepo: whitelistRepo,
	}
}
