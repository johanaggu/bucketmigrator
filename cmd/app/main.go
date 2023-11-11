package main

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/johanaggu/bucketmigrator/cmd/app/internal/migrator"
)

func main() {
	d, err := migrator.NewDownloader(context.Background(), migrator.Config{
		AccessKeyID: os.Getenv("DOWNLOADER_AKID"),
		SecretKey: os.Getenv("DOWNLOADER_SECRET_KEY"),
		Bucket: "origin-bucket-migrator",
		Region: "us-east-1",
	})
	if err != nil {
		log.Fatal("error creating downloader", err)
	}

	u, err := migrator.NewUploader(context.Background(), migrator.Config{
		AccessKeyID: os.Getenv("UPLOADER_AKID"),
		SecretKey: os.Getenv("UPLOADER_SECRET_KEY"),
		Bucket: "destiny-bucket-migrator",
		Region: "us-east-1",
	})
	if err != nil {
		log.Fatal("error creating uploader", err)
	}

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
