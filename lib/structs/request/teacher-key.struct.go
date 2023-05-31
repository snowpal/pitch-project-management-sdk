package request

import (
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
)

type PublishGradesReqBody struct {
	StudentUserIds string `json:"studentUserIds"`
}

type ClassroomIdParam struct {
	StudentId   string
	ResourceIds common.ResourceIdParam
}
