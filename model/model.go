package model

type Transaksi struct {
	IDTransaksi      string            `json:"IDTransaksi"`
	Tanggal          string            `json:"tanggal"`
	Jenis            string            `json:"jenis"`
	Jumlah           int               `json:"jumlah"`
	MataUang         string            `json:"mataUang"`
	MetodePembayaran string            `json:"metodePembayaran,omitempty"`
	Detail           *DetailPembelian  `json:"detail,omitempty"`
	Pengirim         *PengirimTransfer `json:"pengirim,omitempty"`
	Penerima         *PenerimaTransfer `json:"penerima,omitempty"`
	PenyediaJasa     string            `json:"penyediaJasa,omitempty"`
	NomorPelanggan   string            `json:"nomorPelanggan,omitempty"`
	PeriodeTagihan   string            `json:"periodeTagihan,omitempty"`
}

type DetailPembelian struct {
	NamaPedagang string `json:"namaPedagang"`
	Kategori     string `json:"kategori"`
	Deskripsi    string `json:"deskripsi"`
}

type PengirimTransfer struct {
	Nama          string `json:"nama"`
	NomorRekening string `json:"nomorRekening"`
}

type PenerimaTransfer struct {
	Nama          string `json:"nama"`
	NomorRekening string `json:"nomorRekening"`
}
