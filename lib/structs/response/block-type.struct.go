package response

import (
	common2 "github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
)

type BlockTypes struct {
	BlockTypes []BlockType `json:"projectTypes"`
}

type BlockType struct {
	ID       string               `json:"id"`
	Name     string               `json:"projectTypeName"`
	Projects *[]common2.SlimBlock `json:"projects"`

	TeacherReadOnly *bool `json:"teacherReadOnly"`

	Modifier     *common2.ResourceModifier `json:"modifier"`
	LastModified string                    `json:"lastModified"`
}
