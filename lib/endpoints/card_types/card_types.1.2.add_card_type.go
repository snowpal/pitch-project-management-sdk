package cardTypes

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

func AddCardType(jwtToken string, reqBody request.CardTypeReqBody) (response.CardType, error) {
	resCardType := response.CardType{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resCardType, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteCardTypesAddCardType)
	if err != nil {
		return resCardType, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		return resCardType, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resCardType, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resCardType, err
	}

	err = json.Unmarshal(body, &resCardType)
	if err != nil {
		return resCardType, err
	}
	return resCardType, nil
}
