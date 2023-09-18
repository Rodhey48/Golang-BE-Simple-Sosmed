package base

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Base contains common columns for all tables.
type BaseEntity struct {
	Id        string    `json:"id" gorm:"type:uuid;primary_key;unique;default:gen_random_uuid()"`
	IsActive  bool      `json:"is_active" gorm:"default:true;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create;type:time;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"<-:update;type:time;default:CURRENT_TIMESTAMP;not null"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseEntity) BeforeCreate(scope *gorm.DB) (err error) {
	// UUID version 4
	base.Id = uuid.NewString()
	// scope.SetColumn("Id", uuid)
	return
}
