package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"book-trade/config"
)

func UploadToImgur(base64Image string) (string, error) {
	clientID := config.GetEnv("IMGUR_CLIENT_ID", "");

	url := "https://api.imgur.com/3/image"

	data := map[string]string{
		"image": base64Image,
	}
	payload, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Client-ID " + clientID)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	if !result["success"].(bool) {
		return "", fmt.Errorf("Erro ao enviar mensagem: %v", result)
	}

	dataResult := result["data"].(map[string]interface{})
	link := dataResult["link"].(string)

	return link, nil
}