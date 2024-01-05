package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/toko-elektronik/entity"
)

type CheckoutHandler struct {
	Id               int64   `json:"id"`
	UserId           int64   `json:"user_id"`
	Username         string  `json:"username"`
	ProdukId         int64   `json:"produk_id"`
	NamaProduk       string  `json:"nama_produk"`
	Kuantiti         int64   `json:"kuantiti"`
	TotalHarga       float64 `json:"total_harga"`
	MetodePembayaran string  `json:"metode_pembayaran"`
}

func (h *handler) CreateCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	var checkout CheckoutHandler
	if err := json.NewDecoder(r.Body).Decode(&checkout); err != nil {
		log.Println("error:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	checkoutID, err := h.uc.InsertCheckout(context.Background(), entity.Checkout{
		UserId:           checkout.UserId,
		Username:         checkout.Username,
		ProdukId:         checkout.ProdukId,
		NamaProduk:       checkout.NamaProduk,
		Kuantiti:         checkout.Kuantiti,
		TotalHarga:       checkout.TotalHarga,
		MetodePembayaran: checkout.MetodePembayaran,
	})
	if err != nil {
		log.Println("error:", err)
		http.Error(w, fmt.Sprintf("Error creating : %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": checkoutID})
}

func (h *handler) GetCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	checkout, err := h.uc.GetCheckout(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Get : %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("contetn-type", "application/json")
	json.NewEncoder(w).Encode(checkout)
}
func (h *handler) UpdateCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	var checkout CheckoutHandler
	if err := json.NewDecoder(r.Body).Decode(&checkout); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid id : %v", err), http.StatusBadRequest)
		return
	}
	checkoutID, err := h.uc.UpdateCheckout(context.Background(), entity.Checkout{
		Id:               id,
		UserId:           checkout.UserId,
		Username:         checkout.Username,
		ProdukId:         checkout.ProdukId,
		NamaProduk:       checkout.NamaProduk,
		Kuantiti:         checkout.Kuantiti,
		TotalHarga:       checkout.TotalHarga,
		MetodePembayaran: checkout.MetodePembayaran,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro creating: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": checkoutID})
}

func (h *handler) DeleteCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting :%v", err), http.StatusInternalServerError)
		return
	}
	err = h.uc.DeleteCheckout(context.Background(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error delete : %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "sucses delete"})

}
