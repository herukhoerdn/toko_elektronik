package repository

import (
	"context"
	"fmt"

	"github.com/toko-elektronik/entity"
)

func (r *repository) InsertCheckout(ctx context.Context, checkout entity.Checkout) (int64, error) {
	query := "INSERT INTO checkout(user_id,username,produk_id,nama_produk,kuantiti,total_harga,metode_pembayaran) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, checkout.UserId, checkout.Username, checkout.ProdukId, checkout.NamaProduk, checkout.Kuantiti, checkout.TotalHarga, checkout.MetodePembayaran).Scan(&checkout.Id)
	if err != nil {
		return 0, err
	}
	return checkout.Id, nil
}

func (r *repository) GetCheckout(ctx context.Context) ([]entity.Checkout, error) {
	var checkout []entity.Checkout
	query := "SELECT id,user_id,username,produk_id,nama_produk,kuantiti,total_harga,metode_pembayaran FROM checkout"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return checkout, err
	}
	defer rows.Close()

	for rows.Next() {
		var chec entity.Checkout

		if err := rows.Scan(&chec.Id, chec.UserId, chec.Username, chec.ProdukId, chec.NamaProduk, chec.Kuantiti, chec.TotalHarga, chec.MetodePembayaran); err != nil {
			return checkout, err
		}
		checkout = append(checkout, chec)
	}
	return checkout, nil
}

func (r *repository) UpdateCheckout(ctx context.Context, checkout entity.Checkout) (int64, error) {
	query := "UPDATE checkout SET user_id$2,username$3,produk_id$4,nama_produk$5,kuantiti$6,total_harga$7,metode_pembayaran$8 WHERE id=$1 RETURNING id "
	err := r.db.QueryRowContext(ctx, query, checkout.Id, checkout.UserId, checkout.Username, checkout.ProdukId, checkout.NamaProduk, checkout.Kuantiti, checkout.TotalHarga, checkout.MetodePembayaran).Scan(&checkout.Id)

	if err != nil {
		return 0, err
	}
	return checkout.Id, nil
}

func (r * repository) DeleteCheckout(ctx context.Context,id int64) error{
	query := "DELETE FROM checkout WHERE id =$1"
	_,err := r.db.ExecContext(ctx,query,id)

	if err !=nil {
		fmt.Println("Error:",err)
		return err
	}
	return nil
}
