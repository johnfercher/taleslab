package main

import "fmt"

func main() {
	offset := 108
	remainder := offset % 100

	fmt.Println(offset - remainder)
	fmt.Println(remainder)
}
