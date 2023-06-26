package cards

import (
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func DeleteCardTask(jwtToken string, taskParam request.TaskIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteCardsDeleteCardTask,
		*taskParam.TaskId,
		taskParam.KeyId,
		*taskParam.ProjectId,
		*taskParam.CardId,
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
