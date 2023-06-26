package projectKeys

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func AddProjectProjectList(
	jwtToken string,
	reqBody request.AddProjectListReqBody,
	projectListParam common.ResourceIdParam,
) (response.ProjectList, error) {
	resProjectList := response.ProjectList{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resProjectList, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers2.GetRoute(
		lib.RouteProjectKeysAddProjectList,
		projectListParam.ProjectId,
		projectListParam.KeyId,
	)
	if err != nil {
		return resProjectList, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		return resProjectList, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resProjectList, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resProjectList, err
	}

	err = json.Unmarshal(body, &resProjectList)
	if err != nil {
		return resProjectList, err
	}
	return resProjectList, nil
}
