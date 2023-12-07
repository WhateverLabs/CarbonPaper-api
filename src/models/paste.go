package models

import (
	"carbon-paper/src/types"
	"time"
)

type Paste struct {
	ID string `gorm:"primaryKey,uniqueIndex" json:"id"`

	types.PasteRequestBody `gorm:"embedded" json:",inline"`

	CreatedAt time.Time `json:"-"`
	ExpiresAt time.Time `json:"expiresAt"`
}
