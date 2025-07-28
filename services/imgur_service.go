package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"book-trade/config"
)

func UploadToImgur(base64Image string) (link string, deleteHash string, err error) {
	clientID := config.GetEnv("IMGUR_CLIENT_ID", "");

	url := "https://api.imgur.com/3/image"

	data := map[string]string{
		"image": base64Image,
	}
	payload, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Authorization", "Client-ID " + clientID)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	if !result["success"].(bool) {
		return "", "", fmt.Errorf("Erro ao enviar mensagem: %v", result)
	}

	dataResult := result["data"].(map[string]interface{})
	link = dataResult["link"].(string)
	deleteHash = dataResult["deletehash"].(string)

	return deleteHash, link, nil
}

func DeleteFromImgur(deleteHash string) error {
	clientID := config.GetEnv("IMGUR_CLIENT_ID", "")
	url := fmt.Sprintf("https://api.imgur.com/3/image/%s", deleteHash)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Client-ID "+clientID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	if !result["success"].(bool) {
		return fmt.Errorf("Erro ao deletar imagem do Imgur: %v", result)
	}

	return nil
}