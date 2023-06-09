package response

import (
	common2 "github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
)

type Notifications struct {
	//TODO(Anish,3,03/23/23): Rename `badges` to `notifications` first on gateway post processor and then here.
	Notifications []Notification `json:"badges"`
}

type Notification struct {
	ID       string               `json:"id"`
	Type     string               `json:"badgeType"` // See above
	Text     string               `json:"badgeText"` // See above
	Unread   bool                 `json:"unread"`
	Resource NotificationResource `json:"resource"`
}

type NotificationResource struct {
	ID           string `json:"id"`
	ResourceType string `json:"resourceType"`
	IsTask       *bool  `json:"isTask"`

	KeyName     *string `json:"keyName"`
	ProjectName *string `json:"projectName"`
	CardName    *string `json:"cardName"`

	Key      *common2.SlimKey       `json:"key"`
	Projects *[]common2.SlimProject `json:"projects"`

	StartTime *string `json:"eventStartTime"`
	EndTime   *string `json:"eventEndTime"`
}

type Performer struct {
	Username  string `json:"userName"`
	FirstName string `json:"firstName"`
}

type NotificationEvent struct {
	PerformedOn string     `json:"performedOn"`
	PerformedBy *Performer `json:"performedBy"`
}
