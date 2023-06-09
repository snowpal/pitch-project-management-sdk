package projects

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetProjectsAvailableToBeLinkedToThisKey(jwtToken string, keyId string) ([]response.Project, error) {
	resProjects := response.Projects{}
	route, err := helpers2.GetRoute(lib.RouteProjectsGetProjectsAvailableToBeLinkedToThisKey, keyId)
	if err != nil {
		return resProjects.Projects, err
	}

	req, _ := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resProjects.Projects, err
	}

	helpers2.AddUserHeaders(jwtToken, req)
	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resProjects.Projects, err
	}

	defer helpers2.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		return resProjects.Projects, err
	}

	err = json.Unmarshal(body, &resProjects)
	if err != nil {
		return resProjects.Projects, err
	}
	return resProjects.Projects, nil
}
