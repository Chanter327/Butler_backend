package structs

type DeleteReq struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type DeleteRes struct {
	Status string`json:"status"`
	Message string `json:"message"`
	Code int `json:"code"`
}