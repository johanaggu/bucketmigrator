package migrator

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// uploader
type uploader struct {
	awsUploader *manager.Uploader
	bucket      string
}

func NewUploader(ctx context.Context, cfg Config) (*uploader, error) {
	conf, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretKey, "")))
	if err != nil {
		return &uploader{}, err
	}

	client := s3.NewFromConfig(conf)
	u := manager.NewUploader(client)

	return &uploader{
		awsUploader: u,
		bucket:      cfg.Bucket,
	}, nil
}

func (u *uploader) Upload(ctx context.Context, file io.Reader, key string) error {
	_, err := u.awsUploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(u.bucket),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}
