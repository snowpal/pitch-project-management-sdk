package projects

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
)

type BulkArchiveProjectsReqBody struct {
	ProjectIds string `json:"projectIds"`
}

func BulkArchiveProjects(jwtToken string, reqBody BulkArchiveProjectsReqBody, keyId string) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers.GetRoute(lib.RouteProjectsBulkArchiveProjects, keyId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
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
