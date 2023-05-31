package projects

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
)

type AddProjectTypeIdParam struct {
	KeyId         string
	ProjectId     string
	ProjectTypeId string
}

func AddCardTypeToCard(jwtToken string, podParam AddProjectTypeIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectsAddProjectTypeToProject,
		podParam.ProjectId,
		podParam.ProjectTypeId,
		podParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
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
