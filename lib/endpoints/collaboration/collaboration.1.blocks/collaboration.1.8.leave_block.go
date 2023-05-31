package collaboration

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
)

func LeaveBlock(jwtToken string, projectParam common.ResourceIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteCollaborationLeaveBlock,
		projectParam.BlockId,
		projectParam.KeyId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, nil)
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
