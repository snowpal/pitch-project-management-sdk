package projects

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
)

type AddProjectTypeIdParam struct {
	KeyId         string
	ProjectId     string
	ProjectTypeId string
}

func AddCardTypeToCard(jwtToken string, cardParam AddProjectTypeIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectsAddProjectTypeToProject,
		cardParam.ProjectId,
		cardParam.ProjectTypeId,
		cardParam.KeyId,
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
