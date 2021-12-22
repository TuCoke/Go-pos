package service

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"mian/model"
	"mian/util"
	"time"
)

func SaveHouse(house *model.House) error {
	house.AddDate    = util.FormatDate(time.Now(),1)
	house.UpdateDate = util.FormatDate(time.Now(),1)
	oldHouse:=&model.House{}
	err:=g.DB().Table("house").Struct(oldHouse,"hash_id=?",house.HashId)
	//添加
	if err==sql.ErrNoRows {
		//_,err=g.DB().Insert("house",house)
		_,err=g.DB().Table("house").Insert(house)
		//更新
	} else if err==nil {
		switch house.CityId {
		//舟山
		case 32:

		}
		_,err=g.DB().Table("house").FieldsEx("id,hash_id,add_date").Update(house,"id=?",oldHouse.Id)
	}
	return err
}

//楼盘列表
func HouseList(where string) ([]*model.HouseViewModel,error) {
	sql:=fmt.Sprintf("SELECT m.id,m.promotion_name,m.info,m.hash_id FROM house m where 1=1 %s order by m.id asc",where)
	var house []*model.HouseViewModel
	err:=g.DB().GetStructs(&house,sql)
	return house,err
}
