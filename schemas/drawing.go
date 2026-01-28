package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Drawing struct {
	gorm.Model
	Name      string
	Price     *float64
	Customer  string
	StartDate time.Time
	EndDate   *time.Time
}

type DrawingResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitzero"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Customer  string    `json:"customer"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
