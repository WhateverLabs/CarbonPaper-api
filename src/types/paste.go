package types

type PasteRequestBody struct {
	SenderNameHeader     string `gorm:"not null" json:"senderNameHeader" binding:"required"`
	SenderNameCiphertext string `gorm:"not null" json:"senderNameCiphertext" binding:"required"`

	BodyHeader     string `gorm:"not null" json:"bodyHeader" binding:"required"`
	BodyCiphertext string `gorm:"not null" json:"bodyCiphertext" binding:"required"`

	KekHash            string `json:"kekHash"`
	PasswordHashSalt   string `json:"passwordHashSalt"`
	EncryptedKeyHeader string `json:"encryptedKeyHeader"`

	OneView bool `json:"oneView"`

	// gorm ignore
	ExpiresInSeconds int `gorm:"-" json:"expiresInSeconds" binding:"required"`
}
