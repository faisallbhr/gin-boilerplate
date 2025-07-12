package requests

type UpdateUserRequest struct {
	Name  string `json:"name" binding:"required,min=3,max=255"`
	Email string `json:"email" binding:"required,email"`
}

type UpdatePasswordRequest struct {
	CurrentPassword      string `json:"current_password" binding:"required"`
	NewPassword          string `json:"new_password" binding:"required,min=6,max=255"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=NewPassword"`
}
