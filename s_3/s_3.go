package s_3

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	AWS_S3_REGION = "eu-west-2"
	AWS_S3_BUCKET = ""
)

// ConnectAWS connect to aws
func ConnectAWS() *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		panic(err)
	}
	return sess
}

// Hello test function to print
func Hello(filename string) string {
	key := strings.Split(filename, "/")
	return key[0]
}

var sess = ConnectAWS()

// HandleUpload function to upload local file to s3
func HandleUpload(filename string, bucketname string, keyname string, wg *sync.WaitGroup) {

	defer wg.Done()
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

	fmt.Println(filename, ":", status)
}

// HandleDownload function to download file from s3
func HandleDownload(filename string, bucketname string, filepath string, wg *sync.WaitGroup) {
	defer wg.Done()
	status := "Successful"
	if filepath == "" {
		filepath = filename
	}
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println("error creating file: err:", err)
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
	fmt.Println(filename, ":", status)

}
