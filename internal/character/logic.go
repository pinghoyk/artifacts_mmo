package character

import "time"

func (c *Character) IsReady() bool {
	return time.Now().After(c.NextActionTime)
}

func (c *Character) IsInventoryFull() bool {
	count := 0
	for _, item := range c.Inventory {
		if item.Code != "" && item.Quantity > 0 {
			count++
		}
	}
	// Оставляем 1-2 слота запаса
	return count >= c.InventoryMaxItems-2
}

// Парсинг времени кулдауна из строки API
func (c *Character) ParseCooldown() error {
	if c.CooldownExpiration == "" {
		c.NextActionTime = time.Time{}
		return nil
	}
	t, err := time.Parse(time.RFC3339, c.CooldownExpiration)
	if err != nil {
		return err
	}
	c.NextActionTime = t
	return nil
}
