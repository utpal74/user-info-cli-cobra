package model

// User represents information about a user.
type User struct {
	UserId   int    `json:"user_id,omitempty"`
	Name     string `json:"name"`
	MobileNo string `json:"mobile_no"`
}
