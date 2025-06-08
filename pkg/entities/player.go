package entities

type Player struct {
	ID      string   `json:"player_id"`
	Metrics []Metric `json:"metrics"`
}
