package model

type CategoryStat struct {
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	Total        int    `json:"total"`
}

type StatsAll struct {
	TtlBook       int64          `json:"ttlBook"`
	TtlInventory  int64          `json:"ttlInventory"`
	TtlCategory   int64          `json:"ttlCategory"`
	DataBook      []CategoryStat `json:"dataBook"`
	DataInventory []CategoryStat `json:"dataInventory"`
}
