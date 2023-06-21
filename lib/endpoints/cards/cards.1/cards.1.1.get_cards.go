package cards

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetCards(jwtToken string, cardsParam request.GetCardsParam) ([]response.Card, error) {
	resCards := response.Cards{}
	route, err := helpers.GetRoute(
		lib.RouteCardsGetCards,
		*cardsParam.ProjectId,
		strconv.Itoa(cardsParam.BatchIndex),
		cardsParam.KeyId,
	)
	if err != nil {
		return resCards.Cards, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resCards.Cards, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		return resCards.Cards, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resCards.Cards, err
	}

	err = json.Unmarshal(body, &resCards)
	if err != nil {
		return resCards.Cards, err
	}
	return resCards.Cards, nil
}
