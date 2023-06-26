package keys

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetKeyCardAndProjectScaleValues(jwtToken string, scaleParam request.ScaleIdParam) (response.ScaleValues, error) {
	resScaleValues := response.ScaleValues{}
	route, err := helpers2.GetRoute(
		lib.RouteKeysGetProjectScaleValues,
		scaleParam.KeyId,
		scaleParam.ScaleId,
	)
	if err != nil {
		return resScaleValues, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resScaleValues, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resScaleValues, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resScaleValues, err
	}

	err = json.Unmarshal(body, &resScaleValues)
	if err != nil {
		return resScaleValues, err
	}
	return resScaleValues, nil
}
