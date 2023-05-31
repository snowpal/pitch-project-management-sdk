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
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"
)

func UpdateProjectScaleValue(
	jwtToken string,
	reqBody request.UpdateScaleValueReqBody,
	projectParam common.ResourceIdParam,
) (response.UpdateProjectScaleValue, error) {
	resProjectScaleValue := response.UpdateProjectScaleValue{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resProjectScaleValue, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteProjectsUpdateProjectScaleValue,
		projectParam.ProjectId,
		projectParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resProjectScaleValue, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resProjectScaleValue, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectScaleValue, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProjectScaleValue, err
	}

	err = json.Unmarshal(body, &resProjectScaleValue)
	if err != nil {
		fmt.Println(err)
		return resProjectScaleValue, err
	}
	return resProjectScaleValue, nil
}