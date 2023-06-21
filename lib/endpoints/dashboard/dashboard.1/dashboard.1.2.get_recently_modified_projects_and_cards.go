package dashboard

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetRecentlyModifiedProjectsAndCards(jwtToken string) (response.Dashboard, error) {
	resDashboard := response.Dashboard{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetRecentlyModifiedProjectsAndCards)
	if err != nil {
		return resDashboard, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resDashboard, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resDashboard, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resDashboard, err
	}

	err = json.Unmarshal(body, &resDashboard)
	if err != nil {
		return resDashboard, err
	}
	return resDashboard, nil
}
