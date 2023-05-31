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
	CopyKeyName   = "Insurance"
	CopyBlockName = "Car Insurance"
)

func GrantAclOnCustomBlock() {
	log.Info("Objective: Add Custom Block, Share Block, Grant Read Access, Copy Block, Grant Admin Access")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	key, err := recipes.AddCustomKey(user, CopyKeyName)
	if err != nil {
		return
	}

	log.Info("Add custom project")
	recipes.SleepBefore()
	project, err := recipes.AddBlock(user, CopyBlockName, key)
	if err != nil {
		return
	}
	log.Printf(".Block %s added successfully", project.Name)
	recipes.SleepAfter()

	log.Info("Share project with read access")
	recipes.SleepBefore()
	err = recipes.SearchUserAndShareBlock(user, project, "api_read_user", lib.ReadAcl)
	if err != nil {
		return
	}
	log.Printf(".Block %s shared with %s with read access level", project.Name, lib.ReadUser)
	recipes.SleepAfter()

	log.Info("Copy project and see acl is not copied")
	recipes.SleepBefore()
	anotherBlock, err := copyBlock(user, project)
	if err != nil {
		return
	}
	log.Printf(".Block %s copied but %s don't have access on copied project", project.Name, lib.ReadUser)
	recipes.SleepAfter()

	log.Info("Share project with admin access")
	recipes.SleepBefore()
	err = recipes.SearchUserAndShareBlock(user, anotherBlock, "api_admin_user", lib.AdminAcl)
	if err != nil {
		return
	}
	log.Printf(".Block %s shared with %s with admin access", project.Name, lib.ReadUser)
	recipes.SleepAfter()
}

func copyBlock(user response.User, project response.Block) (response.Block, error) {
	resBlock, err := projects.CopyBlock(
		user.JwtToken,
		request.CopyMoveBlockParam{
			BlockId:       project.ID,
			KeyId:         project.Key.ID,
			TargetKeyId:   project.Key.ID,
			AllPods:       true,
			AllTasks:      true,
			AllChecklists: true,
		},
	)
	if err != nil {
		return resBlock, err
	}
	return resBlock, nil
}
