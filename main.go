package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/biranjan/go_s3/s_3"
)

func main() {

	//flag.StringVar(&nameVar, "upload", "tst", "file path to upload")
	//downloadFile := flag.String("download", "", "file name to download")
	uploadCmd := flag.NewFlagSet("upload", flag.ExitOnError)
	//uploadFile := uploadCmd.String("filename", "", "filename")
	uploadKey := uploadCmd.String("key", "", "upload key for s3")

	downloadCmd := flag.NewFlagSet("download", flag.ExitOnError)
	//downloadFile := downloadCmd.String("filename", "", "filename")
	downloadPath := downloadCmd.String("filepath", "", "filename")

	if len(os.Args) < 2 {
		fmt.Println("expected 'upload' or 'download' subcommands")
		os.Exit(1)
	}

	bucketname := os.Getenv("my_bucket") // mytestbucketbi
	var wg sync.WaitGroup

	switch os.Args[1] {
	case "upload":
		uploadCmd.Parse(os.Args[2:])
		//fmt.Println("Upload:", s_3.HandleUpload(*uploadFile, bucketname, *uploadKey))
		fileList := uploadCmd.Args()
		for i := 0; i < len(fileList); i++ {
			wg.Add(1)
			go s_3.HandleUpload(fileList[i], bucketname, *uploadKey, &wg)
		}
		wg.Wait()
	case "download":
		downloadCmd.Parse(os.Args[2:])
		fileList := downloadCmd.Args()
		for i := 0; i < len(fileList); i++ {
			wg.Add(1)
			go s_3.HandleDownload(fileList[i], bucketname, *downloadPath, &wg)
		}
		wg.Wait()
		//fmt.Println("Download:", s_3.HandleDownload(*downloadFile, bucketname, *downloadPath))
	default:
		fmt.Println("Expected 'download' or 'upload' command")
		os.Exit(1)

	}
}
