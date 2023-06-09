package cards

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func UpdateCardScaleValue(
	jwtToken string,
	reqBody request.UpdateScaleValueReqBody,
	cardParam common.ResourceIdParam,
) (response.UpdateCardScaleValue, error) {
	resCardScaleValue := response.UpdateCardScaleValue{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resCardScaleValue, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteCardsUpdateCardScaleValue,
		cardParam.CardId,
		cardParam.KeyId,
		cardParam.ProjectId,
	)
	if err != nil {
		return resCardScaleValue, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return resCardScaleValue, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resCardScaleValue, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resCardScaleValue, err
	}

	err = json.Unmarshal(body, &resCardScaleValue)
	if err != nil {
		return resCardScaleValue, err
	}
	return resCardScaleValue, nil
}
