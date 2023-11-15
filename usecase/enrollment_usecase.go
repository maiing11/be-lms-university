package usecase

import (
	"fmt"

	"enigmacamp.com/be-lms-university/model/dto"
	"enigmacamp.com/be-lms-university/model/entity"
	"enigmacamp.com/be-lms-university/repository"
)

type EnrollmentUseCase interface {
	RegisterNewEnrollment(payload dto.EnrollmentRequestDto) (entity.Enrollment, error)
}

type enrollmentUseCase struct {
	repo     repository.EnrollmentRepository
	userUC   UserUseCase
	courseUC CourseUseCase
}

func (e *enrollmentUseCase) RegisterNewEnrollment(payload dto.EnrollmentRequestDto) (entity.Enrollment, error) {
	var newEnrollmentDetails []entity.EnrollmentDetail
	course, err := e.courseUC.FindById(payload.CourseId)
	if err != nil {
		return entity.Enrollment{}, err
	}

	for _, v := range payload.Users {
		user, err := e.userUC.FindById(v)
		if err != nil {
			return entity.Enrollment{}, err
		}
		newEnrollmentDetails = append(newEnrollmentDetails, entity.EnrollmentDetail{User: user})
	}

	newEnrollment := entity.Enrollment{
		Course:            course,
		EnrollmentDetails: newEnrollmentDetails,
	}

	enrollment, err := e.repo.Create(newEnrollment)
	if err != nil {
		return entity.Enrollment{}, fmt.Errorf("failed to create enrollment: %s", err.Error())
	}

	return enrollment, nil
}

func NewEnrollmentUseCase(
	repo repository.EnrollmentRepository,
	userUC UserUseCase,
	courseUC CourseUseCase,
) EnrollmentUseCase {
	return &enrollmentUseCase{
		repo:     repo,
		userUC:   userUC,
		courseUC: courseUC,
	}
}
