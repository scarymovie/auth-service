package repository

import (
	"auth-service/internal/repository"
	userpb "auth-service/proto"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type postgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) repository.UserRepository {
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) Create(name string) (*userpb.User, error) {
	id := uuid.New()

	query := `INSERT INTO users (uid, name) VALUES ($1, $2)`
	_, err := r.db.Exec(query, id, name)
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &userpb.User{
		Id:   id.String(),
		Name: name,
	}, nil
}

func (r *postgresUserRepository) FindById(uid string) (*userpb.User, error) {
	query := `SELECT uid, name FROM users WHERE uid = $1`

	row := r.db.QueryRow(query, uid)

	var user userpb.User
	if err := row.Scan(&user.Id, &user.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("query error: %w", err)
	}

	return &user, nil
}
