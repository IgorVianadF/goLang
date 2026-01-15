package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Drawing struct {
	gorm.Model
	Name      string
	Price     int64
	Customer  string
	StartDate time.Time
	EndDate   time.Time
}
