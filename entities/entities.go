package entities

type Event struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Member struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
