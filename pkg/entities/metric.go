package entities

type metric struct {
	id    int    `json:"id"`
	name  string `json:"name"`
	value int    `json:"value"`
}
