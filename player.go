package vbcore

// Player defines a player with less properties than User. It's used to show up
// in the frontend to other users
type Player struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
}
