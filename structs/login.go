package structs

type LoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	UserName string `json:"userName,omitempty"`
	Code int `json:"code"`
}