package projectKeys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

type ReorderCardsReqBody struct {
	ProjectListId       string `json:"sourceProjectListId"`
	CardIds             string `json:"sourceProjectListCardIds"`
	TargetProjectListId string `json:"targetProjectListId"`
	TargetCardIds       string `json:"targetProjectListCardIds"`
}

func ReorderCards(
	jwtToken string,
	reqBody ReorderCardsReqBody,
	cardParam common.ResourceIdParam,
) ([]response.Card, error) {
	resCards := response.Cards{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resCards.Cards, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteProjectKeysReorderCards,
		cardParam.ProjectId,
		cardParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resCards.Cards, err
	}
	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resCards.Cards, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCards.Cards, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCards.Cards, err
	}

	err = json.Unmarshal(body, &resCards)
	if err != nil {
		fmt.Println(err)
		return resCards.Cards, err
	}
	return resCards.Cards, nil
}
