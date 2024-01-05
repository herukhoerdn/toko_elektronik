package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/toko-elektronik/usecase"
)

type handler struct {
	uc usecase.Usecase
}

func InitRoute(r *mux.Router, uc usecase.Usecase) {
	h := handler{
		uc: uc,
	}

	// Define API routes
	r.HandleFunc("/users/{id}", h.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users", h.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users", h.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", h.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/kategori/{id}", h.UpdateKategoriHandler).Methods("PUT")
	r.HandleFunc("/kategori", h.CreateKategoriHandler).Methods("POST")
	r.HandleFunc("/kategori", h.GetKategoriHandler).Methods("GET")
	r.HandleFunc("/kategori/{id}", h.DeleteKategoriHandler).Methods("DELETE")

	r.HandleFunc("/produk/{id}", h.UpdateKategoriHandler).Methods("PUT")
	r.HandleFunc("/produk", h.CreateProdukHandler).Methods("POST")
	r.HandleFunc("/produk", h.GetProdukHandler).Methods("GET")
	r.HandleFunc("/produk/{id}", h.DeleteProdukHandler).Methods("DELETE")

	r.HandleFunc("/metode_pembayaran", h.GetPembayaranHandler).Methods("GET")

	r.HandleFunc("/checkout/{id}", h.UpdateCheckoutHandler).Methods("PUT")
	r.HandleFunc("/checkout", h.CreateCheckoutHandler).Methods("POST")
	r.HandleFunc("/checkout", h.GetProdukHandler).Methods("GET")
	r.HandleFunc("/checkout/{id}", h.DeleteCheckoutHandler).Methods("DELETE")
	//login
	r.HandleFunc("/login", h.LoginHandler).Methods("POST")

	assetsDir := http.Dir("./assets")
	assetsHandler := http.StripPrefix("/assets/", http.FileServer(assetsDir))
	r.PathPrefix("/assets/").Handler(assetsHandler)

}
