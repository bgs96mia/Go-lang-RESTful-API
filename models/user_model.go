package models

type UserResponse struct {
	Username  string `json:"username,omitempty"`
	Name      string `json:"name,omitempty"`
	Token     string `json:"token,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Name     string `json:"name" validate:"required,max=100"`
}

type UpdateUserRequest struct {
	Username string `json:"-" validate:"required,max=100"`
	Password string `json:"password,omitempty" validate:"max=100"`
	Name     string `json:"name,omitempty" validate:"max=100"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
}

type GetUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
}
