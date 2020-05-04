package s_3

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	AWS_S3_REGION = "eu-west-2"
	AWS_S3_BUCKET = ""
)

func ConnectAWS() *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		panic(err)
	}
	return sess
}

// Function to print
func Hello(filename string) string {
	key := strings.Split(filename, "/")
	return key[0]
}

var sess = ConnectAWS()

// Function to test flag
func HandleUpload(filename string, bucketname string, keyname string) string {
	file, err := os.Open(filename)

	if keyname == "" {
		key := strings.Split(filename, "/")
		keyname = key[len(key)-1]

	}

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
		Key:    aws.String(keyname),
		Body:   file,
	})
	defer file.Close()
	if err != nil {
		fmt.Println("error uploading file file: err:", err)
		status = "failed"
	}

	return status
}

func HandleDownload(filename string, bucketname string, filepath string) string {
	status := "Successful"
	f, err := os.Create(filepath)
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
