package character

import "time"

type InventoryItem struct {
	Slot     int    `json:"slot"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type Character struct {
	Name string `json:"name"`

	Cooldown           int    `json:"cooldown"`
	CooldownExpiration string `json:"cooldown_expiration"`

	WeaponSlot string `json:"weapon_slot"`
	Task       string `json:"task"`
	TaskType   string `json:"task_type"`

	InventoryMaxItems int             `json:"inventory_max_items"`
	Inventory         []InventoryItem `json:"inventory"`

	X              int       `json:"x"`
	Y              int       `json:"y"`
	Layer          string    `json:"layer"`
	MapID          int       `json:"map_id"`
	NextActionTime time.Time `json:"-"`
}
