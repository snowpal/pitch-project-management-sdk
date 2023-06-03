package projects

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

type ProjectByTemplateParam struct {
	KeyId        string
	TemplateId   string
	ExcludeCards bool
	ExcludeTasks bool
}

func AddProjectBasedOnTemplate(
	jwtToken string,
	reqBody request.AddProjectReqBody,
	projectParam ProjectByTemplateParam,
) (response.Project, error) {
	resProject := response.Project{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resProject, err
	}
	payload := strings.NewReader(requestBody)
	var route string
	route, err = helpers2.GetRoute(
		lib.RouteProjectsAddProjectBasedOnTemplate,
		projectParam.KeyId,
		projectParam.TemplateId,
		strconv.FormatBool(projectParam.ExcludeCards),
		strconv.FormatBool(projectParam.ExcludeTasks),
	)
	if err != nil {
		fmt.Println(err)
		return resProject, err
	}
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resProject, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProject, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProject, err
	}

	err = json.Unmarshal(body, &resProject)
	if err != nil {
		fmt.Println(err)
		return resProject, err
	}
	return resProject, nil
}
