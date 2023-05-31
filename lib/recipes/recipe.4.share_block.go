package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/collaboration/collaboration.1.projects"
	"github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/notifications"
	"github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/projects/projects.1"
	"github.com/snowpal/pitch-building-projects-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	user2 "github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/user"
)

const (
	KeyName            = "Diwali Festival"
	ProjectName        = "Diwali Function"
	UpdatedProjectName = "Diwali Celebration"
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
	log.Printf(".Notifications for the recent share displayed successfully")
	recipes.SleepAfter()

	log.Printf("Update project name as a write user")
	recipes.SleepBefore()
	var resProject response.Project
	resProject, err = updateProjectAsWriteUser(writeUser, project)
	if err != nil {
		return
	}
	log.Printf(".Write user updated project name to %s successfully", resProject.Name)
	recipes.SleepAfter()

	log.Printf("Grant admin access to a user with read access")
	err = makeReadUserAsAdmin(user, project)
	if err != nil {
		return
	}
	log.Printf(".Admin access has been granted successfully")
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
	key, err := recipes.AddProjectKey(user, KeyName)
	if err != nil {
		return project, err
	}
	project, err = recipes.AddProject(user, ProjectName, key)
	if err != nil {
		return project, err
	}
	err = recipes.SearchUserAndShareProject(user, project, "api_read_user", lib.ReadAcl)
	if err != nil {
		return project, err
	}
	err = recipes.SearchUserAndShareProject(user, project, "api_write_user", lib.WriteAcl)
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
			log.Printf(".Notification %d: %s", index, notification.Text)
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
