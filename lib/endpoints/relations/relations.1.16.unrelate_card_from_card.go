package relations

import (
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func unrelateCardFromCard(jwtToken string, route string) error {
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

func UnrelateKeyCardFromKeyCard(jwtToken string, relationParam request.CardToCardRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelateCardFromCard,
		relationParam.CardId,
		relationParam.SourceKeyId,
		relationParam.TargetCardId,
		relationParam.TargetKeyId,
	)
	if err != nil {
		return err
	}
	err = unrelateCardFromCard(jwtToken, route)
	if err != nil {
		return err
	}
	return nil
}

func UnrelateKeyCardFromCard(jwtToken string, relationParam request.CardToCardRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelateCardFromCard,
		relationParam.CardId,
		relationParam.SourceKeyId,
		relationParam.TargetCardId,
		relationParam.TargetKeyId,
		relationParam.TargetProjectId,
	)
	if err != nil {
		return err
	}
	err = unrelateCardFromCard(jwtToken, route)
	if err != nil {
		return err
	}
	return nil
}

func UnrelateCardFromCard(jwtToken string, relationParam request.CardToCardRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelateCardFromCard,
		relationParam.CardId,
		relationParam.SourceKeyId,
		relationParam.SourceProjectId,
		relationParam.TargetCardId,
		relationParam.TargetKeyId,
		relationParam.TargetProjectId,
	)
	if err != nil {
		return err
	}
	err = unrelateCardFromCard(jwtToken, route)
	if err != nil {
		return err
	}
	return nil
}
