package projectKeys

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
)

func CopyProjectBlock(jwtToken string, projectParam request.CopyMoveBlockParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysCopyProjectBlock,
		projectParam.BlockId,
		projectParam.KeyId,
		projectParam.TargetKeyId,
		strconv.FormatBool(projectParam.AllPods),
		strconv.FormatBool(projectParam.AllTasks),
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
