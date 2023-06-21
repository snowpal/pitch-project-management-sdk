package cards

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func UpdateCardComment(
	jwtToken string,
	reqBody request.CommentReqBody,
	commentParam request.CommentIdParam,
) (response.Comment, error) {
	resComment := response.Comment{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resComment, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteCardsUpdateCardComment,
		*commentParam.CommentId,
		commentParam.KeyId,
		*commentParam.ProjectId,
		*commentParam.CardId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return resComment, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resComment, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resComment, err
	}

	err = json.Unmarshal(body, &resComment)
	if err != nil {
		return resComment, err
	}
	return resComment, nil
}
