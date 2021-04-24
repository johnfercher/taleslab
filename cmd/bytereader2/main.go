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
		{175, 0},
		{192, 168},
		{128, 162},
		{64, 156},
		{150, 0},
		{192, 143},
		{128, 137},
	}

	for _, testx := range test {
		for _, testy := range testx {
			fmt.Printf("%b ", testy)
		}
		fmt.Println()
	}

	for _, value := range test {
		reader := bytes.NewReader(value)
		bufReader := bufio.NewReader(reader)

		value, err := byteparser.BufferToUint16(bufReader)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(value)
	}
}
