package projectKeys

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
)

func BulkCopyCardsInProjectList(jwtToken string, projectListParam request.CopyMoveProjectListCardsParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysBulkCopyCardsInProjectList,
		projectListParam.ProjectListId,
		projectListParam.KeyId,
		projectListParam.ProjectId,
		projectListParam.TargetKeyId,
		projectListParam.TargetProjectId,
		projectListParam.TargetProjectListId,
		strconv.FormatBool(projectListParam.AllTasks),
		strings.Join(projectListParam.CardIds, ","),
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
