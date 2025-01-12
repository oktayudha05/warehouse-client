package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"warehouse-client/models"
)

func RegisterPengunjung(username, password, apiURL string) error {
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