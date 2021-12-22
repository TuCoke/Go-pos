package model


type Presell struct {
	PresellId int `json:"presell_id"`
	HouseId int `json:"house_id"`
	TotalCount int `json:"total_count"`
	DealCount int `json:"deal_count"`
	PresellName string `json:"presell_name"`
	Persellno string `json:"persellno"`
	PresellDate *string `json:"presell_date"`
	Info string `json:"info"`
	Stats string `json:"stats"`
	StatsTime int64 `json:"stats_time"`
	AddDate string `json:"add_date"`
	UpdateDate string `json:"update_date"`
	HashId string `json:"hash_id"`
	PresellHashId string `json:"presell_hash_id"`

	CityId int `json:"-"`
	AreaId int `json:"-"`
}

type PresellViewModel struct {
	PresellId int `json:"presell_id"`
	HashId string `json:"hash_id"`
	PresellHashId string `json:"presell_hash_id"`
	PromotionName string `json:"promotion_name"`
	PresellName string `json:"presell_name"`
	HouseInfo string `json:"house_info"`
	PresellInfo string `json:"presell_info"`
}
