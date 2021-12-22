package page_processor

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/gogf/gf/frame/g"
	"go_spider/core/common/com_interfaces"
	"go_spider/core/common/page"
	"go_spider/core/common/page_items"
	"go_spider/core/common/request"
	"go_spider/core/spider"
	"mian/model"
	"mian/service"
	"mian/util"
)

//region 抓取数据
type HousePageProcesser struct {}
func NewHousePageProcesser() *HousePageProcesser {return &HousePageProcesser{}}
func (this *HousePageProcesser) Request(req *request.Request) {}

func (this *HousePageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		return
	}

	if p.GetStatusCode() != 200 {
		p.SetStatus(true, fmt.Sprintf("status code：%d", p.GetStatusCode()))
		return
	}
	fmt.Println("json的格式开头", p.GetBodyStr())
	city:=p.GetRequest().GetMeta().(*CityLiShui)
	var houseList []*model.House
	var presellList []*model.Presell

    for _,house:=range gjson.Get(p.GetBodyStr(),"data.jrZrcjxx").Array(){
		// fmt.Println("house", house.Get("id").String())
		//id
		id:=house.Get("id").String()
		//预售项目名称
		name:=house.Get("ysxmmc").String()
		lpmc:=house.Get("lpmc").String()
		//开发商
		// developer:=house.Get("xmgs").String()
		//地址
		//address:=house.Get("xmzl").String()
		//预售证
		//presellName:=house.Get("xkz").String()
		//公示时间
		//qfrq:=house.Get("qfrq").String()
		kprq:=house.Get("kprq").String()
		//楼盘
		info,_:=json.Marshal(g.Map{"id":id})
		// 生成 url + id
		hashId:= util.Base64Encrypt(city.Host+"_"+id +"")
		houseList = append(houseList, &model.House{
			//城市
			CityId:city.CityId,
			//备案名
			Name:lpmc,
			//推广名
			PromotionName:lpmc,
			//销售状态
			Status:"",
			//地址
			Address:"",
			//开发商
			Developer:"",
			Property:"",
			Info:string(info),
			HashId:hashId,
		})

		//预售证
		presellHashId:=hashId
		var presellDate *string
		if len(kprq)>0 {
			presellDate=&kprq
		}else {
			presellDate=nil
		}
		presellList= append(presellList, &model.Presell{
			PresellName:name,
			PresellDate:presellDate,
			Info:string(info),
			HashId:hashId,
			PresellHashId:presellHashId,
		})
	}
	p.AddField("houseList",houseList)
	p.AddField("presellList",presellList)
}
//endregion

type HousePipeline struct {}
func NewHousePipeline() *HousePipeline {return &HousePipeline{}}

func (this *HousePipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {
	allItems:=items.GetAll()
	houseList:=allItems["houseList"].([]*model.House)
	presellList:=allItems["presellList"].([]*model.Presell)
	succ:=true
	for _,house:=range houseList{
		err:=service.SaveHouse(house)
		if err!=nil {
			g.Log().Errorf("%s\n",err.Error())
			succ=false
		}
	}
	for _,presell:=range presellList{
		err:=service.SavePresell(presell)
		if err!=nil {
			g.Log().Errorf("%s\n",err.Error())
			succ=false
		}
	}
	if succ {
		fmt.Printf("楼盘写入成功 %d条，预售证写入成功 %d条 \n",len(houseList),len(presellList))
	}else {
		fmt.Println("楼盘写入失败")
	}
}
//endregion


func (city *CityLiShui) House() {
	spider:=spider.NewSpider(NewHousePageProcesser(), "").SetThreadnum(2).SetSleepTime("rand",300,600).AddPipeline(NewHousePipeline())
	for i:=1;i<=59;i++{
		url:=fmt.Sprintf("http://183.246.198.122:8099/api/Hourse/QueryTrading?xzqh=1&pageSize=25&pageNumber=%d",i)
		req:=request.NewRequest(url, "text","","GET","",city.Header(nil),nil,nil,city)
		spider.AddRequest(req)
	}
	spider.Run()
	g.Log().Info(fmt.Sprintf("丽水市楼盘抓取完成 共%d",59))
}
