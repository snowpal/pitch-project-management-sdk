package cards

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"
)

func GetCard(jwtToken string, podParam common.ResourceIdParam) (response.Card, error) {
	resCard := response.Card{}
	route, err := helpers2.GetRoute(lib.RouteCardsGetCard, podParam.CardId, podParam.KeyId, podParam.ProjectId)
	if err != nil {
		fmt.Println(err)
		return resCard, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
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
	return resCard, nil
}
