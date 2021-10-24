package usecase

import (
	"context"
	"spos/stores/constants"
	"spos/stores/models"
	"spos/stores/repository"

	"github.com/s-pos/go-utils/utils/response"
)

const errGlobal = constants.Error

var (
	message = constants.Message
	reason  = constants.Reason
)

type usecase struct {
	repo repository.Repository
}

type Usecase interface {
	// NewStore will create new store with type 'Offline' or 'Online'
	// if type is 'Online' then source store will be 'Tokopedia' or 'Shopee'
	NewStore(ctx context.Context, req models.RequestNewStore) response.Output
}

func NewUsecase(r repository.Repository) Usecase {
	return &usecase{
		repo: r,
	}
}
