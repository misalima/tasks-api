package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"tasks-api/internal/core/domain"
	"tasks-api/internal/core/interfaces/irepository"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

var _ irepository.UserLoader = (*UserRepository)(nil)

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) InsertUser(ctx context.Context, user domain.User) (*domain.User, error) {
	sql := `INSERT INTO users (id, username, password, email) VALUES ($1, $2, $3, $4) RETURNING id;`

	err := r.pool.QueryRow(ctx, sql, user.ID, user.Username, user.Password, user.Email).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	sql := `SELECT id, username, password, email FROM users WHERE email = $1;`
	row := r.pool.QueryRow(ctx, sql, email)
	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	sql := `SELECT id, username, password, email FROM users WHERE id = $1;`
	row := r.pool.QueryRow(ctx, sql, id)
	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
