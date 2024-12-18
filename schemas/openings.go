package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Opening struct {
	ID        uuid.UUID `gorm:"type:text;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete support
	Role      string
	Company   string
	Location  string
	Remote    bool
	Link      string
	Salary    int64
}

type OpeningResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}

func (opening *Opening) BeforeCreate(tx *gorm.DB) (err error) {
	if opening.ID == uuid.Nil {
		opening.ID = uuid.New() // Generate UUID in Go
	}
	return
}
