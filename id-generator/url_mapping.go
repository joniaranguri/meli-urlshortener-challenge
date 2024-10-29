package main

import (
	"fmt"
	"gorm.io/gorm"
)

// UrlMapping represents a row in the url_mapping table
type UrlMapping struct {
	ShortUrlId string  `json:"short_url" gorm:"column:short_url;primary_key"`
	LongUrl    *string `json:"long_url" gorm:"column:long_url"`
	UserId     *string `json:"user_id" gorm:"column:user_id"`
	Active     bool    `json:"active" gorm:"column:active"`
}

// TableName overrides the table name used by UrlMapping to `url_mapping`
func (u *UrlMapping) TableName() string {
	return "url_mapping"
}

// generateAndInsertIds generates IDs in base62 and inserts them into the database
func generateAndInsertIds(db *gorm.DB, startId, count int) (int, error) {
	const batchSize = 10000

	for i := 0; i < count; i += batchSize {
		end := i + batchSize
		if end > count {
			end = count
		}

		var urlMappings []UrlMapping

		for j := i; j < end; j++ {
			id := startId + j
			base62Id := toBase62(id)
			urlMappings = append(urlMappings, UrlMapping{
				ShortUrlId: base62Id,
				LongUrl:    nil,
				UserId:     nil,
				Active:     true,
			})
		}

		if err := db.CreateInBatches(urlMappings, 1000).Error; err != nil {
			return startId + i, fmt.Errorf("failed to insert batch starting with ID %d: %v", startId+i, err)
		}
	}

	return startId + count, nil
}

// updateCurrentId updates the current_id in the ids_index table
func updateCurrentId(db *gorm.DB, currentId int) error {
	return db.Table("ids_index").Update("current_id", fmt.Sprintf("%d", currentId)).Error
}
