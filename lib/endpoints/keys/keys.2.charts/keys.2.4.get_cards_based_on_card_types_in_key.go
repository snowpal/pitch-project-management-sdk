package keys

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetCardsBasedOnCardTypesInKey(jwtToken string, keyId string) (response.CardTypesKey, error) {
	resCardTypesKey := response.CardTypesKey{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetCardsBasedOnCardTypesInKey, keyId)
	if err != nil {
		return resCardTypesKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resCardTypesKey, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resCardTypesKey, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resCardTypesKey, err
	}

	err = json.Unmarshal(body, &resCardTypesKey)
	if err != nil {
		return resCardTypesKey, err
	}
	return resCardTypesKey, nil
}
