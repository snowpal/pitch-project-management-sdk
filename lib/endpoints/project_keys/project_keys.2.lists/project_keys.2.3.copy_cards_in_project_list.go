package projectKeys

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func CopyCardsInProjectList(jwtToken string, projectListParam request.CopyMoveProjectListCardsParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysCopyCardsInProjectList,
		projectListParam.ProjectListId,
		projectListParam.KeyId,
		projectListParam.ProjectId,
		projectListParam.TargetKeyId,
		projectListParam.TargetProjectId,
		projectListParam.TargetProjectListId,
		strconv.FormatBool(projectListParam.AllCards),
		strconv.FormatBool(projectListParam.AllTasks),
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
