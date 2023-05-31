package cards

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	helpers2 "github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"
)

func GetCards(jwtToken string, podsParam request.GetCardsParam) ([]response.Card, error) {
	resCards := response.Cards{}
	route, err := helpers.GetRoute(
		lib.RouteCardsGetCards,
		*podsParam.ProjectId,
		strconv.Itoa(podsParam.BatchIndex),
		podsParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resCards.Cards, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resCards.Cards, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCards.Cards, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCards.Cards, err
	}

	err = json.Unmarshal(body, &resCards)
	if err != nil {
		fmt.Println(err)
		return resCards.Cards, err
	}
	return resCards.Cards, nil
}
