package repository

import (
	"context"

	"github.com/toko-elektronik/entity"
)

func (r *repository) Login(ctx context.Context, user entity.User) (entity.User, error) {
	var resp entity.User
	query := "SELECT username, password FROM users WHERE username = $1 AND password = $2"
	row := r.db.QueryRowContext(ctx, query, user.Username, user.Password)
	if row.Err() != nil {
		return user, row.Err()
	}

	if err := row.Scan(&resp.Username, &resp.Password); err != nil {
		return resp, err
	}

	return resp, nil

}
