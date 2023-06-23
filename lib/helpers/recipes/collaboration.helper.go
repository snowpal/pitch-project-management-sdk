package recipes

import (
	"strings"

	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/collaboration/collaboration.1.projects"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func SearchUserAndShareProject(user response.User, project response.Project, searchToken string, acl string) error {
	projectIdParam := common.ResourceIdParam{
		ProjectId: project.ID,
		KeyId:     project.Key.ID,
	}

	searchedUsers, err := collaboration.GetUsersThisProjectCanBeSharedWith(
		user.JwtToken,
		common.SearchUsersParam{
			SearchToken: strings.Split(searchToken, "@")[0],
			ResourceIds: projectIdParam,
		})
	if err != nil {
		return err
	}

	// For the purpose of this recipe, it does not matter which user from the list we end up picking, hence we go with
	// the first one.
	_, err = collaboration.ShareProjectWithCollaborator(
		user.JwtToken,
		request.ProjectAclReqBody{Acl: acl},
		common.AclParam{
			UserId:      searchedUsers[0].ID,
			ResourceIds: projectIdParam,
		})
	if err != nil {
		return err
	}
	return nil
}
