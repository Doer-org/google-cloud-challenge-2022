//go:build ignore

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	openapi, err := os.OpenFile("../../../../openapi/openapi.json",os.O_RDWR,0666)
	if err != nil {
		panic(
			fmt.Sprintf("openapi.json not found : %v",err),
		)
	}
	defer openapi.Close()

	generate_openapi, err := os.OpenFile("openapi.json",os.O_RDWR,0666)
	if err != nil {
		panic(
			fmt.Sprintf("generate openapi.json not found : %v",err),
		)
	}
	defer generate_openapi.Close()

	_,err = io.Copy(openapi, generate_openapi)
	if err != nil {
		panic(
			fmt.Sprintf("coopy failed : %v",err),
		)
	}
}
