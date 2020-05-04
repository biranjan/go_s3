package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/biranjan/go_s3/s_3"
)

func main() {

	//flag.StringVar(&nameVar, "upload", "tst", "file path to upload")
	//downloadFile := flag.String("download", "", "file name to download")
	uploadCmd := flag.NewFlagSet("upload", flag.ExitOnError)
	uploadFile := uploadCmd.String("filename", "", "filename")

	downloadCmd := flag.NewFlagSet("download", flag.ExitOnError)
	downloadFile := downloadCmd.String("filename", "", "filename")

	if len(os.Args) < 2 {
		fmt.Println("expected 'upload' or 'download' subcommands")
		os.Exit(1)
	}

	bucketname := "mytestbucketbi"

	switch os.Args[1] {
	case "upload":
		uploadCmd.Parse(os.Args[2:])
		fmt.Println("hello", s_3.HandleUpload(*uploadFile, bucketname))
	case "download":
		downloadCmd.Parse(os.Args[2:])
		fmt.Println("Upload", s_3.HandleUpload(*downloadFile, bucketname))
	default:
		fmt.Println("Expected 'download' or 'upload' command")
		os.Exit(1)

	}
}
