package entities

type player struct {
	id 			int 			`json:"id"`
	metrics []metric 	`json:"metrics"`
}