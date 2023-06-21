package recipes

import (
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

var KeyName = "Snowpal"
var ProjectName = "Mobile Development"
var ProjectCardName = "Flutter"

func CreateContent(user response.User) (KeyWithProjects, error) {
	var err error
	var KeyWithProjects KeyWithProjects

	log.Info("Creating project key")
	key, err := recipes.AddProjectKey(user, KeyName)
	if err != nil {
		return KeyWithProjects, err
	}

	var ProjectWithCards ProjectWithCards

	log.Info("Creating project inside ", KeyName, " key.")
	var project response.Project
	project, err = recipes.AddProject(user, ProjectName, key)
	if err != nil {
		return KeyWithProjects, err
	}
	ProjectWithCards.Project = project

	log.Info("Creating card inside ", ProjectName, " project.")
	var card response.Card
	card, err = recipes.AddProjectCard(user, ProjectCardName, project)
	if err != nil {
		return KeyWithProjects, err
	}
	ProjectWithCards.ProjectCards = append(ProjectWithCards.ProjectCards, card)

	KeyWithProjects.Projects = append(KeyWithProjects.Projects, ProjectWithCards)

	return KeyWithProjects, err
}
