package templates

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetProjectTemplates(jwtToken string) ([]response.ProjectTemplate, error) {
	resProjectTemplates := response.ProjectTemplates{}
	route, err := helpers2.GetRoute(lib.RouteTemplatesGetProjectTemplates)
	if err != nil {
		fmt.Println(err)
		return resProjectTemplates.Templates, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resProjectTemplates.Templates, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectTemplates.Templates, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProjectTemplates.Templates, err
	}

	err = json.Unmarshal(body, &resProjectTemplates)
	if err != nil {
		fmt.Println(err)
		return resProjectTemplates.Templates, err
	}
	return resProjectTemplates.Templates, nil
}
