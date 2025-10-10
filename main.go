package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"errors"
	"log"
	"os"
	"io"
	"strings"
	"bytes"

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


func main() {

}