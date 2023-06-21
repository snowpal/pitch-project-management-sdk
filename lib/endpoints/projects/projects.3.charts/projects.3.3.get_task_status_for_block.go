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

func GetTaskStatusForProject(jwtToken string, taskParam common.ResourceIdParam) (response.TasksStatusProject, error) {
	resProjectTasksStatus := response.TasksStatusProject{}
	route, err := helpers2.GetRoute(
		lib.RouteProjectsGetTaskStatusForProject,
		taskParam.KeyId,
		taskParam.ProjectId,
	)
	if err != nil {
		return resProjectTasksStatus, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resProjectTasksStatus, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resProjectTasksStatus, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resProjectTasksStatus, err
	}

	err = json.Unmarshal(body, &resProjectTasksStatus)
	if err != nil {
		return resProjectTasksStatus, err
	}
	return resProjectTasksStatus, nil
}
