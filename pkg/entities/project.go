package entities

type Project struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}
