package recipes

import (
	"github.com/snowpal/pitch-building-projects-sdk/lib"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/request"

	keyPods "github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/key_pods/key_pods.1"
	cards "github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/project_pods/project_pods.1"
	projects "github.com/snowpal/pitch-building-projects-sdk/lib/endpoints/projects/projects.1"
	recipes "github.com/snowpal/pitch-building-projects-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

const (
	Key1Name         = "Taxes"
	AnotherKeyName   = "State Taxes"
	Block1Name       = "Form 1040"
	AnotherBlockName = "Form 1120S"
	Pod1Name         = "Income"
	BlockPod1Name    = "Expenses"
)

// AddAndLinkResources Add project, pod & project pod to a key and link them into another key
func AddAndLinkResources() {
	log.Info("Objective: Add keys and projects, and link projects")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	var newKey response.Key
	log.Info("Add a new custom key")
	recipes.SleepBefore()
	newKey, err = recipes.AddCustomKey(user, Key1Name)
	if err != nil {
		return
	}
	log.Printf(".Key, %s is added successfully.", newKey.Name)
	recipes.SleepAfter()

	var (
		newPod      response.Pod
		newBlock    response.Block
		newBlockPod response.Pod
	)
	newPod, newBlock, newBlockPod, err = addProjectsAndPods(user, newKey)
	if err != nil {
		return
	}

	log.Info("Add another key")
	recipes.SleepBefore()
	var anotherKey response.Key
	anotherKey, err = recipes.AddCustomKey(user, AnotherKeyName)
	if err != nil {
		return
	}

	log.Info("Add project")
	recipes.SleepBefore()
	var anotherBlock response.Block
	anotherBlock, err = recipes.AddBlock(user, AnotherBlockName, newKey)
	if err != nil {
		return
	}
	log.Printf(".Block, %s is created successfully.", newBlock.Name)
	recipes.SleepAfter()

	err = linkResources(user, anotherKey, anotherBlock, newBlock, newBlockPod, newPod)
	if err != nil {
		return
	}
}

func linkResources(
	user response.User,
	anotherKey response.Key,
	anotherBlock response.Block,
	newBlock response.Block,
	newBlockPod response.Pod,
	newPod response.Pod,
) error {
	log.Info("Link key pod into the other key")
	recipes.SleepBefore()
	err := keyPods.LinkPodToKey(user.JwtToken,
		common.ResourceIdParam{
			PodId: newPod.ID,
			KeyId: anotherKey.ID,
		})
	if err != nil {
		return err
	}
	log.Printf(".Block Pod, %s is linked successfully to %s Key.", newPod.Name, anotherKey.Name)
	recipes.SleepAfter()

	log.Info("Link project into the other key")
	recipes.SleepBefore()
	err = projects.LinkBlockToKey(user.JwtToken,
		common.ResourceIdParam{
			BlockId: newBlock.ID,
			KeyId:   anotherKey.ID,
		})
	if err != nil {
		return err
	}
	log.Printf(".Block, %s is linked successfully to %s Key.", newBlock.Name, anotherKey.Name)
	recipes.SleepAfter()

	log.Info("Link key pod into the other project")
	recipes.SleepBefore()
	err = cards.LinkPodToBlock(user.JwtToken,
		common.ResourceIdParam{
			PodId:   newBlockPod.ID,
			BlockId: anotherBlock.ID,
			KeyId:   anotherKey.ID,
		})
	if err != nil {
		return err
	}
	log.Printf(".Pod, %s is linked successfully to %s Block.", newPod.Name, anotherBlock.Name)
	recipes.SleepAfter()
	return nil
}

func addProjectsAndPods(user response.User, newKey response.Key) (response.Pod, response.Block, response.Pod, error) {
	log.Info("Add a new key pod into this key")
	recipes.SleepBefore()

	var (
		pod     response.Pod
		project response.Block
	)
	newPod, err := keyPods.AddKeyPod(user.JwtToken,
		request.AddPodReqBody{
			Name: Pod1Name,
		},
		newKey.ID)
	if err != nil {
		return pod, project, pod, err
	}
	log.Printf(".Key Pod, %s is created successfully in %s Key.", newPod.Name, newKey.Name)
	recipes.SleepAfter()

	log.Info("Add a new project")
	recipes.SleepBefore()
	var newBlock response.Block
	newBlock, err = recipes.AddBlock(user, Block1Name, newKey)
	if err != nil {
		return pod, project, pod, err
	}
	log.Printf(".Block, %s is created successfully.", newBlock.Name)
	recipes.SleepAfter()

	log.Info("Add a new project pod in this project")
	recipes.SleepBefore()
	var newBlockPod response.Pod
	newBlockPod, err = cards.AddBlockPod(user.JwtToken,
		request.AddPodReqBody{
			Name: BlockPod1Name,
		},
		common.ResourceIdParam{
			BlockId: newBlock.ID,
			KeyId:   newKey.ID,
		})
	if err != nil {
		return pod, project, pod, err
	}
	log.Printf(".Block Pod, %s is created successfully in %s Block.", newBlockPod.Name, newBlock.Name)
	recipes.SleepAfter()
	return newPod, newBlock, newBlockPod, nil
}
