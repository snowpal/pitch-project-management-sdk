package projects

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"
)

func AddProject(jwtToken string, reqBody request.AddProjectReqBody, keyId string) (response.Project, error) {
	resProject := response.Project{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resProject, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteProjectsAddProject, keyId)
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