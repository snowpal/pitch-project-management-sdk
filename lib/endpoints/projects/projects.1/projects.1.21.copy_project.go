package projects

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func CopyProject(jwtToken string, projectParam request.CopyMoveProjectParam) (response.Project, error) {
	var resProject response.Project
	route, err := helpers2.GetRoute(
		lib.RouteProjectsCopyProject,
		projectParam.ProjectId,
		projectParam.KeyId,
		strconv.FormatBool(projectParam.AllTasks),
		strings.Join(projectParam.CardIds, ","),
		strconv.FormatBool(projectParam.AllCards),
		strconv.FormatBool(projectParam.AllChecklists),
		projectParam.TargetKeyId,
	)
	if err != nil {
		return resProject, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
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
