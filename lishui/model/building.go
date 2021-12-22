package model

type Building struct {
	BuildingId int `json:"building_id"`
	PresellName string `json:"presell_name"`
	HouseId int `json:"house_id"`
	BuildingName string `json:"building_name"`
	TotalCount int `json:"total_count"`
	DealCount int `json:"deal_count"`
	ActiveCount int `json:"active_count"`
	LimitCount int `json:"limit_count"`
	TargetPid string `json:"target_pid"`
	TargetHid int `json:"target_hid"`
	TargetSectionID int `json:"target_sectionID"`
	Info string `json:"info"`
	UpdateDate string `json:"update_date"`
	HashId string `json:"hash_id"`
	PresellHashId string `json:"presell_hash_id"`
	BuildingHashId string `json:"building_hash_id"`
	//城市、区域
	CityId int `json:"-"`
	AreaId int `json:"-"`
}

type BuildingViewModel struct {
	BuildingId int `json:"building_id"`
	HashId string `json:"hash_id"`
	PresellHashId string `json:"presell_hash_id"`
	BuildingHashId string `json:"building_hash_id"`
	PromotionName string `json:"promotion_name"`
	BuildingName string `json:"building_name"`
	HouseInfo string `json:"house_info"`
	BuildingInfo string `json:"building_info"`
}
