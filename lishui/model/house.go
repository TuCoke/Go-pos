package model


type House struct {
	Id int `json:"id"`
	//城市
	CityId int `json:"city_id"`
	//区域
	AreaId int `json:"area_id"`
	HouseId int `json:"house_id"`
	//备案名称
	Name string `json:"name"`
	//推广名称
	PromotionName string `json:"promotion_name"`
	Status string `json:"status"`
	AvgPrice string `json:"avg_price"`
	Address string `json:"address"`
	Developer string `json:"developer"`
	Phone string `json:"phone"`
	Property string `json:"property"`
	Info string `json:"info"`
	AddDate string `json:"add_date"`
	UpdateDate string `json:"update_date"`
	HashId string `json:"hash_id"`
}

type HouseViewModel struct {
	Id int `json:"id"`
	PromotionName string `json:"promotion_name"`
	Info string `json:"info"`
	HashId string `json:"hash_id"`
}
