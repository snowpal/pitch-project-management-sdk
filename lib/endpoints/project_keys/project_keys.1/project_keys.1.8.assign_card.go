package projectKeys

import (
	"net/http"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func AssignCard(
	jwtToken string,
	reqBody request.AssignCardReqBody,
	projectCardParam common.ResourceIdParam,
) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysAssignCard,
		projectCardParam.CardId,
		projectCardParam.KeyId,
		projectCardParam.ProjectId,
	)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
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
