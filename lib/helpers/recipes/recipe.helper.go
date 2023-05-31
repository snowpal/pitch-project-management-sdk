package recipes

import (
	"fmt"
	"github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/collaboration/collaboration.1.projects"
	"time"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/projects/projects.1"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"
)

func sleepWindow(sleepTime time.Duration) {
	time.Sleep(time.Second * sleepTime)
}

func SleepBefore() {
	sleepWindow(1)
}

func SleepAfter() {
	sleepWindow(2)
}

// ValidateDependencies We require that the first recipe be run before anything else as it registers a bunch of users.
// To verify if it was actually run, we do this "random" check.
func ValidateDependencies() (response.User, error) {
	user, err := SignIn(lib.DefaultEmail, lib.Password)
	fmt.Println(user)
	fmt.Println(err)
	if err != nil {
		return user, err
	}

	return user, nil
}

func addKey(user response.User, keyName string, keyType string) (response.Key, error) {
	newKey, err := keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: keyName,
			Type: keyType,
		})
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddCustomKey(user response.User, keyName string) (response.Key, error) {
	newKey, err := addKey(user, keyName, lib.CustomKeyType)
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddTeacherKey(user response.User, keyName string) (response.Key, error) {
	newKey, err := addKey(user, keyName, lib.TeacherKeyType)
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddProjectKey(user response.User, keyName string) (response.Key, error) {
	newKey, err := addKey(user, keyName, lib.ProjectKeyType)
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddProject(user response.User, projectName string, key response.Key) (response.Project, error) {
	newProject, err := projects.AddProject(
		user.JwtToken,
		request.AddProjectReqBody{Name: projectName},
		key.ID)
	if err != nil {
		return newProject, err
	}
	return newProject, nil
}

func SearchUserAndShareProject(user response.User, project response.Project, searchToken string, acl string) error {
	projectIdParam := common.ResourceIdParam{
		ProjectId: project.ID,
		KeyId:     project.Key.ID,
	}

	searchedUsers, err := collaboration.GetUsersThisProjectCanBeSharedWith(
		user.JwtToken,
		common.SearchUsersParam{
			SearchToken: searchToken,
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
