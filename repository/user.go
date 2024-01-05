package repository

import (
	"context"
	"fmt"

	"github.com/toko-elektronik/entity"
)

func (r *repository) InsertUser(ctx context.Context, user entity.User) (int64, error) {
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password).Scan(&user.Id)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (r *repository) GetUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	query := "SELECT id,username,password FROM users where username=$1 "
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User

		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			return users, err
		}

		users = append(users, user)
	}
	return users, nil
}
func (r *repository) UpdateUser(ctx context.Context, user entity.User) (int64, error) {
	query := "UPDATE users SET username=$2, password=$3 WHERE id=$1 RETURNING id"
	err := r.db.QueryRowContext(ctx, query, user.Id, user.Username, user.Password).Scan(&user.Id)
	if err != nil {
		return 0, err
	}
	return user.Id, nil

}
func (r *repository) DeleteUser(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id =$1"
	_, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
