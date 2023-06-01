package cards

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
)

func CopyCard(jwtToken string, cardParam request.CopyMoveCardParam) error {
	route, err := helpers.GetRoute(
		lib.RouteCardsCopyCard,
		cardParam.CardId,
		cardParam.KeyId,
		cardParam.ProjectId,
		strconv.FormatBool(cardParam.AllTasks),
		strconv.FormatBool(cardParam.AllChecklists),
		cardParam.TargetKeyId,
		cardParam.TargetProjectId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
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
