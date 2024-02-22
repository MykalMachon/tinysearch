package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	SourceID  uint      `gorm:"type:uuid REFERENCES sources(id)"`
	Source    Source    `gorm:"foreignKey:SourceID;AssociationForeignKey:ID"`
	Title     string
	Content   string
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`
}
