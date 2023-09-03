package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/tessadem-sdk/pkg/tessadem"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	ctx := context.TODO()

	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	fmt.Println(apiKey)

	client := tessadem.NewClient(apiKey)
	writer := tessadem.NewFileWriter()

	request := &tessadem.AreaRequest{
		Units: tessadem.Meter,
		Northeast: &tessadem.Vector2D{
			X: -22.510677151874123,
			Y: -43.18595653686086,
		},
		Southwest: &tessadem.Vector2D{
			X: -22.503520257853395,
			Y: -43.170512111338226,
		},
	}

	response, err := client.GetArea(ctx, request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = writer.SaveArea(ctx, "file.json", response)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
