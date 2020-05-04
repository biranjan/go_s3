package s_3

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/biranjan/go_s3/config"
)

// Function to print
func Hello(filename string) string {
	return filename
}

var sess = config.ConnectAWS()

// Function to test flag
func HandleUpload(filename string, bucketname string) string {
	file, err := os.Open(filename)
	status := "Successful"
	if err != nil {
		fmt.Println("error opening file: err:", err)
		status = "failed"
		os.Exit(1)
	}
	// upload file to s3
	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketname),
		Key:    aws.String(filename),
		Body:   file,
	})
	defer file.Close()
	if err != nil {
		fmt.Println("error uploading file file: err:", err)
		status = "failed"
	}

	return status
}

func handlerDownload(filename string, bucketname string) string {
	status := "Successful"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("error uploading file file: err:", err)
		status = "failed"
		os.Exit(1)
	}

	// Write the content of S3 object to the file
	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucketname),
		Key:    aws.String(filename),
	})

	defer f.Close()
	if err != nil {
		fmt.Println("error downloading file: err:", err)
		status = "failed"
	}
	return status

}
