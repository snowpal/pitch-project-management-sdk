package dashboard

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetProjectsBasedOnProjectTypes(jwtToken string) ([]response.ProjectTypesKey, error) {
	resProjectTypesKeys := response.ProjectTypesKeys{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetProjectsBasedOnProjectTypes)
	if err != nil {
		return resProjectTypesKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resProjectTypesKeys.Keys, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resProjectTypesKeys.Keys, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resProjectTypesKeys.Keys, err
	}

	err = json.Unmarshal(body, &resProjectTypesKeys)
	if err != nil {
		return resProjectTypesKeys.Keys, err
	}
	return resProjectTypesKeys.Keys, nil
}
