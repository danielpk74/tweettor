package users

// ResponseLogin Get the token string when a user logged in.
type ResponseLogin struct {
	Token string `json:"token,omitempty"`
}
