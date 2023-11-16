package manager

import "enigmacamp.com/be-lms-university/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	CourseUseCase() usecase.CourseUseCase
	EnrollmentUseCase() usecase.EnrollmentUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func (u *useCaseManager) CourseUseCase() usecase.CourseUseCase {
	return usecase.NewCourseUseCase(u.repo.CourseRepo())
}

func (u *useCaseManager) EnrollmentUseCase() usecase.EnrollmentUseCase {
	return usecase.NewEnrollmentUseCase(u.repo.EnrollmentRepo(), u.UserUseCase(), u.CourseUseCase())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
