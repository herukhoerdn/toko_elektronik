package entity

type Checkout struct {
	Id               int64   `db:"id"`
	UserId           int64   `db:"user_id"`
	Username         string  `db:"username"`
	ProdukId         int64   `db:"produk_id"`
	NamaProduk       string  `db:"nama_produk"`
	Kuantiti         int64   `db:"kuantiti"`
	TotalHarga       float64 `db:"total_harga"`
	MetodePembayaran string  `db:"metode_pembayaran"`
}
