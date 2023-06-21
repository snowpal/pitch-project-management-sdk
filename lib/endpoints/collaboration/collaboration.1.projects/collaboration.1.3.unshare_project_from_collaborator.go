package collaboration

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func UnshareProjectWithCollaborator(jwtToken string, projectAclParam common.AclParam) (response.Project, error) {
	resProject := response.Project{}
	route, err := helpers2.GetRoute(
		lib.RouteCollaborationUnshareProjectFromCollaborator,
		projectAclParam.ResourceIds.ProjectId,
		projectAclParam.UserId,
		projectAclParam.ResourceIds.KeyId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		return resProject, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resProject, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resProject, err
	}

	err = json.Unmarshal(body, &resProject)
	if err != nil {
		return resProject, err
	}
	return resProject, nil
}
