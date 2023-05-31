package recipes

import (
	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	projectKeys "github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/project_keys/project_keys.1"
	projectLists "github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/project_keys/project_keys.2.lists"
)

const (
	ProjectKeyName     = "Go Development"
	ProjectProjectName = "Status API"
	CardName           = "Define Endpoints"
	ProjectList1Name   = "Statuses"
	ProjectList2Name   = "Teams"
)

func AddProjectList() {
	log.Info("Objective: Add Project Lists, Project Cards, Move Card between Lists")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	projectKey, err := recipes.AddProjectKey(user, ProjectKeyName)
	if err != nil {
		return
	}

	projectProject, err := recipes.AddProject(user, ProjectProjectName, projectKey)
	if err != nil {
		return
	}

	log.Info("Add 2 project lists")
	recipes.SleepBefore()
	projectList1, err := addProjectList(user, ProjectList1Name, projectProject)
	if err != nil {
		return
	}
	projectList2, err := addProjectList(user, ProjectList2Name, projectProject)
	if err != nil {
		return
	}
	log.Printf(".Both project lists, %s and %s created successfully", projectList1.Name, projectList2.Name)
	recipes.SleepAfter()

	log.Info("Add a project pod into a project list")
	recipes.SleepBefore()
	projectCard, err := addCard(user, CardName, projectList1)
	if err != nil {
		return
	}
	log.Printf(".Project pod %s created inside %s successfully", projectCard.Name, projectList1.Name)
	recipes.SleepAfter()

	log.Printf("Move project pod %s between project lists", projectCard.Name)
	recipes.SleepBefore()
	err = moveCardBetweenLists(user, projectList1, projectList2, projectCard)
	if err != nil {
		return
	}
	log.Printf(".Project pod %s moved from list %s to list %s successfully", projectCard.Name,
		projectList1.Name, projectList2.Name)
	recipes.SleepAfter()
}

func addCard(user response.User, podName string, projectList response.ProjectList) (response.Card, error) {
	newCard, err := projectKeys.AddCard(
		user.JwtToken,
		request.AddCardReqBody{Name: podName},
		request.ProjectListIdParam{
			ProjectListId: projectList.ID,
			ProjectId:     projectList.Project.ID,
			KeyId:         projectList.Key.ID,
		},
	)
	if err != nil {
		return newCard, err
	}
	return newCard, nil
}

func addProjectList(user response.User, projectListName string, project response.Project) (response.ProjectList, error) {
	newProjectList, err := projectLists.AddProjectProjectList(
		user.JwtToken,
		request.AddProjectListReqBody{Name: projectListName},
		common.ResourceIdParam{ProjectId: project.ID, KeyId: project.Key.ID},
	)
	if err != nil {
		return newProjectList, err
	}
	return newProjectList, nil
}

func moveCardBetweenLists(
	user response.User,
	list1 response.ProjectList,
	list2 response.ProjectList,
	pod response.Card,
) error {
	err := projectLists.BulkMoveCardsInProjectList(user.JwtToken, request.CopyMoveProjectListCardsParam{
		ProjectListId:       list1.ID,
		ProjectId:           pod.Project.ID,
		KeyId:               pod.Key.ID,
		TargetProjectListId: list2.ID,
		TargetProjectId:     list2.Project.ID,
		TargetKeyId:         list2.Key.ID,
		CardIds:             []string{pod.ID},
	})
	if err != nil {
		return err
	}
	return nil
}
