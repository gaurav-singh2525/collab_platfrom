package repositories

import (
	"database/sql"

	"collab-code-platform/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) EmailExists(email string) (bool, error) {

	var exists bool

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE email = $1
		)
	`

	err := r.db.QueryRow(
		query,
		email,
	).Scan(&exists)

	return exists, err
}

func (r *UserRepository) CreateUser(
	email string,
	passwordHash string,
) error {

	query := `
		INSERT INTO users (
			email,
			password_hash
		)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(
		query,
		email,
		passwordHash,
	)

	return err
}

func (r *UserRepository) GetUserByEmail(
	email string,
) (*models.User, error) {

	query := `
		SELECT
			id,
			email,
			password_hash,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	user := &models.User{}

	err := r.db.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
