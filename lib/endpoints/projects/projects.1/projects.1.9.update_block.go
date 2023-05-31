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

type UpdateProjectReqBody struct {
	Name              *string `json:"projectName"`
	ProjectId         *string `json:"projectId"`
	SimpleDescription *string `json:"simpleDescription"`
	DueDate           *string `json:"projectDueDate"`
	StartTime         *string `json:"projectStartTime"`
	EndTime           *string `json:"projectEndTime"`
	Color             *string `json:"projectColor"`
	Tags              *string `json:"projectTags"`
	KanbanMode        *bool   `json:"kanbanMode"`
	Completed         bool    `json:"projectCompleted"`
}

func UpdateProject(jwtToken string, reqBody UpdateProjectReqBody, projectParam common.ResourceIdParam) (response.Project, error) {
	resProject := response.Project{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resProject, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteProjectsUpdateProject, projectParam.ProjectId, projectParam.KeyId)
	if err != nil {
		fmt.Println(err)
		return resProject, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
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
