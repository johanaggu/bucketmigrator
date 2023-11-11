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

type downloader struct {
	awsDownloader *manager.Downloader
	bucket      string
}

func NewDownloader(ctx context.Context, cfg Config) (*downloader, error) {
	conf, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretKey, "")))
	if err != nil {
		return &downloader{}, err
	}

	client := s3.NewFromConfig(conf)
	u := manager.NewDownloader(client)

	return &downloader{
		awsDownloader: u,
		bucket:      cfg.Bucket,
	}, nil
}

func (d *downloader) Download(ctx context.Context, file io.WriterAt, key string) error {
	_, err := d.awsDownloader.Download(ctx, file, &s3.GetObjectInput{
		Bucket: aws.String(d.bucket),
		Key:    aws.String(key),
	})

	return err
}
