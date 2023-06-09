package model

type UserResponse struct {
	ID     uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleId uint   `json:"role_id"`
	// CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt *time.Time `json:"updated_at"`
}

type RefreshResponse struct {
	Access_token string `json:"access_token"`
}

/* type UserResponseEpochDTO struct {
	ID           *uint   `json:"user_id"`
	Name         *string `json:"name"`
	AuthServerId *uint   `json:"auth_server_id"`
	Nip          *string `json:"nip"`
	RoleId       *uint   `json:"role_id"`
	RoleName     *string `json:"role_name"`
	Email        *string `json:"email"`
	Label        *string `json:"label"`
	CreatedAt    *uint64 `json:"created_at"`
	UpdatedAt    *uint64 `json:"updated_at"`
} */
