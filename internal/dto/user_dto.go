package dto

type RegisterUserDto struct{
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type RegisterUserResponseDto struct{
	ID string `json:"id"`
}