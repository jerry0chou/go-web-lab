package models

type Teacher struct {
	ID      uint     `json:"id" gorm:"primaryKey"`
	Name    string   `json:"name" gorm:"not null"`
	Email   string   `json:"email" gorm:"uniqueIndex;not null"`
	Subject string   `json:"subject"`
	Courses []Course `json:"courses,omitempty" gorm:"foreignKey:TeacherID"`
}
