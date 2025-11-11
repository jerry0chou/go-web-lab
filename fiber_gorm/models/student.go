package models

type Student struct {
	ID          uint         `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"not null"`
	Email       string       `json:"email" gorm:"uniqueIndex;not null"`
	Age         int          `json:"age"`
	Enrollments []Enrollment `json:"enrollments,omitempty" gorm:"foreignKey:StudentID"`
}
