package cards

import (
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func AddScaleToCard(jwtToken string, cardParam request.ScaleIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteCardsAddScaleToCard,
		*cardParam.CardId,
		cardParam.ScaleId,
		cardParam.KeyId,
		*cardParam.ProjectId,
	)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
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
