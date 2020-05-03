package main

import (
	"fmt"

	"github.com/biranjan/go_s3/s_3"
)

func main() {
	nameVar := "test"

	fmt.Println("hello", s_3.Hello(nameVar))
}
