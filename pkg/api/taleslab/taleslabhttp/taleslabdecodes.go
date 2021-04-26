package taleslabhttp

import (
	"context"
	"encoding/json"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
	"io/ioutil"
	"net/http"
)

func DecodeMapRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err := apierror.New(http.StatusInternalServerError, err.Error())
		apierror.Log(ctx, err)
		return nil, err
	}
	inputMap := &entities.Map{}
	err = json.Unmarshal(bytes, inputMap)
	if err != nil {
		apiErr := apierror.New(http.StatusBadRequest, err.Error())
		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	return inputMap, nil
}
