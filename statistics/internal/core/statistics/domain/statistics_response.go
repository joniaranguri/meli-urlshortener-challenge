package domain

type StatisticsResponse struct {
	ShortUrlId string `json:"short_url"`
	Clicks     uint64 `json:"clicks"`
}
