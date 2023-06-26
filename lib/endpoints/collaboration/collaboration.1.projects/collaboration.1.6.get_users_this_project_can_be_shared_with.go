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

func GetUsersThisProjectCanBeSharedWith(
	jwtToken string,
	projectAclParam common.SearchUsersParam,
) ([]response.SearchUser, error) {
	resUsers := response.SearchUsers{}
	route, err := helpers2.GetRoute(
		lib.RouteCollaborationGetUsersThisProjectCanBeSharedWith,
		projectAclParam.ResourceIds.ProjectId,
		projectAclParam.ResourceIds.KeyId,
		projectAclParam.SearchToken,
	)
	if err != nil {
		return resUsers.SearchUsers, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resUsers.SearchUsers, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resUsers.SearchUsers, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resUsers.SearchUsers, err
	}

	err = json.Unmarshal(body, &resUsers)
	if err != nil {
		return resUsers.SearchUsers, err
	}
	return resUsers.SearchUsers, nil
}
