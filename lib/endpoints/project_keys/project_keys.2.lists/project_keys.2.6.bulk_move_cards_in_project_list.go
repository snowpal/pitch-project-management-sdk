package projectKeys

import (
	"net/http"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func BulkMoveCardsInProjectList(jwtToken string, projectListParam request.CopyMoveProjectListCardsParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysBulkMoveCardsInProjectList,
		projectListParam.ProjectListId,
		projectListParam.KeyId,
		projectListParam.ProjectId,
		projectListParam.TargetKeyId,
		projectListParam.TargetProjectId,
		projectListParam.TargetProjectListId,
		strings.Join(projectListParam.CardIds, ","),
	)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
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
