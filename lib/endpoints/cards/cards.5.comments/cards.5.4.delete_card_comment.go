package cards

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func DeleteCardComment(jwtToken string, commentParam request.CommentIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteCardsDeleteCardComment,
		*commentParam.CommentId,
		commentParam.KeyId,
		*commentParam.ProjectId,
		*commentParam.CardId,
	)
	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
