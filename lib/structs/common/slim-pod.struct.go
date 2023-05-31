package common

type SlimCard struct {
	ID       string        `json:"id"`
	Name     string        `json:"podName"`
	Key      *SlimKey      `json:"key"`
	Project  *SlimProject  `json:"project"`
	Projects []SlimProject `json:"projects"`
}
