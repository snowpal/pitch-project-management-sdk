package projects

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetProject(jwtToken string, projectParam common.ResourceIdParam) (response.Project, error) {
	resProject := response.Project{}
	route, err := helpers2.GetRoute(lib.RouteProjectsGetProject, projectParam.ProjectId, projectParam.KeyId)
	if err != nil {
		return resProject, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resProject, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resProject, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resProject, err
	}

	err = json.Unmarshal(body, &resProject)
	if err != nil {
		return resProject, err
	}
	return resProject, nil
}
