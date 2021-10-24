package controllers

import (
	"net/http"
	"spos/stores/constants"
	"spos/stores/models"

	"github.com/labstack/echo"
	"github.com/s-pos/go-utils/utils/request"
	"github.com/s-pos/go-utils/utils/response"
)

func (c *controller) NewStoreHandler(e echo.Context) error {
	var (
		req     = e.Request()
		ctx     = req.Context()
		payload models.RequestNewStore
	)

	if err := request.BodyValidation(ctx, e, &payload, request.JSON); err != nil {
		return response.Errors(ctx, http.StatusBadRequest, string(constants.BodyRequired), message[constants.FailedCreatedStore], reason[constants.BodyRequired], err).Write(e)
	}

	return c.usecase.NewStore(ctx, payload).Write(e)
}
