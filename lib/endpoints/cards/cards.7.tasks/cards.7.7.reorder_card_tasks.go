package cards

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func ReorderCardTasks(
	jwtToken string,
	reqBody request.ReorderTasksReqBody,
	taskParam request.TaskIdParam,
) ([]response.Task, error) {
	resTasks := response.Tasks{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resTasks.Tasks, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteCardsReorderCardTasks,
		*taskParam.CardId,
		taskParam.KeyId,
		*taskParam.ProjectId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return resTasks.Tasks, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resTasks.Tasks, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resTasks.Tasks, err
	}

	err = json.Unmarshal(body, &resTasks)
	if err != nil {
		return resTasks.Tasks, err
	}
	return resTasks.Tasks, nil
}
