package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/projects/projects.1"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"

	recipes "github.com/snowpal/pitch-project-management-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-project-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

const (
	CopyKeyName     = "Insurance"
	CopyProjectName = "Car Insurance"
)

func GrantAclOnProject() {
	log.Info("Objective: Add Project, Share Project, Grant Read Access, Copy Project, Grant Admin Access")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	var key response.Key
	key, err = keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: CopyKeyName,
			Type: lib.ProjectKeyType,
		})
	if err != nil {
		return
	}

	log.Info("Add project")
	recipes.SleepBefore()
	var project response.Project
	project, err = projects.AddProject(
		user.JwtToken,
		request.AddProjectReqBody{Name: CopyProjectName},
		key.ID)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Project %s added successfully", project.Name))
	recipes.SleepAfter()

	log.Info("Share project with read access")
	recipes.SleepBefore()
	err = recipes.SearchUserAndShareProject(user, project, lib.AdminUser, lib.ReadAcl)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Project %s shared with %s with read access level", project.Name, lib.ReadUser))
	recipes.SleepAfter()

	log.Info("Copy project and see acl is not copied")
	recipes.SleepBefore()
	var anotherProject response.Project
	anotherProject, err = copyProject(user, project)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Project %s copied but %s don't have access on copied project", project.Name, lib.ReadUser))
	recipes.SleepAfter()

	log.Info("Share project with admin access")
	recipes.SleepBefore()
	err = recipes.SearchUserAndShareProject(user, anotherProject, lib.AdminUser, lib.AdminAcl)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Project %s shared with %s with admin access", project.Name, lib.ReadUser))
	recipes.SleepAfter()
}

func copyProject(user response.User, project response.Project) (response.Project, error) {
	resProject, err := projects.CopyProject(
		user.JwtToken,
		request.CopyMoveProjectParam{
			ProjectId:     project.ID,
			KeyId:         project.Key.ID,
			TargetKeyId:   project.Key.ID,
			AllCards:      true,
			AllTasks:      true,
			AllChecklists: true,
		},
	)
	if err != nil {
		return resProject, err
	}
	return resProject, nil
}
