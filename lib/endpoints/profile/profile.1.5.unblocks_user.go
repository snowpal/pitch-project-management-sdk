package profiles

import (
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
)

func UnprojectsUser(jwtToken string, userId string) error {
	route, err := helpers.GetRoute(lib.RouteProfileUnblockUser, userId)
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
