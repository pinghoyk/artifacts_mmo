package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"os"
	"io"
	"strings"

	"github.com/joho/godotenv"
)

type Character struct {
	Name 		string 			`json:"name"`
	Account 	string 			`json:"account"`
	Skin 		string 			`json:"skin"`
	Level 		int 			`json:"level"`
	Xp 			int 			`json:"xp"`
	Gold 		int 			`json:"gold"`
	Hp 			int 			`json:"hp"`
	Cooldown	int 			`json:"cooldown"`
	X 			int 			`json:"x"`
	Y 			int 			`json:"y"`
	Inventory []InventoryItem 	`json:"inventory"`
}

type InventoryItem struct {
	Slot 		int 	`json:"slot"`
	Code 		string 	`json:"code"`
	Quantity 	int 	`json:"quantity"`
}

func mustGetEnv (key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Токен или ссылка не существует! %s", key )
	}
	return value
}

func createRequest(url, token string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	return req, nil
}

func getCharacter(client *http.Client, url, token string) (*Character, error) {
    req, err := createRequest(url, token)
    if err != nil {
        return nil, err
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API вернуло ошибку: %s", resp.Status)
    }

    body, err := io.ReadAll(resp.Body)
    fmt.Println("Сырой ответ сервера (получение персонажа):", string(body))
    if err != nil {
        return nil, err
    }

    var response struct {
        Data Character `json:"data"`
    }
    if err := json.Unmarshal(body, &response); err != nil {
        return nil, err
    }

    return &response.Data, nil
}


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}

	token := mustGetEnv("API_TOKEN")
	baseURL := mustGetEnv("URL")
	characterName := mustGetEnv("CHARACTER_NAME")

	characterURL := fmt.Sprintf("%s/characters/%s", strings.TrimSuffix(baseURL, "/"), characterName)

	client := &http.Client{}
	
	character, err := getCharacter(client, characterURL, token)
	if err != nil {
	    log.Fatal("Ошибка получения персонажа:", err)
	}

	if character == nil {
	    log.Fatal("Персонаж не найден")
	}
	
	fmt.Printf("Текущая позиция: (%d, %d)\n", character.X, character.Y)

}