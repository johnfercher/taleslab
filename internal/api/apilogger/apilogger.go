package apilogger

import (
	"context"
	"fmt"
)

func Info(_ context.Context, message string) {
	fmt.Printf("INFO: %s\n", message)
}

func Error(_ context.Context, message string) {
	fmt.Printf("ERROR: %s\n", message)
}

func Warn(_ context.Context, message string) {
	fmt.Printf("WARN: %s\n", message)
}
