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

func UpdateBlockScaleValue(
	jwtToken string,
	reqBody request.UpdateScaleValueReqBody,
	projectParam common.ResourceIdParam,
) (response.UpdateBlockScaleValue, error) {
	resBlockScaleValue := response.UpdateBlockScaleValue{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteProjectsUpdateBlockScaleValue,
		projectParam.BlockId,
		projectParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}

	err = json.Unmarshal(body, &resBlockScaleValue)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}
	return resBlockScaleValue, nil
}
