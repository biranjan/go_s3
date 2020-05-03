package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/biranjan/go_s3/s_3"
)

func main() {
	//nameVar := "test"
	uploadFile := flag.String("upload", "", "file path to upload")
	downloadFile := flag.String("download", "", "file name to download")

	switch os.Args[1] {
	case "upload":
		flag.Parse()
		fmt.Println("hello", s_3.Hello(*uploadFile))
	case "download":
		flag.Parse()
		fmt.Println("hello", s_3.HandleUpload(*downloadFile))
	default:
		fmt.Println("Expected 'download' or 'upload' command")
		os.Exit(1)

	}
}
