package cloud_storage_utils

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"gopkg.in/mgo.v2/bson"
)

func UploadAndGetPublicUrl(bucketName string, folder string, file *multipart.FileHeader) (string, error) {
	objectId := bson.NewObjectId().Hex()
	object := folder + objectId
	client, err := storage.NewClient(context.Background())
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	filePointer, _ := file.Open()
	wc := bucket.Object(object).NewWriter(context.Background())
	if _, err = io.Copy(wc, filePointer); err != nil {
		fmt.Println(err)
		return "", err
	}
	if err := wc.Close(); err != nil {
		fmt.Println(err)
		return "", err
	}

	url := "https://storage.googleapis.com/" + bucketName + "/" + object
	return url, nil
}
