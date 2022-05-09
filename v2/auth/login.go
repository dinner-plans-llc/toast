package auth

// Message body parameters for a login endpoint request
type Message struct {
	ClientID       string
	ClientSecret   string
	UserAccessType string
}
