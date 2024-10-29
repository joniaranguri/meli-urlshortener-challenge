package domain

type ShortenUrlRequest struct {
	LongUrl string `json:"url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}
