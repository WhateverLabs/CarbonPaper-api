package models

import (
	"carbon-paper/src/types"
	"time"
)

type Paste struct {
	ID string `gorm:"primaryKey,uniqueIndex"`

	PasteRequestBody types.PasteRequestBody `gorm:"embedded"`

	CreatedAt time.Time
	ExpiresAt time.Time
}
