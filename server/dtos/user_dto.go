package dtos

type UserSignUpFormDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserLoginFormDTO struct {
	Email string `json:"email;binding:required,min=2,max=100"`
	Password string `json:"password;binding:required,min=8,max=100" `
}
