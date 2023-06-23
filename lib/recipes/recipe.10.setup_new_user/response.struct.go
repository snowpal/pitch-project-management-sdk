package recipes

import "github.com/snowpal/pitch-project-management-sdk/lib/structs/response"

type KeyWithProjects struct {
	Key      response.Key
	Projects []ProjectWithCards
}

type ProjectWithCards struct {
	Project      response.Project
	ProjectCards []response.Card
}
