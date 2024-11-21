package models

type Product struct {
	Id         int64  `json:"id"`
	NamaProduk string `json:"nama_produk"`
	Deskripsi  string `json:"deskripsi"`
}
