package cards

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
)

func DeleteBlockPodChecklist(jwtToken string, checklistParam request.ChecklistIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteCardsDeleteBlockPodChecklist,
		*checklistParam.PodId,
		*checklistParam.ChecklistId,
		checklistParam.KeyId,
		*checklistParam.BlockId,
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
