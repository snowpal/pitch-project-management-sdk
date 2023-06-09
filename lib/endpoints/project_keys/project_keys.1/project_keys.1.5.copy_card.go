package projectKeys

import (
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func CopyCard(jwtToken string, projectCardParam request.CopyMoveProjectCardParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysCopyCard,
		projectCardParam.CardId,
		projectCardParam.KeyId,
		projectCardParam.ProjectId,
		projectCardParam.TargetKeyId,
		projectCardParam.TargetProjectId,
		projectCardParam.TargetProjectListId,
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
