package request

type AuthReq struct {
	Username string `json:"username" binding:"required,max=20"`
	Passwrod string `json:"password" binding:"required,max=20"`
}
