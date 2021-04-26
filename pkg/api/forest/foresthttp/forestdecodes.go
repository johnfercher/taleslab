package foresthttp

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
	"io/ioutil"
	"net/http"
	"strings"
)

func DecodeForestRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err := apierror.New(http.StatusInternalServerError, err.Error())
		apierror.Log(ctx, err)
		return nil, err
	}
	forest := &entities.Map{}
	err = json.Unmarshal(bytes, forest)
	if err != nil {
		apiErr := apierror.New(http.StatusBadRequest, err.Error())
		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	/*err = validateForestRequest(ctx, forest)
	if err != nil {
		apiErr := apierror.New(http.StatusBadRequest, err.Error())
		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}*/

	return forest, nil
}

func validateForestRequest(ctx context.Context, forest *entities.Map) error {
	errMessages := []string{}

	if forest.Mountains.MinX+forest.Mountains.RandX > forest.Ground.Width {
		errMessages = append(errMessages, "The mountain X size is larger than the world X size.")
	}

	if forest.Mountains.MinY+forest.Mountains.RandY > forest.Ground.Width {
		errMessages = append(errMessages, "The mountain X size is larger than the world Y size.")
	}

	if len(errMessages) > 0 {
		return errors.New(strings.Join(errMessages, ","))
	}
	return nil
}
