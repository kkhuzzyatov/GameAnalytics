package entities

type Metric struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}
