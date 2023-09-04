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

	client := tessadem.NewClient(apiKey)
	writer := tessadem.NewFileWriter()

	request := &tessadem.AreaRequest{
		Units: tessadem.Meter,
		Northeast: &tessadem.Vector2D{
			X: -22.490167623953077,
			Y: -43.07809776047625,
		},
		Southwest: &tessadem.Vector2D{
			X: -22.475654441563602,
			Y: -43.0492586515996,
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
