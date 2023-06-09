package projects

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetProjects(jwtToken string, projectParam request.GetProjectsParam) ([]response.Project, error) {
	resProjects := response.Projects{}
	route, err := helpers2.GetRoute(
		lib.RouteProjectsGetProjects,
		projectParam.KeyId,
		projectParam.Filter,
		strconv.Itoa(projectParam.BatchIndex),
		strconv.FormatBool(projectParam.WriteOrHigherAcl),
	)
	if err != nil {
		return resProjects.Projects, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resProjects.Projects, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resProjects.Projects, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resProjects.Projects, err
	}

	err = json.Unmarshal(body, &resProjects)
	if err != nil {
		return resProjects.Projects, err
	}
	return resProjects.Projects, nil
}
