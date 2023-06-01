package response

type Profile struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Username      string `json:"userName"`
	FirstName     string `json:"firstName"`
	MiddleName    string `json:"middleName"`
	LastName      string `json:"lastName"`
	PhoneNumber   string `json:"phoneNumber"`
	AddressUserBy string `json:"addressUserBy"`
	UserInitial   string `json:"userInitial"`
	AvatarName    string `json:"avatarName"`
	AvatarUrl     string `json:"avatarUrl"`

	IsInactive         bool `json:"inactive"`
	IsAnnualSubscriber bool `json:"isAnnualSubscriber"`

	ProjectedUsers []ProjectedUser `json:"projectedUsers"`
	Preferences    []Preference    `json:"preferences"`
}

type ProjectedUser struct {
	ID        string `json:"id"`
	Username  string `json:"userName"`
	FirstName string `json:"firstName"`
}

type Preference struct {
	SubscribeMail              bool `json:"subscribeMail"`
	SubscribeSms               bool `json:"subscribeSms"`
	ShowResourcesChart         bool `json:"showResourcesChart"`
	ShowScalesChartForProjects bool `json:"showScalesChartForProjects"`
	ShowScalesChartForCards    bool `json:"showScalesChartForCards"`
}
