package pipline

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"go_spider/core/common/com_interfaces"
	"go_spider/core/common/mlog"
	"go_spider/core/common/page_items"
	"mian/model"
	"mian/service"
)

type BuildingPipeline struct {}
func NewBuildingPipeline() *BuildingPipeline {return &BuildingPipeline{}}

func (this *BuildingPipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {
	item,exist:=items.GetItem("buildingList")
	if !exist {
		return
	}
	buildingList:=item.([]*model.Building)
	succ:=true
	for _,building:=range buildingList{
		err:=service.SaveBuilding(building)
		if err!=nil {
			g.Log().Errorf("%s",err.Error())
			succ=false
		}
	}
	if succ {
		this.TaskProcess(items,fmt.Sprintf("楼幢写入成功 %d条",len(buildingList)))
	}else {
		mlog.StraceInst().Println("楼幢写入失败")
	}
}

func (this *BuildingPipeline) TaskProcess(items *page_items.PageItems,message string) {
	item,exist:=items.GetItem("buildingTaskId")
	if !exist {
		mlog.StraceInst().Println(message)
		return
	}
	err:=service.Task.UpdateTask(item.(int))
	if err!=nil {
		g.Log().Errorf("更新任务状态失败，%s",err.Error())
	}else {
		mlog.StraceInst().Println(fmt.Sprintf("%s，更新任务状态成功",message))
	}
}
