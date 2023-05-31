package recipes

import (
	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/attributes"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"

	log "github.com/sirupsen/logrus"
	recipes "github.com/snowpal/pitch-building-projects-sdk/lib/helpers/recipes"
)

const (
	AttrsKeyName     = "Birds"
	AttrsProjectName = "Parrot"
)

// UpdateAttributes sign in, update key attributes, update project attributes, update pod attributes, update project pod attributes,
// get resource attributes
func UpdateAttributes() {
	log.Info("Objective: Update show/hide of key, project, pod & project pod attributes")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
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
	key, err := recipes.AddCustomKey(user, AttrsKeyName)
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
	log.Printf(".Attributes for Key %s updated successfully", key.Name)
	recipes.SleepAfter()

	log.Info("Update project attributes")
	recipes.SleepBefore()
	project, err := recipes.AddProject(user, AttrsProjectName, key)
	if err != nil {
		return
	}
	err = attributes.UpdateProjectAttrs(
		user.JwtToken,
		common.ResourceIdParam{
			ProjectId: project.ID,
			KeyId:     project.Key.ID,
		},
		request.ResourceAttributeReqBody{
			AttributeNames: "tags,rendering_mode",
			ShowAttribute:  false,
		},
	)
	if err != nil {
		return
	}
	log.Printf(".Attributes for project %s updated successfully", key.Name)
	recipes.SleepAfter()
}
