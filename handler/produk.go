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

type ProdukHandler struct {
	Id              int64   `json:"id"`
	KategoriId      int64   `json:"kategori_id"`
	Nama            string  `json:"nama"`
	Detail          string  `json:"detail"`
	Harga           float64 `json:"harga"`
	KetersedianStok int64   `json:"ketersediaan_stok"`
}

func (h *handler) CreateProdukHandler(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	log.Printf("error: %v", err)
	// 	http.Error(w, "Error parsing form", http.StatusBadRequest)
	// 	return
	// }
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("error: %v", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	kategoriIdStr := r.FormValue("kategori_id")
	kategoriId, err := strconv.ParseInt(kategoriIdStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid kategoriId : %v", err), http.StatusBadRequest)
		return
	}

	nama := r.Form.Get("nama")
	detail := r.Form.Get("detail")
	hargaStr := r.Form.Get("harga")
	harga, err := strconv.ParseFloat(hargaStr, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid harga: %v", err), http.StatusBadRequest)
		return
	}

	ketersediaanStokStr := r.Form.Get("ketersediaan_stok")
	ketersediaanStok, err := strconv.ParseInt(ketersediaanStokStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ketersediaanStok : %v", err), http.StatusBadRequest)
		return
	}

	// foto, _, err := r.FormFile("foto")
	// if err != nil {
	// 	log.Printf("error: %v", err)
	// 	http.Error(w, "Error parsing form", http.StatusBadRequest)
	// 	return
	// }
	// fileData, err := ioutil.ReadAll(foto)
	// if err != nil {
	// 	log.Printf("error: %v", err)
	// 	http.Error(w, "Error parsing form", http.StatusBadRequest)
	// 	return
	// }

	// var produk ProdukHandler
	// if err := json.NewDecoder(r.Body).Decode(&produk); err != nil {
	// 	http.Error(w, "Invalid request payload", http.StatusBadRequest)
	// 	return
	// }
	produkID, err := h.uc.InsertProduk(context.Background(), entity.Produk{
		KategoriId: kategoriId,
		Nama:       nama,
		// Foto:            fileData,
		Detail:          detail,
		Harga:           harga,
		KetersedianStok: ketersediaanStok,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating produk : %v ", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": produkID})
}

func (h *handler) GetProdukHandler(w http.ResponseWriter, r *http.Request) {
	produk, err := h.uc.GetProduk(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Get produk : %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(produk)
}

func (h *handler) UpdateProdukHandler(w http.ResponseWriter, r *http.Request) {
	var produk ProdukHandler
	if err := json.NewDecoder(r.Body).Decode(&produk); err != nil {
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
	produkID, err := h.uc.UpdateProduk(context.Background(), entity.Produk{
		Id:              id,
		KategoriId:      produk.KategoriId,
		Nama:            produk.Nama,
		Detail:          produk.Detail,
		Harga:           produk.Harga,
		KetersedianStok: produk.KetersedianStok,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating produk: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "appplication/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": produkID})
}

func (h *handler) DeleteProdukHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting produk : %v", err), http.StatusInternalServerError)
		return
	}
	err = h.uc.DeleteProduk(context.Background(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting produk : %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "sucses delete"})
}
