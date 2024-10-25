package application

import (
	"books_service/internal/core/cerror"
	"books_service/internal/core/dto"
	"books_service/internal/core/models"
	"context"
	"encoding/base64"
	"net/http"
	"slices"
)

var allowedMimeTypes = []string{"image/jpeg", "image/png", "image/webp"}

func (u *useCase) UploadPage(c context.Context, user *models.User, dto *dto.UploadPage) (string, error) {
	if err := u.repo.IsDraftExists(c, dto.ChapterID, user.ID); err != nil {
		return "", err
	}

	file, err := base64.StdEncoding.DecodeString(dto.Image)
	if err != nil {
		return "", err
	}

	if !slices.Contains(allowedMimeTypes, http.DetectContentType(file)) {
		return "", cerror.New(cerror.UNSUPPORTED_FORMAT, "this file format is not supporting")
	}

	filename, err := u.s3.Put(c, file)
	if err != nil {
		return "", err
	}

	return filename, u.repo.AddPageToDraft(c, dto.ChapterID, filename)
}
