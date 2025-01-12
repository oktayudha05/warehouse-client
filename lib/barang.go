package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"warehouse-client/models"
)

func LihatBarang(apiURL string, token string) error {
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

func TambahBarang(apiURL, token string) error {
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