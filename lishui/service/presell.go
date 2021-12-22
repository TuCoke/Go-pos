package service

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"mian/model"
	"mian/util"
	"time"
)


func SavePresell(presell *model.Presell) error {
	presell.AddDate    = util.FormatDate(time.Now(),1)
	presell.UpdateDate = util.FormatDate(time.Now(),1)
	oldPresell:=&model.Presell{}
	err:=g.DB().Table("presell").Struct(oldPresell,"presell_hash_id=?",presell.PresellHashId)
	//添加
	if err==sql.ErrNoRows {
		_,err=g.DB().Table("presell").Insert(presell)
		//更新
	} else if err==nil {
		//不更新统计数据
		fieldsEx:=""
		switch presell.CityId {
		//杭州
		case 4:
			presell.PresellName=oldPresell.PresellName  //预售证名称
			if presell.Stats == oldPresell.Stats {      //总套数、成交套数状态
				fieldsEx+=",stats,stats_time"
			}
		default:
			fieldsEx+=",stats,stats_time"
		}
		//不更新核发时间
		presell.PresellDate=oldPresell.PresellDate
		_,err=g.DB().Table("presell").FieldsEx(fmt.Sprintf("presell_id,presell_hash_id,add_date%s",fieldsEx)).Update(presell,"presell_id=?",oldPresell.PresellId)
	}
	return err
}

func SavePresellStats(presellId int,statsTime int64,stats string) error {
	record,err:=g.DB().GetOne("select stats from presell where presell_id=?",presellId)
	if err!=nil {
		return err
	}
	if v,ok:=record["stats"];ok && v.String()!=stats {
		_,err:=g.DB().Table("presell").Update(g.Map{"stats":stats,"stats_time":statsTime},"presell_id=?",presellId)
		return err
	}
	return nil
}


//预售证列表
func PresellList(where string) ([]*model.PresellViewModel,error) {
	return PresellSortList(where,"n.presell_id ASC")
}
//预售证排序列表
func PresellSortList(where string,order string) ([]*model.PresellViewModel,error) {
	sql:=fmt.Sprintf("SELECT n.presell_id,m.hash_id,n.presell_hash_id,m.promotion_name,n.presell_name,m.info house_info,n.info presell_info FROM house m,presell n where m.hash_id=n.hash_id %s ORDER BY %s",where,order)
	var presell []*model.PresellViewModel
	err:=g.DB().GetStructs(&presell,sql)
	return presell,err
}