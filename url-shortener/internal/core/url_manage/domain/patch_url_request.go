package domain

type PatchUrlRequest struct {
	LongUrl    string `json:"url"`
	ShortUrlId string
}
