package relations

import (
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func unrelateKeyFromCard(jwtToken string, route string) error {
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

func UnrelateKeyFromKeyCard(jwtToken string, relationParam request.ProjectToCardRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelateCardFromKey,
		relationParam.ProjectId,
		relationParam.TargetCardId,
		relationParam.TargetKeyId,
	)
	if err != nil {
		return err
	}
	err = unrelateKeyFromCard(jwtToken, route)
	if err != nil {
		return err
	}
	return nil
}

func UnrelateKeyFromCard(jwtToken string, relationParam request.ProjectToCardRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelateCardFromKey,
		relationParam.ProjectId,
		relationParam.TargetCardId,
		relationParam.TargetKeyId,
		relationParam.TargetProjectId,
	)
	if err != nil {
		return err
	}
	err = unrelateKeyFromCard(jwtToken, route)
	if err != nil {
		return err
	}
	return nil
}
