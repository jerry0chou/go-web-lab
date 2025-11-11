package models

type Enrollment struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	StudentID uint    `json:"student_id" gorm:"not null"`
	CourseID  uint    `json:"course_id" gorm:"not null"`
	Grade     *string `json:"grade,omitempty"`
	Student   Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	Course    Course  `json:"course,omitempty" gorm:"foreignKey:CourseID"`
}
