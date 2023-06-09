package response

import (
	common2 "github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
)

type ProjectTypes struct {
	ProjectTypes []ProjectType `json:"projectTypes"`
}

type ProjectType struct {
	ID       string                 `json:"id"`
	Name     string                 `json:"projectTypeName"`
	Projects *[]common2.SlimProject `json:"projects"`

	Modifier     *common2.ResourceModifier `json:"modifier"`
	LastModified string                    `json:"lastModified"`
}
