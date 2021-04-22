package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/internal/taledecoder"
	"log"
)

func main() {
	forestBase64 := "H4sIAAAAAAAACzv369xFRgYmht71eR8Z67I99uXKL0oMm3OFhYGBoUWRL6kzRdClbaeS1eGIY94gMQaGAw4MDAvsGRianBgYGuwh7AZ7jv///wNprHIQfQKOEH4Fqj4w+wFUXw2KHIPz2nsQOQcgLoDKNSDpA7kFROegyYFAA1SuBlWOCVmuBIu+C0A5BgeIOxOg7AYHAKB2UtcoAQAA"

	board, err := taledecoder.DecodeSlab(forestBase64)
	if err != nil {
		log.Fatal(err)
	}

	boardBytes, err := json.Marshal(board)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(boardBytes))
}

const b = `
{
    "magichex": "CE FA CE D1 ",
    "Version": 1,
    "AssetsCount": 2,
    "Assets": [{
        "UUID": "8daf6ef1-017e-6b48-be6d-1fa261569cd4",
        "AssetsCount": 4,
        "Layout": [{
            "location": {
                "x": 6.0,
                "y": 1.25,
                "z": 65.0
            },
            "size": {
                "x": 1.0,
                "y": 1.25,
                "z": 1.0
            },
            "Rotation": 8
        }, {
            "location": {
                "x": 4.0,
                "y": 1.25,
                "z": 65.0
            },
            "size": {
                "x": 1.0,
                "y": 1.25,
                "z": 1.0
            },
            "Rotation": 4
        }, {
            "location": {
                "x": 9.0,
                "y": 1.25,
                "z": 62.0
            },
            "size": {
                "x": 1.0,
                "y": 1.25,
                "z": 1.0
            },
            "Rotation": 8
        }, {
            "location": {
                "x": 7.0,
                "y": 1.25,
                "z": 63.0
            },
            "size": {
                "x": 1.0,
                "y": 1.25,
                "z": 1.0
            },
            "Rotation": 0
        }]
    }, {
        "UUID": "84210e62-8964-1144-86b9-223ac358c64b",
        "AssetsCount": 4,
        "Layout": [{
            "location": {
                "x": 7.0,
                "y": 3.0,
                "z": 60.0
            },
            "size": {
                "x": 1.0,
                "y": 1.0,
                "z": 1.0
            },
            "Rotation": 8
        }, {
            "location": {
                "x": 6.0,
                "y": 1.0,
                "z": 59.0
            },
            "size": {
                "x": 1.0,
                "y": 1.0,
                "z": 1.0
            },
            "Rotation": 8
        }, {
            "location": {
                "x": 4.0,
                "y": 1.0,
                "z": 63.0
            },
            "size": {
                "x": 1.0,
                "y": 1.0,
                "z": 1.0
            },
            "Rotation": 8
        }, {
            "location": {
                "x": 4.0,
                "y": 1.0,
                "z": 61.0
            },
            "size": {
                "x": 1.0,
                "y": 1.0,
                "z": 1.0
            },
            "Rotation": 8
        }]
    }],
    "Bounds": {
        "location": {
            "x": 6.5,
            "y": 2.0,
            "z": 62.0
        },
        "size": {
            "x": 3.5,
            "y": 2.0,
            "z": 4.0
        },
        "Rotation": 0
    }
}
`
