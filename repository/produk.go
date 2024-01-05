package repository

import (
	"context"
	"fmt"

	"github.com/toko-elektronik/entity"
)

func (r *repository) InsertProduk(ctx context.Context, produk entity.Produk) (int64, error) {
	query := "INSERT INTO produk(kategori_id,nama,foto,detail,harga,ketersedian_stok) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, produk.KategoriId, produk.Nama, produk.Foto, produk.Detail, produk.Harga, produk.KetersedianStok).Scan(&produk.Id)
	if err != nil {
		return 0, err
	}
	return produk.Id, nil
}

func (r *repository) GetProduk(ctx context.Context) ([]entity.Produk, error) {
	var produk []entity.Produk
	query := "SELECT id,kategori_id,nama,foto,detail,harga,ketersedian_stok FROM produk"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return produk, err
	}
	defer rows.Close()

	for rows.Next() {
		var pro entity.Produk

		if err := rows.Scan(&pro.Id, &pro.KategoriId, &pro.Nama, &pro.Foto, &pro.Detail, &pro.Harga, &pro.KetersedianStok); err != nil {
			return produk, err
		}
		pro.Foto = "http://localhost:8080/assets/" + pro.Foto
		produk = append(produk, pro)
	}
	return produk, nil
}

func (r *repository) UpdateProduk(ctx context.Context, produk entity.Produk) (int64, error) {
	query := "UPDATE produk SET kategori_id$2,nama$3,foto$4,detail$5,harga$6,ketersedian_stok$7 WHERE id=$1 RETURNING id"
	err := r.db.QueryRowContext(ctx, query, produk.Id, produk.KategoriId, produk.Nama, produk.Foto, produk.Detail, produk.Harga, produk.KetersedianStok).Scan(&produk.Id)

	if err != nil {
		return 0, err
	}
	return produk.Id, nil
}

func (r *repository) DeleteProduk(ctx context.Context, id int64) error {
	query := "DELETE FROM produk WHERE id =$1"
	_, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
