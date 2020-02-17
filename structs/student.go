package structs

import "time"

// Student Student model
type Student struct {
	ID        uint       `gorm:"primary_key" gorm:"column:id"`
	FirstName string     `gorm:"column:first_name"`
	LastName  string     `gorm:"column:last_name"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `sql:"index"`
}
