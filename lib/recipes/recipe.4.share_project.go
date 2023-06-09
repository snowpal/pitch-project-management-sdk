package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/collaboration/collaboration.1.projects"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/notifications"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/projects/projects.1"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	user2 "github.com/snowpal/pitch-project-management-sdk/lib/endpoints/user"
)

const (
	KeyName            = "Snowpal Pitch"
	ShareProjectName   = "Web App"
	UpdatedProjectName = "Android App"
)

func ShareProject() {
	log.Info("Objective: Create project, share users as read & write, make 1 of them as admin.")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Share a project")
	recipes.SleepBefore()
	var project response.Project
	project, err = shareProject(user)
	if err != nil {
		return
	}

	writeUser, err := getWriteUser(user, project)
	fmt.Println(user.JwtToken)
	if err != nil {
		return
	}

	log.Info("Show notifications as write user")
	recipes.SleepBefore()
	err = showNotificationsAsWriteUser(writeUser)
	if err != nil {
		return
	}
	log.Info(".Notifications for the recent share displayed successfully")
	recipes.SleepAfter()

	log.Info("Update project name as a write user")
	recipes.SleepBefore()
	var resProject response.Project
	resProject, err = updateProjectAsWriteUser(writeUser, project)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Write user updated project name to %s successfully", resProject.Name))
	recipes.SleepAfter()

	log.Info("Grant admin access to a user with read access")
	err = makeReadUserAsAdmin(user, project)
	if err != nil {
		return
	}
	log.Info(".Admin access has been granted successfully")
}

func getWriteUser(user response.User, project response.Project) (response.User, error) {
	var writeUser response.User
	resProject, err := collaboration.GetProjectCollaborators(
		user.JwtToken,
		common.ResourceIdParam{
			ProjectId: project.ID,
			KeyId:     project.Key.ID,
		})
	if err != nil {
		return writeUser, err
	}
	allUsers, err := user2.GetUsers(user.JwtToken)
	for _, sharedUser := range *resProject.SharedUsers {
		for _, userInAll := range allUsers {
			if sharedUser.ID == userInAll.ID {
				writeUser = userInAll
				break
			}
		}
	}
	if err != nil {
		return writeUser, err
	}

	writeUser, err = recipes.SignIn(writeUser.Email, lib.Password)
	if err != nil {
		return writeUser, err
	}
	return writeUser, nil
}

func shareProject(user response.User) (response.Project, error) {
	var project response.Project
	key, err := keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: KeyName,
			Type: lib.ProjectKeyType,
		})
	if err != nil {
		return project, err
	}
	project, err = projects.AddProject(
		user.JwtToken,
		request.AddProjectReqBody{Name: ShareProjectName},
		key.ID)
	if err != nil {
		return project, err
	}
	err = recipes.SearchUserAndShareProject(user, project, lib.ReadUser, lib.ReadAcl)
	if err != nil {
		return project, err
	}
	err = recipes.SearchUserAndShareProject(user, project, lib.WriteUser, lib.WriteAcl)
	if err != nil {
		return project, err
	}
	return project, nil
}

func showNotificationsAsWriteUser(writeUser response.User) error {
	unreadNotifications, err := notifications.GetNotifications(writeUser.JwtToken)
	if err != nil {
		return err
	}
	fmt.Println(len(unreadNotifications))
	for index, notification := range unreadNotifications {
		if notification.Type == "acl" {
			log.Info(fmt.Sprintf(".Notification %d: %s", index, notification.Text))
		}
	}
	return nil
}

func updateProjectAsWriteUser(writeUser response.User, project response.Project) (response.Project, error) {
	const (
		SystemKeyType       = "system"
		customSystemKeyType = "SharedCustomKey"
	)
	systemKeys, _ := keys.GetKeysFilteredByType(writeUser.JwtToken, SystemKeyType)
	var customSystemKey response.Key
	for _, systemKey := range systemKeys {
		if systemKey.Type == customSystemKeyType {
			customSystemKey = systemKey
			break
		}
	}
	updatedProjectName := UpdatedProjectName
	resProject, err := projects.UpdateProject(
		writeUser.JwtToken,
		projects.UpdateProjectReqBody{Name: &updatedProjectName},
		common.ResourceIdParam{
			ProjectId: project.ID,
			KeyId:     customSystemKey.ID,
		})
	if err != nil {
		return resProject, err
	}
	return resProject, nil
}

func makeReadUserAsAdmin(user response.User, project response.Project) error {
	resProject, err := collaboration.GetProjectCollaborators(
		user.JwtToken,
		common.ResourceIdParam{
			ProjectId: project.ID,
			KeyId:     project.Key.ID,
		})
	if err != nil {
		return err
	}

	var readUser response.SharedUser
	for _, sharedUser := range *resProject.SharedUsers {
		if sharedUser.Acl == lib.ReadAcl {
			readUser = sharedUser
			break
		}
	}

	_, err = collaboration.UpdateProjectAcl(
		user.JwtToken,
		request.ProjectAclReqBody{Acl: lib.AdminAcl},
		common.AclParam{
			UserId: readUser.ID,
			ResourceIds: common.ResourceIdParam{
				ProjectId: project.ID,
				KeyId:     project.Key.ID,
			},
		})
	if err != nil {
		return err
	}
	return nil
}
