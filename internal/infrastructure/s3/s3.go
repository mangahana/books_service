package s3

import (
	"books_service/internal/core/configuration"
	"bytes"
	"context"
	"mime"
	"net/http"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type storage struct {
	client     *minio.Client
	bucketName string
}

func New(cfg *configuration.S3Config) (*storage, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	})

	if err != nil {
		return nil, err
	}

	return &storage{client: client, bucketName: cfg.BucketName}, nil
}

func (s *storage) Put(c context.Context, file []byte) (string, error) {
	opts := minio.PutObjectOptions{
		ContentType: http.DetectContentType(file),
	}

	exts, err := mime.ExtensionsByType("image/jpeg")
	if err != nil {
		return "", err
	}

	randomKey, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	filename := randomKey.String() + exts[0]

	_, err = s.client.PutObject(c, s.bucketName, filename, bytes.NewReader(file), int64(len(file)), opts)
	return filename, err
}

func (s *storage) Remove(c context.Context, filename string) error {
	return s.client.RemoveObject(c, s.bucketName, filename, minio.RemoveObjectOptions{})
}
