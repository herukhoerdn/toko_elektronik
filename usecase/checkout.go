package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/toko-elektronik/entity"
)

func (u *usecase) InsertCheckout(ctx context.Context, checkout entity.Checkout) (int64, error) {
	err := validateCheckout(checkout)
	if err != nil {
		return 0, err
	}

	id, err := u.repo.InsertCheckout(ctx, checkout)

	if err != nil {
		log.Println("Error insert")
		return 0, err
	}
	return id, nil
}

func validateCheckout(checkout entity.Checkout) error {
	if checkout.UserId == 0 {
		return errors.New("masukan user id ")
	}
	if checkout.Username == "" {
		return errors.New("masukan username")
	}
	if checkout.ProdukId == 0 {
		return errors.New("masukan produk id ")
	}
	if checkout.NamaProduk == "" {
		return errors.New("masukan nama produk")
	}
	if checkout.Kuantiti == 0 {
		return errors.New("masukan  jumlah belanjaan")
	}
	if checkout.TotalHarga == 0 {
		return errors.New("jumlah barang ")
	}
	if checkout.MetodePembayaran == "" {
		return errors.New("bisa melalui")
	}
	return nil
}

func (u *usecase) GetCheckout(ctx context.Context) ([]entity.Checkout, error) {
	checkout, err := u.repo.GetCheckout(ctx)
	if err != nil {
		log.Println("Error get")
		return checkout, err
	}
	return checkout, nil
}

func (u *usecase) UpdateCheckout(ctx context.Context, checkout entity.Checkout) (int64, error) {
	err := validateCheckout(checkout)
	if err != nil {
		log.Println("Error update")
		return 0, err
	}
	id, err := u.repo.UpdateCheckout(ctx, checkout)

	if err != nil {
		log.Println("Error update")
		return 0, err
	}
	return id, nil
}

func (u *usecase) DeleteCheckout(ctx context.Context, id int64) error {
	err := u.repo.DeleteCheckout(ctx, id)
	if err != nil {
		log.Println("Error delete")
		return err
	}
	return nil
}
