package dto

type EnrollmentRequestDto struct {
	Id       string   `json:"id"`
	CourseId string   `json:"courseId"`
	Users    []string `json:"users"`
	Status   string   `json:"status"`
}
