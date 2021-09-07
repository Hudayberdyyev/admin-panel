package storage

import (
	"context"
	"io"
	"net/http"
	"os"

	"github.com/Hudayberdyyev/admin-panel-backend/logo"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

const (
	NewsBucketPattern    = "newsb"
	ContentBucketPattern = "contentb"
	NewsBucket           = "news"
	ContentBucket        = "content"
)

var ImageStorage *NewsStorage

type NewsStorage struct {
	Client *minio.Client
}

type MinioConfig struct {
	Endpoint       string
	AccessKeyId    string
	SecretAccesKey string
	UseSSL         bool
}

func NewNewsStorage(cfg MinioConfig) (*NewsStorage, error) {
	m, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyId, cfg.SecretAccesKey, ""),
		Secure: cfg.UseSSL,
	})
	return &NewsStorage{Client: m}, err
}

func (s *NewsStorage) UploadImage(ctx context.Context, bucketName string, filePath string, objectName string, authorsId int) error {
	location := "ap-south-1"

	err := s.Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location, ObjectLocking: false})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := s.Client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			// log.Printf("We already own %s\n", bucketName)
		} else {
			return err
		}
	} else {
		logrus.Printf("Successfully created %s\n", bucketName)
	}
	imageReader, err := getImageReader(filePath)
	if err != nil {
		var path string
		switch authorsId {
		case 1:
			path = logo.Turkmenportal
		case 2:
			path = logo.Rozetked
		case 3:
			path = logo.Wylsa
		case 4:
			path = logo.Championat
		case 5:
			path = logo.Ixbt
		}
		imageReader, err = os.Open(path)
		if err != nil {
			return err
		}
	}

	_, err = s.Client.PutObject(ctx, bucketName, objectName, imageReader, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	// fmt.Printf("Successfully uploaded bytes: %s\n", filePath)

	return nil
}

func (s *NewsStorage) RemoveImage(ctx context.Context, bucketName string, objectName string) error {
	err := s.Client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{GovernanceBypass: true})
	return err
}

func getImageReader(URL string) (io.Reader, error) {
	if resp, err := http.Get(URL); err != nil {
		return nil, err
	} else {
		return resp.Body, nil
	}
}
