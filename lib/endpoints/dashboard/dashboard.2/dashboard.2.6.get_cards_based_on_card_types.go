package dashboard

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetCardsBasedOnCardTypes(jwtToken string) ([]response.CardTypesKey, error) {
	resCardTypesKeys := response.CardTypesKeys{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetCardsBasedOnCardTypes)
	if err != nil {
		fmt.Println(err)
		return *resCardTypesKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return *resCardTypesKeys.Keys, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return *resCardTypesKeys.Keys, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return *resCardTypesKeys.Keys, err
	}

	err = json.Unmarshal(body, &resCardTypesKeys)
	if err != nil {
		fmt.Println(err)
		return *resCardTypesKeys.Keys, err
	}
	return *resCardTypesKeys.Keys, nil
}
