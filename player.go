package vbcore

// Player defines a player with less properties than User. It's used to show up
// in the frontend to other users
type Player struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
}

// PlayerInfo is a Player with more information
type PlayerInfo struct {
	UserID     string `json:"userID"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Picture    string `json:"picture"`
	Permission string `json:"permission"`
}
