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
		{0, 0},
		{64, 0},
		{128, 0},
		{192, 0},
		{0, 1},
		{64, 1},
		{128, 1},
		{192, 1},
		{0, 2},
		{64, 2},
		{128, 2},
		{192, 2},
		{0, 3},
		{64, 3},
		{128, 3},
		{192, 3},
		{0, 4},
		{64, 4},
		{128, 4},
		{192, 4},
		{0, 5},
		{64, 5},
		{128, 5},
		{192, 5},
		{0, 6},
	}

	fmt.Println(320.0 / 1536.0 * 360.0)

	for _, value := range test {
		reader := bytes.NewReader(value)
		bufReader := bufio.NewReader(reader)

		value, err := byteparser.BufferToInt16(bufReader)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(value)
	}
}
