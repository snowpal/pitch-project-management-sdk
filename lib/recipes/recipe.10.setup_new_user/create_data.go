package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/cards/cards.1"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/projects/projects.1"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

var KeyName = "Snowpal"
var ProjectName = "Mobile Development"
var ProjectCardName = "Flutter"

func CreateData(user response.User) (KeyWithProjects, error) {
	var err error
	var KeyWithProjects KeyWithProjects

	log.Info("Creating project key")
	var key response.Key
	key, err = keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: KeyName,
			Type: lib.ProjectKeyType,
		})
	if err != nil {
		return KeyWithProjects, err
	}

	var ProjectWithCards ProjectWithCards

	log.Info(fmt.Sprintf("Creating project inside %s key.", KeyName))
	var project response.Project
	project, err = projects.AddProject(
		user.JwtToken,
		request.AddProjectReqBody{Name: ProjectName},
		key.ID)
	if err != nil {
		return KeyWithProjects, err
	}
	ProjectWithCards.Project = project

	log.Info(fmt.Sprintf("Creating card inside %s project.", ProjectName))
	var card response.Card
	card, err = cards.AddCard(
		user.JwtToken,
		request.AddCardReqBody{Name: ProjectCardName},
		common.ResourceIdParam{ProjectId: project.ID, KeyId: key.ID},
	)
	if err != nil {
		return KeyWithProjects, err
	}
	ProjectWithCards.ProjectCards = append(ProjectWithCards.ProjectCards, card)

	KeyWithProjects.Projects = append(KeyWithProjects.Projects, ProjectWithCards)

	return KeyWithProjects, err
}
