package model

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=4"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4,max=10"`
	RoleId   uint   `json:"role_id"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
