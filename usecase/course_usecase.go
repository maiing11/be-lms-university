package usecase

import (
	"fmt"

	"enigmacamp.com/be-lms-university/model/entity"
	"enigmacamp.com/be-lms-university/repository"
)

type CourseUseCase interface {
	FindById(id string) (entity.Course, error)
}

type courseUseCase struct {
	repo repository.CourseRepository
}

func (c *courseUseCase) FindById(id string) (entity.Course, error) {
	course, err := c.repo.Get(id)
	if err != nil {
		return entity.Course{}, fmt.Errorf("course with ID %s not found", id)
	}
	return course, nil
}

func NewCourseUseCase(repo repository.CourseRepository) CourseUseCase {
	return &courseUseCase{repo: repo}
}
