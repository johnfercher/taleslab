package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/johnfercher/taleslab/internal/byteparser"
	"log"
)

func main() {
	test := [][]byte{
		{0, 0, 200, 0},
		{0, 0, 44, 1},
		{0, 0, 144, 1},
	}

	for _, value := range test {
		reader := bytes.NewReader(value)
		bufReader := bufio.NewReader(reader)

		value, err := byteparser.BufferToFloat32(bufReader)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(value)
	}
}
