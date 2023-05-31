package response

// Note(anish,3,03/06/2023) Most of the attributes are common to these templates, so the gateway should send a nested
// structure ideally.

type KeyTemplates struct {
	Templates []KeyTemplatesWithType `json:"templates"`
}

type KeyTemplatesWithType struct {
	Type      string        `json:"type"`
	Templates []KeyTemplate `json:"templates"`
}

type KeyTemplate struct {
	ID           string `json:"id"`
	Name         string `json:"templateName"`
	Description  string `json:"templateDescription"`
	Color        string `json:"color"`
	Tags         string `json:"tags"`
	PreviewUrl   string `json:"previewUrl"`
	Archived     bool   `json:"archived"`
	LastModified string `json:"lastModified"`
}

type ProjectTemplates struct {
	Templates []ProjectTemplate `json:"templates"`
}

type ProjectTemplate struct {
	ID            string  `json:"id"`
	Name          string  `json:"templateName"`
	Description   string  `json:"templateDescription"`
	Color         string  `json:"color"`
	Tags          string  `json:"tags"`
	PreviewUrl    string  `json:"previewUrl"`
	Archived      bool    `json:"archived"`
	KeyTemplateId *string `json:"keyTemplateId"`
	LastModified  string  `json:"lastModified"`
}

type CardTemplates struct {
	Templates []CardTemplate `json:"templates"`
}

type CardTemplate struct {
	ID                string  `json:"id"`
	Name              string  `json:"templateName"`
	Description       string  `json:"templateDescription"`
	Color             string  `json:"color"`
	Tags              string  `json:"tags"`
	PreviewUrl        string  `json:"previewUrl"`
	Archived          bool    `json:"archived"`
	KeyTemplateId     *string `json:"keyTemplateId"`
	ProjectTemplateId *string `json:"projectTemplateId"`
	LastModified      string  `json:"lastModified"`
}
