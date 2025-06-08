package entities

type Project struct {
	ID      string   `json:"project_id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}
