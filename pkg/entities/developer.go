package entities

type developer struct {
	id 								int 			`json:"id"`
	login 						string 		`json:"login"`
	hash_of_password 	int 			`json:"hash_of_password"`
	projects					[]project	`json:"project"`
} 