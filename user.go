package gameAnalytics


// User represents a user in the game analytics system.
// means developer's user, not a game user.
type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	Username string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}
