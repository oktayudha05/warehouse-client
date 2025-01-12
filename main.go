package main

import (
	"fmt"
	"os"

	"warehouse-client/controller"
	"warehouse-client/lib"

	"github.com/joho/godotenv"
)

var token, role, apiURL string
func init(){
	err := godotenv.Load(".env")
	if err != nil {
		panic("error load .env")
	}
	apiURL = os.Getenv("apiURL")
}

func showMenu() {
	if role == "karyawan" {
		var pilihan int
		fmt.Println("\nMenu Karyawan:")
		fmt.Println("1. Lihat Barang")
		fmt.Println("2. Tambah Barang")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			err := lib.LihatBarang(apiURL, token)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			err := lib.TambahBarang(apiURL, token)
			if err != nil {
				fmt.Println("Error:", err)
			}
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	} else if role == "pengunjung" {
		var pilihan int
		fmt.Println("\nMenu Pengunjung:")
		fmt.Println("1. Lihat Barang")
		fmt.Println("2. Exit")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			err := lib.LihatBarang(apiURL, token)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	} else {
		fmt.Println("Role tidak ditemukan.")
	}
}

func dashboard() {
	fmt.Println("\nSelamat datang di Aplikasi Warehouse!")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("0. Exit")
	fmt.Print("Pilih opsi: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		fmt.Println("1. Login Karyawan")
		fmt.Println("2. Login Pengunjung")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih opsi: ")
		var pilihanLog int
		fmt.Scanln(&pilihanLog)
		switch pilihanLog {
		case 1:
			role = "karyawan"
			var username, password string
			fmt.Print("Username: ")
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			fmt.Scanln(&password)
			err := controller.Login(username, password, apiURL, role, &token)
			if err != nil {
				fmt.Println("login gagal: ", err)
				return
			}
			fmt.Printf("login karyawan sebagai %s berhasil\n", username)
			showMenu()
		
		case 2:
			role = "pengunjung"
			var username, password string
			fmt.Print("Username: ")
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			fmt.Scanln(&password)
			err := controller.Login(username, password, apiURL, role, &token)
			if err != nil {
				fmt.Println("login gagal: ", err)
				return
			}
			fmt.Printf("login pengunjung sebagai %s berhasil\n", username)
			showMenu()

		case 0:
			fmt.Println("Kembali ke menu utama")
			return
		}

	case 2:
		fmt.Println("1. Register Karyawan")
		fmt.Println("2. Register Pengunjung")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih opsi: ")
		var pilihanReg int
		fmt.Scanln(&pilihanReg)

		switch pilihanReg {
		case 1:
			var username, password, nama, jabatan string
			fmt.Print("Username: ")
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			fmt.Scanln(&password)
			fmt.Print("Nama Lengkap: ")
			fmt.Scanln(&nama)
			fmt.Print("Jabatan: ")
			fmt.Scanln(&jabatan)
			err := controller.RegisterKaryawan(username, password, nama, jabatan, apiURL)
			if err != nil {
				fmt.Println("Register gagal:", err)
				return
			}
			fmt.Println("Register berhasil!")
		case 2:
			var username, password string
			fmt.Print("Username: ")
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			fmt.Scanln(&password)
			err := controller.RegisterPengunjung(username, password, apiURL)
			if err != nil {
				fmt.Println("Register gagal:", err)
				return
			}
			fmt.Println("Register berhasil!")

		case 0:
			fmt.Println("Kembali ke menu utama")
			return
		}
	case 0:
		fmt.Println("Terima kasih!")
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func main() {
	for {
		dashboard()
	}
}
