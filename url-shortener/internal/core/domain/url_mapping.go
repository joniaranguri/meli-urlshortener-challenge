package domain

type UrlMapping struct {
	ShortUrlId string `json:"short_url" gorm:"column:short_url;primary_key"`
	LongUrl    string `json:"long_url" gorm:"column:long_url"`
	UserId     string `json:"user_id" gorm:"column:user_id"`
	Active     bool   `json:"active" gorm:"column:active"`
}

// TableName overrides the table name used by UrlMapping to `url_mapping`
func (u *UrlMapping) TableName() string {
	return "url_mapping"
}
