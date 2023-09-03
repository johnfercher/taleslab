package file

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"net/http"
	"os"
)

func SaveCodes(base64MatrixContents [][]string, pathFile string) apierror.ApiError {
	f, err := os.Create(pathFile)
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
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
		return apierror.New(http.StatusInternalServerError, err.Error())
	}

	fmt.Printf("File generated in %s", pathFile)

	return nil
}

func SaveJSON(json string, pathFile string) apierror.ApiError {
	f, err := os.Create(pathFile)
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
	}

	defer f.Close()

	_, err = f.WriteString(json)
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
	}

	fmt.Printf("File generated in %s", pathFile)

	return nil
}
