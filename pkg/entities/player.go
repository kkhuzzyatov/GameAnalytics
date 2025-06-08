package entities

type Player struct {
	Id      int      `json:"id"`
	Metrics []Metric `json:"metrics"`
}
