package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"spos/stores/constants"
	"spos/stores/models"

	"github.com/s-pos/go-utils/utils/response"
)

func (u *usecase) NewStore(ctx context.Context, req models.RequestNewStore) response.Output {
	_, err := u.repo.FindStoreByName(req.Name)
	if err == nil {
		err = fmt.Errorf("name store already taken")
		return response.Errors(ctx, http.StatusBadRequest, string(constants.StoreNameAlreadyTaken), message[constants.FailedCreatedStore], reason[constants.StoreNameAlreadyTaken], err)
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return response.Errors(ctx, http.StatusInternalServerError, string(constants.ErrorFetch), message[constants.FailedCreatedStore], errGlobal, err)
	}

	_, err = u.repo.FindStoreByName(req.Username)
	if err == nil {
		err = fmt.Errorf("domain store already taken")
		return response.Errors(ctx, http.StatusBadRequest, string(constants.StoreDomainAlreadyTaken), message[constants.FailedCreatedStore], reason[constants.StoreDomainAlreadyTaken], err)
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return response.Errors(ctx, http.StatusInternalServerError, string(constants.ErrorFetch), message[constants.FailedCreatedStore], errGlobal, err)
	}

	store := models.NewStore()
	store.SetName(req.Name)
	store.SetType(req.Type)
	store.SetSource(req.Source)
	store.SetDescription(req.Description)
	store.SetDomain(req.Username)

	errCode, errMessage, errReason, code, err := uploadToGoogleStorage(ctx, store, req.Image)
	if err != nil {
		return response.Errors(ctx, code, errCode, errMessage, errReason, err)
	}

	store, err = u.repo.InsertStore(store)
	if err != nil {
		return response.Errors(ctx, http.StatusInternalServerError, string(constants.ErrorInsert), message[constants.FailedCreatedStore], errGlobal, err)
	}

	return response.Success(ctx, http.StatusCreated, string(constants.SuccessCreateStore), message[constants.SuccessCreateStore], store)
}
