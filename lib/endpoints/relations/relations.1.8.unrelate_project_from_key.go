package relations

import (
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func UnrelateKeyFromProject(jwtToken string, relationParam request.KeyToProjectRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelateProjectFromKey,
		relationParam.KeyId,
		relationParam.TargetProjectId,
	)
	if err != nil {
		return err
	}

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
