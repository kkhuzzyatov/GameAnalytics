package entities

type Developer struct {
	ID             string    `json:"developer_id"`
	Login          string    `json:"login"`
	HashedPassword int       `json:"hashed_password"`
	Projects       []Project `json:"project"`
}
