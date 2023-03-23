package models

type JakeModel struct {
	Categories []interface{} `json:"categories"`
	CreatedAt  string        `json:"created_at"`
	IconUrl    string        `json:"icon_url"`
	Id         string        `json:"id"`
	UpdatedAt  string        `json:"updated_at"`
	Url        string        `json:"url"`
	Value      string        `json:"value"`
}

type ResponseJake struct {
	Id    string `json:"id"`
	Url   string `json:"url"`
	Value string `json:"value"`
}
