package projects

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"
)

func AllowArchivalOfProject(
	jwtToken string,
	reqBody common.AllowArchivalReqBody,
	projectParam common.ResourceIdParam,
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
		lib.RouteProjectsAllowArchivalOfProject,
		projectParam.ProjectId,
		projectParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resProject, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resProject, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
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
