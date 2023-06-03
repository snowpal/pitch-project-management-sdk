package projectKeys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func AddCard(
	jwtToken string,
	reqBody request.AddCardReqBody,
	projectListParam request.ProjectListIdParam,
) (response.Card, error) {
	resCard := response.Card{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resCard, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers.GetRoute(
		lib.RouteProjectKeysAddACard,
		projectListParam.ProjectId,
		projectListParam.KeyId,
		projectListParam.ProjectListId,
	)
	if err != nil {
		fmt.Println(err)
		return resCard, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resCard, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCard, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCard, err
	}

	err = json.Unmarshal(body, &resCard)
	if err != nil {
		fmt.Println(err)
		return resCard, err
	}
	return resCard, nil
}
