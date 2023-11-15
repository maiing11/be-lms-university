package repository

import (
	"database/sql"
	"time"

	"enigmacamp.com/be-lms-university/model/entity"
)

type EnrollmentRepository interface {
	Create(payload entity.Enrollment) (entity.Enrollment, error)
}

type enrollmentRepository struct {
	db *sql.DB
}

func (e *enrollmentRepository) Create(payload entity.Enrollment) (entity.Enrollment, error) {
	// Transactional
	tx, err := e.db.Begin()
	if err != nil {
		return entity.Enrollment{}, err
	}

	// insert enrollment
	var enrollment entity.Enrollment
	var enrollmentDetails []entity.EnrollmentDetail
	err = tx.QueryRow(`
	INSERT INTO enrollments (course_id, status, updated_at) VALUES ($1, $2, $3)
	RETURNING id, status, created_at, updated_at`,
		payload.Course.Id,
		"active",
		time.Now(),
	).Scan(
		&enrollment.Id,
		&enrollment.Status,
		&enrollment.CreatedAt,
		&enrollment.UpdatedAt,
	)

	if err != nil {
		return entity.Enrollment{}, tx.Rollback()
	}

	for _, v := range payload.EnrollmentDetails {
		var enrollmentDetail entity.EnrollmentDetail
		err = tx.QueryRow(`
		INSERT INTO enrollment_details (enrollment_id, user_id, updated_at) VALUES ($1, $2, $3)
		RETURNING id, enrollment_id, created_at, updated_at`,
			enrollment.Id,
			v.User.Id,
			time.Now(),
		).Scan(
			&enrollmentDetail.Id,
			&enrollmentDetail.EnrollmentId,
			&enrollmentDetail.CreatedAt,
			&enrollmentDetail.UpdatedAt,
		)
		enrollmentDetail.User = v.User
		enrollmentDetails = append(enrollmentDetails, enrollmentDetail)
	}

	if err := tx.Commit(); err != nil {
		return entity.Enrollment{}, err
	}

	enrollment.Course = payload.Course
	enrollment.EnrollmentDetails = enrollmentDetails

	return enrollment, nil
}

func NewEnrollmentRepository(db *sql.DB) EnrollmentRepository {
	return &enrollmentRepository{db: db}
}
