package response

type AuthResp struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	// ExpireAt
}
