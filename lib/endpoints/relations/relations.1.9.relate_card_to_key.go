package relations

import (
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func relateKeyToCard(jwtToken string, route string) error {
	req, err := http.NewRequest(http.MethodPatch, route, nil)
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

func RelateKeyToCard(jwtToken string, relationParam request.KeyToCardRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsRelateCardToKey,
		relationParam.KeyId,
		relationParam.TargetCardId,
		relationParam.TargetKeyId,
		relationParam.TargetProjectId,
	)
	if err != nil {
		return err
	}
	err = relateKeyToCard(jwtToken, route)
	if err != nil {
		return err
	}
	return nil
}
