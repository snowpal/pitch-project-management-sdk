package response

import (
	common2 "github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
)

type SearchResources struct {
	Results []SearchResource `json:"results"`
}

type SearchResource struct {
	ID   string `json:"id"`
	Type string `json:"type"`

	// Relation attribute
	IsRelated bool `json:"isRelated"`

	KeyName     *string `json:"keyName"`
	KeyType     *string `json:"keyType"`
	ProjectName *string `json:"projectName"`
	CardName    *string `json:"podName"`

	Key      *common2.SlimKey       `json:"key"`
	Project  *common2.SlimProject   `json:"project"`
	Projects *[]common2.SlimProject `json:"projects"`

	Modifier common2.ResourceModifier `json:"modifier"`
}

type Relations struct {
	Relationships Relationships `json:"relationships"`
}

type Relationships struct {
	Keys     []common2.SlimKey     `json:"keys"`
	Projects []common2.SlimProject `json:"projects"`
	Cards    []common2.SlimCard    `json:"pods"`
}
