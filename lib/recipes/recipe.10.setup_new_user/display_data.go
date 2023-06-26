package setupnewuser

import (
	"fmt"

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
	log.Info(fmt.Sprintf("- %s | %s", email, user.JwtToken))
}

func displayAllKeys(user response.User) {
	allKeys, err := keys.GetKeys(user.JwtToken, 0)
	if err != nil {
		return
	}
	log.Info("List of Keys")
	for kIndex, key := range allKeys {
		if key.Type != lib.ProjectKeyType {
			continue
		}

		log.Info(fmt.Sprintf("%d. %s | %s", kIndex+1, key.Name, key.Type))
		allProjects, err := projects.GetProjects(user.JwtToken, request.GetProjectsParam{
			KeyId: key.ID,
		})
		if err != nil {
			return
		}

		log.Info(fmt.Sprintf("List of Prjects inside %s", key.Name))
		for bIndex, project := range allProjects {
			log.Info(fmt.Sprintf("%d. %s", bIndex+1, project.Name))

			allCards, err := cards.GetCards(user.JwtToken, request.GetCardsParam{
				KeyId:     key.ID,
				ProjectId: &project.ID,
			})
			if err != nil {
				return
			}

			log.Info(fmt.Sprintf("List of Cards inside %s and %s", project.Name, key.Name))
			for bpIndex, card := range allCards {
				log.Info(fmt.Sprintf("%d. %s", bpIndex+1, card.Name))
			}
		}
	}
}

func displayAllNotifications(user response.User) {
	allNotifications, err := notifications.GetNotifications(user.JwtToken)
	if err != nil {
		return
	}

	for index, notification := range allNotifications {
		log.Info(fmt.Sprintf("%d. %s", index+1, notification.Text))
	}
}

func DisplayData(user response.User, anotherUserEmail string) {
	log.Info("## Registered Users")
	displayUser(user.Email)
	displayUser(anotherUserEmail)

	log.Info(fmt.Sprintf("## Resources Created for user: %s", user.Email))
	displayAllKeys(user)

	anotherUser, err := recipes.SignIn(anotherUserEmail, lib.Password)
	if err != nil {
		return
	}

	log.Info(fmt.Sprintf("## Notifications for shared data as user: %s", anotherUserEmail))
	displayAllNotifications(anotherUser)
}
