package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/tessadem-sdk/pkg/tessadem"
)

func main() {
	ctx := context.TODO()

	fileReader := tessadem.NewFileReader()

	response, err := fileReader.ReadArea(ctx, "file.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(response)
}
