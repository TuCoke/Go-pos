package service

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"mian/model"
	"mian/util"
	"time"
)

//保存楼幢
func SaveBuilding(building *model.Building) error {
	building.UpdateDate = util.FormatDate(time.Now(),1)
	oldBuilding:=&model.Building{}
	err:=g.DB().Table("building").Struct(oldBuilding,"building_hash_id=?",building.BuildingHashId)
	//添加
	if err==sql.ErrNoRows {
		_,err=g.DB().Table("building").Insert(building)
		//更新
	} else if err==nil {
		switch building.CityId {
		//宁波市
		case 28:
			//楼盘hashid
			building.HashId       =oldBuilding.HashId
			//预售证hashid
			building.PresellHashId=oldBuilding.PresellHashId
			//预售证名称
			building.PresellName  =oldBuilding.PresellName
			//info
			building.Info         =oldBuilding.Info
		//金华市
		case 29:
			//市区
			if building.AreaId==0 {
				building.PresellName   =oldBuilding.PresellName    //预售证
				building.PresellHashId =oldBuilding.PresellHashId  //预售证hashid
			}
		}
		_,err=g.DB().Table("building").FieldsEx("building_id,building_hash_id").Update(building,"building_id=?",oldBuilding.BuildingId)
	}
	return err
}

func SaveBuildingStats(buildingHashId string,statsTime int64,stats string) error {
	record,err:=g.DB().GetOne("SELECT stats FROM building where building_hash_id=?",buildingHashId)
	if err!=nil {
		return err
	}
	if v,ok:=record["stats"];ok && v.String()!=stats {
		_,err:=g.DB().Table("building").Update(g.Map{"stats":stats,"stats_time":statsTime},"building_hash_id=?",buildingHashId)
		return err
	}
	return nil
}

//楼幢列表
func BuildingList(where string) ([]*model.BuildingViewModel,error) {
	sql:=fmt.Sprintf("SELECT o.building_id,m.hash_id,n.presell_hash_id,o.building_hash_id,m.promotion_name,o.building_name,m.info house_info,o.info building_info from house m,presell n,building o where m.hash_id=n.hash_id and n.presell_hash_id=o.presell_hash_id %s order by o.building_id asc",where)
	var building []*model.BuildingViewModel
	err:=g.DB().GetStructs(&building,sql)
	return building,err
}