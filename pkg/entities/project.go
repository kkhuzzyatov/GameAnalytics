package entities

type project struct {
	id 			int 			`json:"id"`
	name 		string		`json:"name"`
	players []player 	`json:"players"`
}