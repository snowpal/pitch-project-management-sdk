package recipes

import (
	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/relations"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"

	log "github.com/sirupsen/logrus"
	recipes "github.com/snowpal/pitch-building-projects-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"
)

const (
	RelationKeyName   = "Animals"
	RelationBlockName = "Tiger"
	RelationPodName   = "Cat"
)

func AddRelation() {
	log.Info("Objective: Add and Remove Relations")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	log.Info("Create a key and project & pod into that key")
	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Relate the project with key pod")
	project, pod, err := addRelation(user)
	if err != nil {
		return
	}
	log.Printf(".Block %s is related with pod %s successfully", project.Name, pod.Name)

	log.Info("Unrelate the project from key pod")
	err = removeRelation(user, project, pod)
	if err != nil {
		return
	}
	log.Printf(".Block %s is unrelated from pod %s successfully", project.Name, pod.Name)
}

func removeRelation(user response.User, project response.Block, pod response.Pod) error {
	err := relations.UnrelateBlockFromKeyPod(
		user.JwtToken,
		request.BlockToPodRelationParam{
			BlockId:     project.ID,
			TargetPodId: pod.ID,
			TargetKeyId: pod.Key.ID,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func addRelation(user response.User) (response.Block, response.Pod, error) {
	var (
		project response.Block
		pod     response.Pod
	)
	key, err := recipes.AddCustomKey(user, RelationKeyName)
	if err != nil {
		return project, pod, err
	}
	project, err = recipes.AddBlock(user, RelationBlockName, key)
	if err != nil {
		return project, pod, err
	}
	pod, err = recipes.AddPod(user, RelationPodName, key)
	if err != nil {
		return project, pod, err
	}
	err = relations.RelateBlockToKeyPod(
		user.JwtToken,
		request.BlockToPodRelationParam{
			BlockId:     project.ID,
			TargetPodId: pod.ID,
			TargetKeyId: pod.Key.ID,
		},
	)
	if err != nil {
		return project, pod, err
	}
	return project, pod, nil

}
