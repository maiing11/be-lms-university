package entity

import "time"

type Course struct {
	Id              string `json:"id"`
	CourseFullName  string `json:"courseFullName`
	CourseShortName string `json:"courseShortName"`
	Description     string `json:"description"`
	CourseStartDate time.Time
	CourseEndDate   time.Time
	CourseImage     string    `json:"courseImage`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
