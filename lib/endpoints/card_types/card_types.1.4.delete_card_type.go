package cardTypes

import (
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
)

func DeleteCardType(jwtToken string, cardTypeId string) error {
	route, err := helpers.GetRoute(lib.RouteCardTypesDeleteCardType, cardTypeId)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}
