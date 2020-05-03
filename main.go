package main

import (
	"fmt"

	"gihub.com/biranjan/go_s3/s3"
)

func main() {
	nameVar := "test"

	fmt.Println("hello", s3.Hello(nameVar))
}
