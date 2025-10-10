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


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}

}