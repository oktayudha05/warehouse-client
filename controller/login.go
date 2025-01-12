package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"warehouse-client/models"
)

func Login(username, password, apiURL, role string, token *string) error {
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

	*token = loginRes.Token
	return nil
}