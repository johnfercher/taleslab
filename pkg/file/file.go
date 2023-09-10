package file

import (
	"fmt"
	"os"
)

func SaveCodes(base64MatrixContents [][]string, pathFile string) error {
	f, err := os.Create(pathFile)
	if err != nil {
		return err
	}

	defer f.Close()

	contents := ""
	for _, base64Contents := range base64MatrixContents {
		for _, base := range base64Contents {
			content := "```" + base + "```"
			contents += content + "\n"
		}
		contents += "\n"
	}

	_, err = f.WriteString(contents)
	if err != nil {
		return err
	}

	fmt.Printf("File generated in %s", pathFile)

	return nil
}

func SaveJSON(json string, pathFile string) error {
	f, err := os.Create(pathFile)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(json)
	if err != nil {
		return err
	}

	fmt.Printf("File generated in %s", pathFile)

	return nil
}
