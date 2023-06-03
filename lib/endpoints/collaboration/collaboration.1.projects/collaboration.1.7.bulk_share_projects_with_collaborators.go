package collaboration

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
)

type ProjectBulkShareReqBody struct {
	Acl        string `json:"projectAcl"`
	ProjectIds string `json:"projectIds"`
}

func ShareProjectsWithCollaborators(
	jwtToken string,
	reqBody ProjectBulkShareReqBody,
	projectAclParam common.AclParam,
) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		lib.RouteCollaborationBulkShareProjectsWithCollaborators,
		projectAclParam.UserId,
		projectAclParam.ResourceIds.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
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
