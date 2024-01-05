package usecase

import (
	"context"
	"log"

	"github.com/toko-elektronik/entity"
)

func (u *usecase) Login(ctx context.Context, user entity.User) (entity.User, error) {
	login, err := u.repo.Login(ctx, user)
	if err != nil {
		log.Printf("Error login: %v", err)
		return entity.User{}, err
	}
	return login, nil
}
