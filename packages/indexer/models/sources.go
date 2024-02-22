package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Source struct {
	gorm.Model
	Id            uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name          string
	Description   string
	Url           string    `gorm:"unique"`
	LastIndexedAt time.Time `gorm:"default:null"`
}
