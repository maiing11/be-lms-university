package entity

import "time"

type User struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstName`
	LastName  string    `json:"lastName`
	Email     string    `json:"email"`
	Username  string    `json:"username" form:"username" binding:"required"`
	Password  string    `json:"password" form:"password" binding:"required"`
	Role      string    `json:"role"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u User) IsValidRole() bool {
	return u.Role == "student" || u.Role == "instructor" || u.Role == "admin"
}
