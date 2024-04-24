package main

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/johanaggu/bucketmigrator/cmd/app/internal/migrator"
)

func main() {
	// Downloader client
	downloaderAKID := os.Getenv("DOWNLOADER_AKID")
	downloaderSecretKey := os.Getenv("DOWNLOADER_SECRET_KEY")
	downloaderBucketName := "origin-bucket-migrator"
	downloaderBucketRegion := "us-east-1"

	d, err := migrator.NewDownloader(context.Background(), migrator.Config{
		AccessKeyID: downloaderAKID,
		SecretKey:   downloaderSecretKey,
		Bucket:      downloaderBucketName,
		Region:      downloaderBucketRegion,
	})
	if err != nil {
		log.Fatal("error creating downloader", err)
	}

	// Uploader client
	uploaderAKID := os.Getenv("UPLOADER_AKID")
	uploaderSecretKey := os.Getenv("UPLOADER_SECRET_KEY")
	uploaderBucketName := "destiny-bucket-migrator"
	uploaderBucketRegion := "us-east-1"

	u, err := migrator.NewUploader(context.Background(), migrator.Config{
		AccessKeyID: uploaderAKID,
		SecretKey:   uploaderSecretKey,
		Bucket:      uploaderBucketName,
		Region:      uploaderBucketRegion,
	})
	if err != nil {
		log.Fatal("error creating uploader", err)
	}

	// all objects to migrate
	key := "motorcycles/duke.jpg"

	f, err := os.Create(path.Join(".tmp", key))
	if err != nil {
		log.Fatal("error creating file", err)
	}

	err = d.Download(context.Background(), f, key)
	if err != nil {
		log.Fatal("error  download file", err)
	}

	f, err = os.Open(path.Join(".tmp", key))
	if err != nil {
		log.Fatal("error open file", err)
	}

	err = u.Upload(context.Background(), f, key)
	if err != nil {
		log.Fatal("error upload file", err)
	}

}
