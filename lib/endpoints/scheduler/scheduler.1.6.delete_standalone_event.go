package scheduler

import (
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
)

func DeleteStandaloneEvent(jwtToken string, standaloneEventId string) error {
	route, err := helpers.GetRoute(lib.RouteSchedulerDeleteStandaloneEvent, standaloneEventId)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodDelete, route, nil)
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
