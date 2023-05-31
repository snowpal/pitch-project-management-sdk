package cards

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

func UpdateCardScaleValue(
	jwtToken string,
	reqBody request.UpdateScaleValueReqBody,
	podParam common.ResourceIdParam,
) (response.UpdateCardScaleValue, error) {
	resCardScaleValue := response.UpdateCardScaleValue{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resCardScaleValue, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteCardsUpdateCardScaleValue,
		podParam.CardId,
		podParam.KeyId,
		podParam.ProjectId,
	)
	if err != nil {
		fmt.Println(err)
		return resCardScaleValue, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resCardScaleValue, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCardScaleValue, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCardScaleValue, err
	}

	err = json.Unmarshal(body, &resCardScaleValue)
	if err != nil {
		fmt.Println(err)
		return resCardScaleValue, err
	}
	return resCardScaleValue, nil
}
