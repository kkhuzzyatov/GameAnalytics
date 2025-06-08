package entities

type Developer struct {
	Id           int       `json:"id"`
	Login        string    `json:"login"`
	HashPassword int       `json:"hash_of_password"`
	Projects     []Project `json:"project"`
}
