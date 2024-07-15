package structs

type RegistrationReq struct {
	UserName string `json:"userName" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegistrationRes struct {
	Status string `json:"status"`
	Message  string `json:"message"`
	UserName string `json:"userName,omitempty"`
	Code int `json:"code"`
}