package character

import (
	"time"
)

// InventoryItem представляет предмет в инвентаре
type InventoryItem struct {
	Slot     int    `json:"slot"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

// Effect представляет активный эффект на персонаже
type Effect struct {
	Code  string `json:"code"`
	Value int    `json:"value"`
}

// SkillData вспомогательная структура для удобной работы с навыками в коде
// (не используется для прямого маппинга JSON, так как в JSON поля плоские)
type SkillData struct {
	Level int `json:"level"`
	XP    int `json:"xp"`
	MaxXP int `json:"max_xp"`
}

// Character представляет данные персонажа
type Character struct {
	// --- Основная информация ---
	Name    string `json:"name"`
	Account string `json:"account"`
	Skin    string `json:"skin"`

	// --- Прогресс персонажа ---
	Level int `json:"level"`
	XP    int `json:"xp"`
	MaxXP int `json:"max_xp"`
	Gold  int `json:"gold"`
	Speed int `json:"speed"`

	// --- Навыки (Skills) ---
	// Добыча
	MiningLevel      int `json:"mining_level"`
	MiningXP         int `json:"mining_xp"`
	MiningMaxXP      int `json:"mining_max_xp"`
	WoodcuttingLevel int `json:"woodcutting_level"`
	WoodcuttingXP    int `json:"woodcutting_xp"`
	WoodcuttingMaxXP int `json:"woodcutting_max_xp"`
	FishingLevel     int `json:"fishing_level"`
	FishingXP        int `json:"fishing_xp"`
	FishingMaxXP     int `json:"fishing_max_xp"`

	// Крафт
	WeaponcraftingLevel  int `json:"weaponcrafting_level"`
	WeaponcraftingXP     int `json:"weaponcrafting_xp"`
	WeaponcraftingMaxXP  int `json:"weaponcrafting_max_xp"`
	GearcraftingLevel    int `json:"gearcrafting_level"`
	GearcraftingXP       int `json:"gearcrafting_xp"`
	GearcraftingMaxXP    int `json:"gearcrafting_max_xp"`
	JewelrycraftingLevel int `json:"jewelrycrafting_level"`
	JewelrycraftingXP    int `json:"jewelrycrafting_xp"`
	JewelrycraftingMaxXP int `json:"jewelrycrafting_max_xp"`
	CookingLevel         int `json:"cooking_level"`
	CookingXP            int `json:"cooking_xp"`
	CookingMaxXP         int `json:"cooking_max_xp"`
	AlchemyLevel         int `json:"alchemy_level"`
	AlchemyXP            int `json:"alchemy_xp"`
	AlchemyMaxXP         int `json:"alchemy_max_xp"`

	// --- Характеристики (Stats) ---
	Hp             int `json:"hp"`
	MaxHp          int `json:"max_hp"`
	Haste          int `json:"haste"`
	CriticalStrike int `json:"critical_strike"`
	Wisdom         int `json:"wisdom"`
	Prospecting    int `json:"prospecting"`
	Initiative     int `json:"initiative"`
	Threat         int `json:"threat"`

	// Атака по стихиям
	AttackFire  int `json:"attack_fire"`
	AttackEarth int `json:"attack_earth"`
	AttackWater int `json:"attack_water"`
	AttackAir   int `json:"attack_air"`

	// Урон
	Dmg      int `json:"dmg"`
	DmgFire  int `json:"dmg_fire"`
	DmgEarth int `json:"dmg_earth"`
	DmgWater int `json:"dmg_water"`
	DmgAir   int `json:"dmg_air"`

	// Сопротивления
	ResFire  int `json:"res_fire"`
	ResEarth int `json:"res_earth"`
	ResWater int `json:"res_water"`
	ResAir   int `json:"res_air"`

	// --- Состояние и Эффекты ---
	Effects []Effect `json:"effects"`

	// --- Позиция и Карта ---
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Layer string `json:"layer"`
	MapID int    `json:"map_id"`

	// --- Кулдауны и Действия ---
	Cooldown           int       `json:"cooldown"`
	CooldownExpiration time.Time `json:"cooldown_expiration"`
	NextActionTime     time.Time `json:"-"`

	// --- Экипировка (Slots) ---
	WeaponSlot    string `json:"weapon_slot"`
	RuneSlot      string `json:"rune_slot"`
	ShieldSlot    string `json:"shield_slot"`
	HelmetSlot    string `json:"helmet_slot"`
	BodyArmorSlot string `json:"body_armor_slot"`
	LegArmorSlot  string `json:"leg_armor_slot"`
	BootsSlot     string `json:"boots_slot"`
	Ring1Slot     string `json:"ring1_slot"`
	Ring2Slot     string `json:"ring2_slot"`
	AmuletSlot    string `json:"amulet_slot"`
	Artifact1Slot string `json:"artifact1_slot"`
	Artifact2Slot string `json:"artifact2_slot"`
	Artifact3Slot string `json:"artifact3_slot"`
	Utility1Slot  string `json:"utility1_slot"`
	Utility1Qty   int    `json:"utility1_slot_quantity"`
	Utility2Slot  string `json:"utility2_slot"`
	Utility2Qty   int    `json:"utility2_slot_quantity"`
	BagSlot       string `json:"bag_slot"`

	// --- Задания ---
	Task         string `json:"task"`
	TaskType     string `json:"task_type"`
	TaskProgress int    `json:"task_progress"`
	TaskTotal    int    `json:"task_total"`

	// --- Инвентарь ---
	InventoryMaxItems int             `json:"inventory_max_items"`
	Inventory         []InventoryItem `json:"inventory"`
}

// GetSkills возвращает карту всех навыков для удобного доступа в логике игры.
func (c *Character) GetSkills() map[string]SkillData {
	return map[string]SkillData{
		"mining":          {Level: c.MiningLevel, XP: c.MiningXP, MaxXP: c.MiningMaxXP},
		"woodcutting":     {Level: c.WoodcuttingLevel, XP: c.WoodcuttingXP, MaxXP: c.WoodcuttingMaxXP},
		"fishing":         {Level: c.FishingLevel, XP: c.FishingXP, MaxXP: c.FishingMaxXP},
		"weaponcrafting":  {Level: c.WeaponcraftingLevel, XP: c.WeaponcraftingXP, MaxXP: c.WeaponcraftingMaxXP},
		"gearcrafting":    {Level: c.GearcraftingLevel, XP: c.GearcraftingXP, MaxXP: c.GearcraftingMaxXP},
		"jewelrycrafting": {Level: c.JewelrycraftingLevel, XP: c.JewelrycraftingXP, MaxXP: c.JewelrycraftingMaxXP},
		"cooking":         {Level: c.CookingLevel, XP: c.CookingXP, MaxXP: c.CookingMaxXP},
		"alchemy":         {Level: c.AlchemyLevel, XP: c.AlchemyXP, MaxXP: c.AlchemyMaxXP},
	}
}
