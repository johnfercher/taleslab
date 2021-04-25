package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/johnfercher/taleslab/internal/byteparser"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
)

func main() {
	test := [][]byte{
		{128, 56},
	}

	for _, value := range test {
		reader := bytes.NewReader(value)
		bufReader := bufio.NewReader(reader)

		value, err := byteparser.BufferToUint16(bufReader)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(slabdecoder.DecodeY(value))
	}
}
