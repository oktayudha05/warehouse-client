package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterKaryawanRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
	Jabatan  string `json:"jabatan"`
}

type RegisterPengunjungRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type BarangRequest struct {
	Nama   string `json:"nama"`
	Jumlah int    `json:"jumlah"`
}

type TampilBarang struct {
	Nama   string `json:"nama"`
	Jumlah int    `json:"jumlah"`
	Harga  int    `json:"harga"`
}

type LoginRes struct {
	Data struct {
		Nama     string `json:"nama"`
		Username string `json:"username"`
		Jabatan  string `json:"jabatan"`
	} `json:"data"`
	Message string `json:"message"`
	Token   string `json:"token"`
}