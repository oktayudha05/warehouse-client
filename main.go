package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"warehouse-client/models"
)

const apiURL = "http://localhost:3000"
var token, role string

// Fungsi login
func login(username, password string) error {
	loginReq := models.LoginRequest{Username: username, Password: password}
	reqBody, _ := json.Marshal(loginReq)
	url := fmt.Sprintf("%s/%s/login", apiURL, role)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var loginRes models.LoginRes
	err = json.Unmarshal(body, &loginRes)
	if err != nil {
		return err
	}

	token = loginRes.Token
	return nil
}

// Fungsi register
func registerKaryawan(username, password, nama, jabatan string) error {
	registerReq := models.RegisterKaryawanRequest{Username: username, Password: password, Nama: nama, Jabatan: jabatan}
	reqBody, _ := json.Marshal(registerReq)
	url := fmt.Sprintf("%s/karyawan/register", apiURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("register failed: %s", resp.Status)
	}
	return nil
}

func registerPengunjung(username, password string) error {
	registerReq := models.RegisterPengunjungRequest{Username: username, Password: password}
	reqBody, _ := json.Marshal(registerReq)
	url := fmt.Sprintf("%s/pengunjung/register", apiURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("register failed: %s", resp.Status)
	}
	return nil
}

// Fungsi untuk tambah barang
func tambahBarang() error {
	var nama string
	var jumlah int
	fmt.Println("Masukkan nama barang:")
	fmt.Scanln(&nama)
	fmt.Println("Masukkan jumlah barang:")
	fmt.Scanln(&jumlah)

	barangReq := models.BarangRequest{Nama: nama, Jumlah: jumlah}
	reqBody, _ := json.Marshal(barangReq)
	req, err := http.NewRequest("POST", apiURL+"/barang", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to add barang: %s", resp.Status)
	}
	fmt.Println("Barang berhasil ditambahkan!")
	return nil
}

// Fungsi untuk lihat barang
func lihatBarang() error {
	req, err := http.NewRequest("GET", apiURL+"/barang", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get barang: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var tampilBarang []models.TampilBarang
	err = json.Unmarshal(body, &tampilBarang)
	if err != nil {
		return err
	}
	fmt.Println("Barang yang ada:")
	for _, barang := range tampilBarang {
		fmt.Printf("Nama: %s, Jumlah: %d, Harga: %d\n", barang.Nama, barang.Jumlah, barang.Harga)
	}
	return nil
}

// Menampilkan menu setelah login
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
			err := lihatBarang()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			err := tambahBarang()
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
			err := lihatBarang()
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

// Fungsi dashboard
func dashboard() {
	fmt.Println("\nSelamat datang di Aplikasi Warehouse!")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Exit")
	fmt.Print("Pilih opsi: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		fmt.Println("Login sebagai (karyawan/pengunjung):")
		fmt.Scanln(&role)
		if role != "karyawan" && role != "pengunjung" {
			fmt.Println("Role tidak valid.")
			return
		}
		var username, password string
		fmt.Print("Username: ")
		fmt.Scanln(&username)
		fmt.Print("Password: ")
		fmt.Scanln(&password)
		err := login(username, password)
		if err != nil {
			fmt.Println("Login gagal:", err)
			return
		}
		fmt.Println("Login berhasil!")
		showMenu()
	case 2:
		fmt.Println("1. Register Karyawan")
		fmt.Println("2. Register Pengunjung")
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
			err := registerKaryawan(username, password, nama, jabatan)
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
			err := registerPengunjung(username, password)
			if err != nil {
				fmt.Println("Register gagal:", err)
				return
			}
			fmt.Println("Register berhasil!")
		}
	case 3:
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
