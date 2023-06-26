package projectKeys

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

type CardByTemplateParam struct {
	ProjectListId string
	CardParam     request.CardByTemplateParam
}

func AddCardBasedOnTemplate(
	jwtToken string,
	reqBody request.AddCardReqBody,
	projectCardParam CardByTemplateParam,
) (response.Card, error) {
	resCard := response.Card{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resCard, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers2.GetRoute(
		lib.RouteProjectKeysAddCardBasedOnTemplate,
		*projectCardParam.CardParam.ProjectId,
		projectCardParam.CardParam.KeyId,
		projectCardParam.ProjectListId,
		projectCardParam.CardParam.TemplateId,
		strconv.FormatBool(projectCardParam.CardParam.ExcludeTasks),
	)
	if err != nil {
		return resCard, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		return resCard, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resCard, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resCard, err
	}

	err = json.Unmarshal(body, &resCard)
	if err != nil {
		return resCard, err
	}
	return resCard, err
}
