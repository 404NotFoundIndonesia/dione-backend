package dto

type UserData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	AvatarUrl string `json:"avatar_url"`
	Role      string `json:"role,omitempty"`
}
