package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/slabdecoder"
	"github.com/johnfercher/taleslab/internal/slabencoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACzv369xFJgZGBgYGt3SdZ+fP7XfrVvghbvHxkgITAwicABHMDBAApAFijUOaMAAAAA=="

	slab, err := slabdecoder.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	slabBase64, err := slabencoder.Encode(slab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}

// |*|
//
//
//
// |*|
// Y = 8000
// 0xCE 0xFA 0xCE 0xD1 0x2 0x0 0x1 0x0 0x0 0x0 0x46 0x67 0x2C 0xE6 0xCF 0xCE 0xBF 0x46 0x8B 0x20 0xF8 0x17 0x38 0xF1 0xD2 0x20 0x2 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x40 0x1F 0x0 0x3 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x3 0x0 0x0
// 206 250 206 209 2 0 1 0 0 0 70 103 44 230 207 206 191 70 139 32 248 23 56 241 210 32 2 0 0 0 0 0 0 0 64 31 0 3 0 0 0 0 0 0 0 3 0 0

// |*|
//
//
// |*|
// Y = 25
// 0xCE 0xFA 0xCE 0xD1 0x2 0x0 0x1 0x0 0x0 0x0 0x46 0x67 0x2C 0xE6 0xCF 0xCE 0xBF 0x46 0x8B 0x20 0xF8 0x17 0x38 0xF1 0xD2 0x20 0x2 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x19 0x0 0x3 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x3 0x0 0x0
// 206 250 206 209 2 0 1 0 0 0 70 103 44 230 207 206 191 70 139 32 248 23 56 241 210 32 2 0 0 0 0 0 0 0 0 [25 0] 3 0 0 0 0 0 0 0 3 0 0

// |*|
//
// |*|
// Y = 4800
// 0xCE 0xFA 0xCE 0xD1 0x2 0x0 0x1 0x0 0x0 0x0 0x46 0x67 0x2C 0xE6 0xCF 0xCE 0xBF 0x46 0x8B 0x20 0xF8 0x17 0x38 0xF1 0xD2 0x20 0x2 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0xC0 0x12 0x0 0x3 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x3 0x0 0x0
// 206 250 206 209 2 0 1 0 0 0 70 103 44 230 207 206 191 70 139 32 248 23 56 241 210 32 2 0 0 0 0 0 0 0 [192 18] 0 3 0 0 0 0 0 0 0 3 0 0

// |*|
// |*|
// Y = 3200
// 0xCE 0xFA 0xCE 0xD1 0x2 0x0 0x1 0x0 0x0 0x0 0x46 0x67 0x2C 0xE6 0xCF 0xCE 0xBF 0x46 0x8B 0x20 0xF8 0x17 0x38 0xF1 0xD2 0x20 0x2 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x80 0xC 0x0 0x3 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x3 0x0 0x0
// 206 250 206 209 2 0 1 0 0 0 70 103 44 230 207 206 191 70 139 32 248 23 56 241 210 32 2 0 0 0 0 0 0 0 [128 12] 0 3 0 0 0 0 0 0 0 3 0 0

// |*||*| ->
// X = 200
// 0xCE 0xFA 0xCE 0xD1 0x2 0x0 0x1 0x0 0x0 0x0 0x46 0x67 0x2C 0xE6 0xCF 0xCE 0xBF 0x46 0x8B 0x20 0xF8 0x17 0x38 0xF1 0xD2 0x20 0x2 0x0 0x0 0x0 0xC8 0x0 0x0 0x0 0x0 0x0 0x0 0x3 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x3 0x0 0x0
// 206 250 206 209 2 0 1 0 0 0 70 103 44 230 207 206 191 70 139 32 248 23 56 241 210 32 2 0 0 0 [200 0] 0 0 0 0 0 3 0 0 0 0 0 0 0 3 0 0

// |*| |*| ->
// X = 300
// 0xCE 0xFA 0xCE 0xD1 0x2 0x0 0x1 0x0 0x0 0x0 0x46 0x67 0x2C 0xE6 0xCF 0xCE 0xBF 0x46 0x8B 0x20 0xF8 0x17 0x38 0xF1 0xD2 0x20 0x2 0x0 0x0 0x0 0x2C 0x1 0x0 0x0 0x0 0x0 0x0 0x3 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x3 0x0 0x0
// 206 250 206 209 2 0 1 0 0 0 70 103 44 230 207 206 191 70 139 32 248 23 56 241 210 32 2 0 0 0 [44 1] 0 0 0 0 0 3 0 0 0 0 0 0 0 3 0 0

// |*|  |*| ->
// X = 400
// 0xCE 0xFA 0xCE 0xD1 0x2 0x0 0x1 0x0 0x0 0x0 0x46 0x67 0x2C 0xE6 0xCF 0xCE 0xBF 0x46 0x8B 0x20 0xF8 0x17 0x38 0xF1 0xD2 0x20 0x2 0x0 0x0 0x0 0x90 0x1 0x0 0x0 0x0 0x0 0x0 0x3 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x3 0x0 0x0
// 206 250 206 209 2 0 1 0 0 0 70 103 44 230 207 206 191 70 139 32 248 23 56 241 210 32 2 0 0 0 [144 1] 0 0 0 0 0 3 0 0 0 0 0 0 0 3 0 0
