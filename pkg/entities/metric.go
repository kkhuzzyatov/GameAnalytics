package entities

type Metric struct {
	ID    string `json:"metric_id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}
