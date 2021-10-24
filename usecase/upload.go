package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"spos/stores/constants"
	"spos/stores/models"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/s-pos/go-utils/logger"
	"google.golang.org/api/option"
)

func uploadToGoogleStorage(ctx context.Context, store *models.Store, reqFile string) (string, string, string, int, error) {
	var (
		imageName      = uuid.New().String()
		err            error
		bucket         = os.Getenv("GCS_BUCKET_NAME")
		credentialPath = os.Getenv("GCS_PATH")
	)

	if store.IsSourceOffline() && store.IsTypeOffline() && reflect.ValueOf(reqFile).IsZero() {
		err = fmt.Errorf("image not found")
		logger.FieldMandatory("image", "Image is required").Body(ctx)
		return string(constants.StoreImageNotFound), message[constants.FailedCreatedStore], reason[constants.StoreImageNotFound], http.StatusBadRequest, err
	}

	fileDecode, err := base64.StdEncoding.DecodeString(reqFile)
	if err != nil {
		return string(constants.ImageNotEncoded), message[constants.FailedCreatedStore], reason[constants.ImageNotEncoded], http.StatusBadRequest, err
	}

	imageName = fmt.Sprintf("%s.jpg", strings.ReplaceAll(imageName, "-", ""))

	err = ioutil.WriteFile(imageName, fileDecode, 0666)
	if err != nil {
		return string(constants.ImageUploadFailed), message[constants.FailedCreatedStore], reason[constants.ImageUploadFailed], http.StatusInternalServerError, err
	}

	file, err := os.Open(imageName)
	if err != nil {
		return string(constants.ImageUploadFailed), message[constants.FailedCreatedStore], reason[constants.ImageUploadFailed], http.StatusInternalServerError, err
	}
	defer file.Close()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialPath))
	if err != nil {
		return string(constants.ImageUploadFailed), message[constants.FailedCreatedStore], reason[constants.ImageUploadFailed], http.StatusInternalServerError, err
	}

	sw := client.Bucket(bucket).Object(fmt.Sprintf("store/%s", imageName)).NewWriter(ctx)
	defer sw.Close()

	_, err = io.Copy(sw, file)
	if err != nil {
		return string(constants.ImageUploadFailed), message[constants.FailedCreatedStore], reason[constants.ImageUploadFailed], http.StatusInternalServerError, err
	}

	store.SetLogo(imageName)

	os.Remove(file.Name())
	return "", "", "", http.StatusCreated, nil
}
