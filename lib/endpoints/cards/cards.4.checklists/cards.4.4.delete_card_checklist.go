package cards

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
)

func DeleteCardChecklist(jwtToken string, checklistParam request.ChecklistIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteCardsDeleteCardChecklist,
		*checklistParam.CardId,
		*checklistParam.ChecklistId,
		checklistParam.KeyId,
		*checklistParam.ProjectId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

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
