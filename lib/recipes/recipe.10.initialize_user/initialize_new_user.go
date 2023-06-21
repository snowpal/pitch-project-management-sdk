package recipes

import (
	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func InitializeNewUser() {
	log.Info("Objective: Initialize new user with a dynamic email address and create content for that user")
	userEmail, err := RegisterNewUser()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(userEmail, lib.Password)
	if err != nil {
		return
	}

	log.Info("Creating content for ", user.Email)
	var keyWithProjects KeyWithProjects
	keyWithProjects, err = CreateContent(user)
	if err != nil {
		return
	}

	log.Info("Register another user to share content")
	var anotherUserEmail string
	anotherUserEmail, err = RegisterNewUser()
	if err != nil {
		return
	}

	log.Info("Share content with ", anotherUserEmail)
	err = ShareContent(user, anotherUserEmail, keyWithProjects)
	if err != nil {
		return
	}

	DisplayContent(user, anotherUserEmail)
}