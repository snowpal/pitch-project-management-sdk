package relations

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
)

func unrelateProjectToCard(jwtToken string, route string) error {
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

func UnrelateProjectFromKeyCard(jwtToken string, relationParam request.ProjectToCardRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelateCardFromProject,
		relationParam.ProjectId,
		relationParam.TargetCardId,
		relationParam.TargetKeyId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = unrelateProjectToCard(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UnrelateProjectFromCard(jwtToken string, relationParam request.ProjectToCardRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelateCardFromProject,
		relationParam.ProjectId,
		relationParam.TargetCardId,
		relationParam.TargetKeyId,
		relationParam.TargetProjectId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = unrelateProjectToCard(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
