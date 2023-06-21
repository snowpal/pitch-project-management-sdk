package projects

import (
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func DeleteProjectComment(jwtToken string, commentParam request.CommentIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectsDeleteProjectComment,
		*commentParam.CommentId,
		commentParam.KeyId,
		*commentParam.ProjectId,
	)
	req, err := http.NewRequest(http.MethodDelete, route, nil)
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
