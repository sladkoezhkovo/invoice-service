package handler

import (
	"context"
	"github.com/sladkoezhkovo/invoice-service/api"
	"github.com/sladkoezhkovo/invoice-service/internal/entity"
)

type InvoiceService interface {
	Create(invoice *entity.Invoice) error
}

type server struct {
	api.UnimplementedInvoiceServer
	service InvoiceService
}

func (s *server) Create(ctx context.Context, req *api.Request) (*api.Empty, error) {

	invoice := entity.Invoice{
		LastName:   req.LastName,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		OrgName:    req.OrgName,
	}

	switch req.Type {
	case api.AccountType_FACTORY:
		invoice.AccountType = entity.AccountTypeFactory
	case api.AccountType_SHOP:
		invoice.AccountType = entity.AccountTypeShop
	}

	if err := s.service.Create(&invoice); err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}
