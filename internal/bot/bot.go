// internal/bot/bot.go
package bot

import (
	"fmt"

	"github.com/pinghoyk/artifacts_mmo/internal/api"
	"github.com/pinghoyk/artifacts_mmo/internal/character"
)

type Bot struct {
	client     *api.Client
	characters map[string]*character.Character
}

func NewBot(token string) *Bot {
	return &Bot{
		client:     api.NewClient(token),
		characters: make(map[string]*character.Character),
	}
}

// LoadCharacters загружает персонажей из API и инициализирует их
func (b *Bot) LoadCharacters() error {
	chars, err := b.client.GetCharacters()
	if err != nil {
		return fmt.Errorf("ошибка получения персонажей: %w", err)
	}

	for i := range chars {
		chars[i].ParseCooldown()
		b.characters[chars[i].Name] = &chars[i]
		fmt.Printf("[✓] %s | Оружие: %s | Задача: %s | Инвентарь: %d/%d\n",
			chars[i].Name, chars[i].WeaponSlot, chars[i].Task,
			countInventoryItems(chars[i].Inventory), chars[i].InventoryMaxItems)
	}
	return nil
}

// GetCharacter возвращает персонажа по имени (или nil)
func (b *Bot) GetCharacter(name string) *character.Character {
	return b.characters[name]
}

// GetAllCharacters возвращает срез всех персонажей
func (b *Bot) GetAllCharacters() []*character.Character {
	result := make([]*character.Character, 0, len(b.characters))
	for _, char := range b.characters {
		result = append(result, char)
	}
	return result
}

// 🎯 Твоя "простая функция" — получить инфу о персонаже
func (b *Bot) GetCharacterInfo(name string) map[string]interface{} {
	char := b.characters[name]
	if char == nil {
		return map[string]interface{}{"error": "character not found"}
	}

	return map[string]interface{}{
		"name":           char.Name,
		"level":          char.Level,
		"gold":           char.Gold,
		"position":       fmt.Sprintf("(%d, %d) [%s]", char.X, char.Y, char.Layer),
		"ready":          char.IsReady(),
		"inventory_full": char.IsInventoryFull(2),
		"task": map[string]interface{}{
			"name":     char.Task,
			"type":     char.TaskType,
			"progress": fmt.Sprintf("%d/%d", char.TaskProgress, char.TaskTotal),
		},
	}
}

// ShowStatus — красивая печать статуса (для отладки/лога)
func (b *Bot) ShowStatus(name string) {
	info := b.GetCharacterInfo(name)
	if err, ok := info["error"]; ok {
		fmt.Printf("❌ %v\n", err)
		return
	}

	fmt.Printf("\n🎮 %s (Lvl %d)\n", info["name"], info["level"])
	fmt.Printf("💰 Золото: %d\n", info["gold"])
	fmt.Printf("📍 Позиция: %s\n", info["position"])
	fmt.Printf("⏱️  Готов: %v\n", info["ready"])
	fmt.Printf("🎒 Инвентарь: %s\n",
		map[bool]string{true: "⚠️ почти полон", false: "✅ ок"}[info["inventory_full"]])
	fmt.Println()
}

// Вспомогательная функция — считаем заполненные слоты инвентаря
func countInventoryItems(inventory []character.InventoryItem) int {
	count := 0
	for _, item := range inventory {
		if item.Code != "" && item.Quantity > 0 {
			count++
		}
	}
	return count
}
