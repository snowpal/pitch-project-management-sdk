package scales

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetProjectsUsingScale(jwtToken string, scaleId string) ([]response.Project, error) {
	resProjects := response.Projects{}
	route, err := helpers2.GetRoute(lib.RouteScalesGetProjectsUsingScale, scaleId)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
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
