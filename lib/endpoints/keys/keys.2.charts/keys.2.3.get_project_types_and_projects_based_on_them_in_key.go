package keys

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetProjectTypesAndProjectsBasedOnThemInKey(jwtToken string, keyId string) (response.ProjectTypesKey, error) {
	resProjectTypesKey := response.ProjectTypesKey{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetProjectTypesAndProjectsBasedOnThemInKey, keyId)
	if err != nil {
		return resProjectTypesKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resProjectTypesKey, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resProjectTypesKey, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resProjectTypesKey, err
	}

	err = json.Unmarshal(body, &resProjectTypesKey)
	if err != nil {
		return resProjectTypesKey, err
	}
	return resProjectTypesKey, nil
}
