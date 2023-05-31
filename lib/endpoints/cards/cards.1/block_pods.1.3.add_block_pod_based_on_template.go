package cards

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"
)

func AddCardBasedOnTemplate(
	jwtToken string,
	reqBody request.AddCardReqBody,
	cardParam request.CardByTemplateParam,
) (response.Card, error) {
	resCard := response.Card{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resCard, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers2.GetRoute(
		lib.RouteCardsAddCardBasedOnTemplate,
		*cardParam.ProjectId,
		cardParam.KeyId,
		cardParam.TemplateId,
		strconv.FormatBool(cardParam.ExcludeTasks),
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCard, err
	}

	defer helpers2.CloseBody(res.Body)

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
	return resCard, err
}
