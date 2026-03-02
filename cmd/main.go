package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/pinghoyk/artifacts_mmo/internal/bot"
)

func findProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "."
}

func main() {
	projectRoot := findProjectRoot()
	envPath := filepath.Join(projectRoot, ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Предупреждение: не удалось загрузить .env из %s: %v", envPath, err)
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("Ошибка: переменная окружения TOKEN не задана")
	}

	gamebot := bot.NewBot(token)
	if err := gamebot.LoadCharacters(); err != nil {
		fmt.Println("Ошибка загрузки персонажей:", err)
		return
	}
	fmt.Printf("\n[DONE] Загружено персонажей: %d\n\n", len(gamebot.GetAllCharacters()))

}
