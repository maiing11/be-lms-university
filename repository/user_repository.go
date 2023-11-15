package repository

import (
	"database/sql"

	"enigmacamp.com/be-lms-university/model/entity"
)

// TODO
/*
1. Kita harus siapkan sebuah kontrak (interface)
2. Interface ini yang akan dilempar ke service lain (injection)
3. Ini berguna untuk memudahkan Unit Testing
4. Biasanya nama interface seperti nama file dan dia ter-expose (public)
5. Setelah itu kita buatkan struct untuk dikirim sebagai receiver method
6. Method-method inilah yang digunakan sebagi isian kontrak dari interface
7. Kita buatkan function sebagai perantara untuk memanggil interface, agar method-method yang dibuat bisa dipanggil keluar (function ini biasanya disebut sebagi constructor, diawali dengan keyword `New`)
*/

type UserRepository interface {
	// CRUD methods
	Get(id string) (entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Get(id string) (entity.User, error) {
	var user entity.User
	err := u.db.QueryRow(`
	SELECT
		id, first_name, last_name, email, username, role, photo, created_at, updated_at
	FROM
		users
	WHERE
		id = $1
	`, id).
		Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Role,
			&user.Photo,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

	// jika scan gagal
	if err != nil {
		return entity.User{}, err
	}

	// jika berhasil
	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
