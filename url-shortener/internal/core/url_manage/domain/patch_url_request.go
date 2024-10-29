package domain

type PatchUrlRequest struct {
	LongUrl    string `json:"url" binding:"required"`
	UserId     string `json:"user_id" binding:"required"`
	ShortUrlId string
}
