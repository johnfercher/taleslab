package file

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"net/http"
	"os"
)

func SaveContentInFile(base64Content string, pathFile string) apierror.ApiError {
	f, err := os.Create(pathFile)
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
	}

	defer f.Close()

	_, err = f.WriteString("```" + base64Content + "```")
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
	}

	fmt.Printf("File generated in %s", pathFile)

	return nil
}
