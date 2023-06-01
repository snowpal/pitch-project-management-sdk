package relations

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
)

func relateProjectToCard(jwtToken string, route string) error {
	req, err := http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RelateProjectToKeyCard(jwtToken string, relationParam request.ProjectToCardRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsRelateCardToProject,
		relationParam.ProjectId,
		relationParam.TargetCardId,
		relationParam.TargetKeyId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = relateProjectToCard(jwtToken, route)
	if err != nil {
		return err
	}
	return nil
}

func RelateProjectToCard(jwtToken string, relationParam request.ProjectToCardRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsRelateCardToProject,
		relationParam.ProjectId,
		relationParam.TargetCardId,
		relationParam.TargetKeyId,
		relationParam.TargetProjectId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err != nil {
		fmt.Println(err)
		return err
	}
	err = relateProjectToCard(jwtToken, route)
	if err != nil {
		return err
	}
	return nil
}
