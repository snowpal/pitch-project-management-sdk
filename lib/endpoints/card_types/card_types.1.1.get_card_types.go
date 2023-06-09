package cardTypes

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetCardTypes(jwtToken string, includeCounts bool) ([]response.CardType, error) {
	resCardTypes := response.CardTypes{}
	route, err := helpers2.GetRoute(lib.RouteCardTypesGetCardTypes, strconv.FormatBool(includeCounts))
	if err != nil {
		return resCardTypes.CardTypes, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resCardTypes.CardTypes, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resCardTypes.CardTypes, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resCardTypes.CardTypes, err
	}

	err = json.Unmarshal(body, &resCardTypes)
	if err != nil {
		return resCardTypes.CardTypes, err
	}
	return resCardTypes.CardTypes, nil
}
