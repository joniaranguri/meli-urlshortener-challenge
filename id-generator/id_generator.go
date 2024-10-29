package main

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// toBase62 encodes an integer to a base62 string
func toBase62(num int) string {
	result := ""
	for num > 0 {
		result = string(base62Chars[num%62]) + result
		num /= 62
	}

	// Pad the result with leading zeroes to ensure it's 7 characters long
	if len(result) < 7 {
		result = strings.Repeat("0", 7-len(result)) + result
	}
	return result
}

// getCurrentId retrieves the current ID from the ids_index table
func getCurrentId(db *gorm.DB) (int, error) {
	var index struct {
		CurrentID string
	}

	err := db.Table("ids_index").Select("current_id").Take(&index).Error
	if err != nil {
		return 0, fmt.Errorf("failed to get current_id: %v", err)
	}

	currentId, err := strconv.Atoi(index.CurrentID)
	if err != nil {
		return 0, fmt.Errorf("failed to convert current_id to int: %v", err)
	}
	return currentId, nil
}
