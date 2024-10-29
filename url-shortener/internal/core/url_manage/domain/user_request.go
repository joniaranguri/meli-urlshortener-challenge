package domain

type UserRequest struct {
	UserId string `json:"user_id" binding:"required"`
}
