package main

import (
	"fmt"
	"github.com/qoqozhang/go-basic-test.git/input_static_params/version"
)

var mainParam string = "vMain"

func main() {
	fmt.Println("main: ", mainParam)
	fmt.Println("GitCommit: ", version.Version)
}
