package controllers

import (
	"spos/stores/constants"
	"spos/stores/usecase"

	"github.com/labstack/echo"
)

var (
	message = constants.Message
	reason  = constants.Reason
)

type controller struct {
	usecase usecase.Usecase
}

type Controller interface {
	// NewStoreHandler handler for create new store
	NewStoreHandler(e echo.Context) error
}

func NewController(uc usecase.Usecase) Controller {
	return &controller{
		usecase: uc,
	}
}
