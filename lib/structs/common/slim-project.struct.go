package common

type SlimProject struct {
	ID   string   `json:"id"`
	Name string   `json:"projectName"`
	Key  *SlimKey `json:"key"`
}
