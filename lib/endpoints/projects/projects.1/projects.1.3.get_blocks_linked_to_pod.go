package projects

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"
)

func GetProjectsLinkedToCards(jwtToken string, projectParam common.ResourceIdParam) ([]response.Project, error) {
	resProjects := response.Projects{}
	route, err := helpers2.GetRoute(
		lib.RouteProjectsGetProjectsLinkedToCard,
		projectParam.CardId,
		projectParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}

	req, _ := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}

	helpers2.AddUserHeaders(jwtToken, req)
	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}

	defer helpers2.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}

	err = json.Unmarshal(body, &resProjects)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}
	return resProjects.Projects, nil
}