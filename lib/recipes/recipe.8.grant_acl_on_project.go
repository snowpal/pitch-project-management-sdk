package recipes

import (
	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/projects/projects.1"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"

	recipes "github.com/snowpal/pitch-building-projects-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"

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

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	key, err := recipes.AddProjectKey(user, CopyKeyName)
	if err != nil {
		return
	}

	log.Info("Add custom project")
	recipes.SleepBefore()
	project, err := recipes.AddProject(user, CopyProjectName, key)
	if err != nil {
		return
	}
	log.Printf(".Project %s added successfully", project.Name)
	recipes.SleepAfter()

	log.Info("Share project with read access")
	recipes.SleepBefore()
	err = recipes.SearchUserAndShareProject(user, project, "api_read_user", lib.ReadAcl)
	if err != nil {
		return
	}
	log.Printf(".Project %s shared with %s with read access level", project.Name, lib.ReadUser)
	recipes.SleepAfter()

	log.Info("Copy project and see acl is not copied")
	recipes.SleepBefore()
	anotherProject, err := copyProject(user, project)
	if err != nil {
		return
	}
	log.Printf(".Project %s copied but %s don't have access on copied project", project.Name, lib.ReadUser)
	recipes.SleepAfter()

	log.Info("Share project with admin access")
	recipes.SleepBefore()
	err = recipes.SearchUserAndShareProject(user, anotherProject, "api_admin_user", lib.AdminAcl)
	if err != nil {
		return
	}
	log.Printf(".Project %s shared with %s with admin access", project.Name, lib.ReadUser)
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
