package entity

type Produk struct {
	Id              int64   `db:"id"`
	KategoriId      int64   `db:"kategori_id"`
	Nama            string  `db:"nama"`
	Foto            string  `db:"foto"`
	Detail          string  `db:"detail"`
	Harga           float64 `db:"harga"`
	KetersedianStok int64   `db:"ketersedian_stok"`
}
