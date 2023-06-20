package recipes

import (
	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/cards/cards.1"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/notifications"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/projects/projects.1"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func displayUser(email string) {
	user, err := recipes.SignIn(email, lib.Password)
	if err != nil {
		return
	}
	log.Info("- ", email, " | ", user.JwtToken)
}

func displayAllKeys(user response.User) {
	keys, err := keys.GetKeys(user.JwtToken, 0)
	if err != nil {
		return
	}
	log.Info("List of Keys")
	for kIndex, key := range keys {
		if key.Type != lib.ProjectKeyType {
			continue
		}

		log.Info(kIndex+1, ". ", key.Name, " | ", key.Type)
		projects, err := projects.GetProjects(user.JwtToken, request.GetProjectsParam{
			KeyId: key.ID,
		})
		if err != nil {
			return
		}

		log.Info("List of Prjects inside ", key.Name)
		for bIndex, project := range projects {
			log.Info(bIndex+1, ". ", project.Name)

			cards, err := cards.GetCards(user.JwtToken, request.GetCardsParam{
				KeyId:     key.ID,
				ProjectId: &project.ID,
			})
			if err != nil {
				return
			}

			log.Info("List of Cards inside ", project.Name, " and ", key.Name)
			for bpIndex, card := range cards {
				log.Info(bpIndex+1, ". ", card.Name)
			}
		}
	}
}

func displayAllNotifications(user response.User) {
	notifications, err := notifications.GetNotifications(user.JwtToken)
	if err != nil {
		return
	}

	for index, notification := range notifications {
		log.Info(index+1, ". ", notification.Text)
	}
}

func DisplayContent(user response.User, anotherUserEmail string) {
	log.Info("## Registered Users")
	displayUser(user.Email)
	displayUser(anotherUserEmail)

	log.Info("## Resources Created for user: ", user.Email)
	displayAllKeys(user)

	anotherUser, err := recipes.SignIn(anotherUserEmail, lib.Password)
	if err != nil {
		return
	}

	log.Info("## Notifications for shared content as user: ", anotherUserEmail)
	displayAllNotifications(anotherUser)
}
