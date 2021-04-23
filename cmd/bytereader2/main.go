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
		{0, 6},
	}

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
