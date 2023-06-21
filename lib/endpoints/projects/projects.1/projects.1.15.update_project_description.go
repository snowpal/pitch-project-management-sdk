package projects

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

type UpdateProjectDescReqBody struct {
	Description   string  `json:"projectDescription"`
	TaggedUserIds *string `json:"taggedUserIds"`
}

func UpdateCardDescription(
	jwtToken string,
	reqBody UpdateProjectDescReqBody,
	cardParam common.ResourceIdParam,
) (response.Project, error) {
	resProject := response.Project{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resProject, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers2.GetRoute(
		lib.RouteProjectsUpdateProjectDescription,
		cardParam.ProjectId,
		cardParam.KeyId,
	)
	if err != nil {
		return resProject, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return resProject, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resProject, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resProject, err
	}

	err = json.Unmarshal(body, &resProject)
	if err != nil {
		return resProject, err
	}
	return resProject, nil
}
