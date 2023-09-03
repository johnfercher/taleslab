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
			X: -27.44524479006217,
			Y: -48.519853037517144,
		},
		Southwest: &tessadem.Vector2D{
			X: -27.459521825397537,
			Y: -48.55104453459989,
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
