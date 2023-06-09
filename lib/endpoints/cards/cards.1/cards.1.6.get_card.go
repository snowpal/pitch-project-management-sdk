package cards

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetCard(jwtToken string, cardParam common.ResourceIdParam) (response.Card, error) {
	resCard := response.Card{}
	route, err := helpers2.GetRoute(lib.RouteCardsGetCard, cardParam.CardId, cardParam.KeyId, cardParam.ProjectId)
	if err != nil {
		return resCard, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
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
	return resCard, nil
}
