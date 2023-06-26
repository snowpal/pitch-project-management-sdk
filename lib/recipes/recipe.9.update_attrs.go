package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/attributes"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-project-management-sdk/lib/endpoints/projects/projects.1"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	recipes "github.com/snowpal/pitch-project-management-sdk/lib/helpers/recipes"
)

const (
	AttrsKeyName     = "Birds"
	AttrsProjectName = "Parrot"
)

// UpdateAttributes sign in, update key attributes, update project attributes, update card attributes,
// get resource attributes
func UpdateAttributes() {
	log.Info("Objective: Update show/hide of key, project & card attributes")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Get displayable attributes")
	recipes.SleepBefore()
	resourceAttrs, _ := attributes.GetResourceAttrs(user.JwtToken)
	if err != nil {
		return
	}
	log.Info(resourceAttrs)

	log.Info("Update key attributes")
	recipes.SleepBefore()
	var key response.Key
	key, err = keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: AttrsKeyName,
			Type: lib.ProjectKeyType,
		})
	if err != nil {
		return
	}
	err = attributes.UpdateKeyAttrs(
		user.JwtToken,
		key.ID,
		request.ResourceAttributeReqBody{
			AttributeNames: "tags,rendering_mode",
			ShowAttribute:  false,
		},
	)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Attributes for Key %s updated successfully", key.Name))
	recipes.SleepAfter()

	log.Info("Update project attributes")
	recipes.SleepBefore()
	var project response.Project
	project, err = projects.AddProject(
		user.JwtToken,
		request.AddProjectReqBody{Name: AttrsProjectName},
		key.ID)
	if err != nil {
		return
	}
	err = attributes.UpdateProjectAttrs(
		user.JwtToken,
		common.ResourceIdParam{
			ProjectId: project.ID,
			KeyId:     key.ID,
		},
		request.ResourceAttributeReqBody{
			AttributeNames: "tags,rendering_mode",
			ShowAttribute:  false,
		},
	)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Attributes for project %s updated successfully", key.Name))
	recipes.SleepAfter()
}
