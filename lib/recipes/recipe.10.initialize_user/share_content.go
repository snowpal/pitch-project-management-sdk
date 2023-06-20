package recipes

import (
	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func ShareContent(user response.User, anotherUserEmail string, keyWithProjects KeyWithProjects) error {
	log.Info("Share key content with ", anotherUserEmail)

	log.Info("Share project ", keyWithProjects.Projects[0].Project.Name, " for project key, ", keyWithProjects.Key.Name)
	err := recipes.SearchUserAndShareProject(user, keyWithProjects.Projects[0].Project, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}

	return nil
}
