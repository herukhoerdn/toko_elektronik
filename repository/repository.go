package repository

import (
	"context"
	"database/sql"

	"github.com/toko-elektronik/entity"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	InsertUser(ctx context.Context, user entity.User) (int64, error)
	GetUsers(ctx context.Context) ([]entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (int64, error)
	DeleteUser(ctx context.Context, id int64) error

	InsertKategori(ctx context.Context, kategori entity.Kategori) (int64, error)
	GetKategori(ctx context.Context) ([]entity.Kategori, error)
	Updatekategori(ctx context.Context, kategori entity.Kategori) (int64, error)
	DeleteKategori(ctx context.Context, id int64) error

	InsertProduk(ctx context.Context, produk entity.Produk) (int64, error)
	GetProduk(ctx context.Context) ([]entity.Produk, error)
	UpdateProduk(ctx context.Context, kategori entity.Produk) (int64, error)
	DeleteProduk(ctx context.Context, id int64) error

	GetPembayaran(ctx context.Context) ([]entity.Pembayaran, error)

	InsertCheckout(ctx context.Context, checkout entity.Checkout) (int64, error)
	GetCheckout(ctx context.Context) ([]entity.Checkout, error)
	UpdateCheckout(ctx context.Context, checkout entity.Checkout) (int64, error)
	DeleteCheckout(ctx context.Context, id int64) error

	//login
	Login(ctx context.Context, user entity.User) (entity.User, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
