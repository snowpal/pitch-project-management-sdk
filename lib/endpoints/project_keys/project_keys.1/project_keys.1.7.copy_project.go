package projectKeys

import (
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func CopyProject(jwtToken string, projectParam request.CopyMoveProjectParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysCopyProject,
		projectParam.ProjectId,
		projectParam.KeyId,
		projectParam.TargetKeyId,
		strconv.FormatBool(projectParam.AllCards),
		strconv.FormatBool(projectParam.AllTasks),
	)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
	if err != nil {
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}
