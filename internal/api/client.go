package api

import (
	"artifacts_mmo/internal/character"
	"fmt"
)

type Client struct {
	Token string
}

func NewClient(token string) *Client {
	return &Client{Token: token}
}

// Заглушка: потом здесь будет реальный HTTP запрос
func (c *Client) GetCharacters() ([]character.Character, error) {
	fmt.Println("[API] Получение списка персонажей...")
	// TODO: Реальный запрос к API
	return []character.Character{}, nil
}

// Заглушка для сбора ресурсов
func (c *Client) StartGathering(charName string) error {
	fmt.Printf("[API] %s начинает сбор ресурсов\n", charName)
	return nil
}

// Заглушка для сдачи в банк
func (c *Client) DepositResources(charName string) error {
	fmt.Printf("[API] %s сдает ресурсы в банк\n", charName)
	return nil
}
