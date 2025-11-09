package models

type Course struct {
	ID          uint         `json:"id" gorm:"primaryKey"`
	Title       string       `json:"title" gorm:"not null"`
	Description string       `json:"description"`
	TeacherID   uint         `json:"teacher_id" gorm:"not null"`
	Teacher     Teacher      `json:"teacher,omitempty" gorm:"foreignKey:TeacherID"`
	Enrollments []Enrollment `json:"enrollments,omitempty" gorm:"foreignKey:CourseID"`
}
