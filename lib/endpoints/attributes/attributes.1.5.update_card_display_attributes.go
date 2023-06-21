package attributes

import (
	"net/http"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func UpdateCardAttrs(jwtToken string, cardParam common.ResourceIdParam, attribute request.ResourceAttributeReqBody) error {
	requestBody, err := helpers.GetRequestBody(attribute)
	if err != nil {
		return err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers.GetRoute(
		lib.RouteAttributesUpdateCardDisplayAttributes,
		cardParam.CardId,
		cardParam.KeyId,
		cardParam.ProjectId,
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
