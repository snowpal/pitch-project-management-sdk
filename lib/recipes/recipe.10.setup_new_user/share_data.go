package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func ShareData(user response.User, anotherUserEmail string, keyWithProjects KeyWithProjects) error {
	log.Info(fmt.Sprintf("Share project with %s", anotherUserEmail))

	project := keyWithProjects.Projects[0].Project
	log.Info(fmt.Sprintf("Share project %s  for project key, %s", project.Name, keyWithProjects.Key.Name))
	err := recipes.SearchUserAndShareProject(user, project, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}

	return nil
}
