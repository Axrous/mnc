package service

import (
	"context"

	"github.com/Axrous/mnc/helper"
	"github.com/Axrous/mnc/model/domain"
	"github.com/Axrous/mnc/model/web"
	"github.com/Axrous/mnc/repository"
)

type TransactionService interface {
	Payment(ctx context.Context, request web.TransactionPaymentRequest)
}

type transactionServiceImpl struct {
	repo repository.TransactionRepository
	whitetlistRepo repository.WhitelistRepository
}

// Payment implements TransactionService.
func (service *transactionServiceImpl) Payment(ctx context.Context, request web.TransactionPaymentRequest) {
	_, err := service.whitetlistRepo.FindById(domain.Whitelist{Id: ctx.Value("id").(string)})
	helper.PanicIfError(err)

	service.repo.Save(domain.Transaction{
		Id:       "1",
		Customer: domain.Customer{
			Id:       ctx.Value("id").(string),
		},
		Merchant: domain.Merchant{
			Id:   request.MerchantId,
		},
		Amount:   request.Amount,
	})
}

func NewTransactionService(repo repository.TransactionRepository, whitelistRepo repository.WhitelistRepository) TransactionService {
	return &transactionServiceImpl{
		repo:           repo,
		whitetlistRepo: whitelistRepo,
	}
}
