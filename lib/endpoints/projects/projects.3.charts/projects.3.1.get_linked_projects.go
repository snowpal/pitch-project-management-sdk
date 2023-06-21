package projects

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetLinkedCards(jwtToken string, projectParam common.ResourceIdParam) (response.LinkedResources, error) {
	resLinkedCards := response.LinkedResources{}
	route, err := helpers2.GetRoute(
		lib.RouteProjectsGetLinkedCards,
		projectParam.KeyId,
		projectParam.ProjectId,
	)
	if err != nil {
		return resLinkedCards, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resLinkedCards, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resLinkedCards, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resLinkedCards, err
	}

	err = json.Unmarshal(body, &resLinkedCards)
	if err != nil {
		return resLinkedCards, err
	}
	return resLinkedCards, nil
}
