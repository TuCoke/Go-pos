package page_processor

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/guonaihong/gout"
	"github.com/tidwall/gjson"
	"go_spider/core/common/page"
	"go_spider/core/common/request"
	"go_spider/core/spider"
	"mian/model"
	"mian/pipline"
	"mian/service"
	"mian/util"
	"net/http"
)

var ids string

type BuildingRecord struct {
	City    *CityLiShui
	Presell *model.PresellViewModel
}

//region 抓取数据
type BuildingPageProcesser struct{}

func NewBuildingPageProcesser() *BuildingPageProcesser { return &BuildingPageProcesser{} }
func (this *BuildingPageProcesser) Request(req *request.Request) {
}

func (this *BuildingPageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		return
	}
	fmt.Println("请求的状态code", p.GetStatusCode())
	if p.GetStatusCode() == 403 {
		//region 获取token
		header := gout.H{
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
			"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
		}
		text, statusCode := "", 0
		gout.POST("http://183.246.198.92:8003/epoint-web-lishui/rest/oauth2/token").
			SetHeader(header).SetBody("client_id=28a990ac-49c4-425a-a2c5-a896a82fe9ad&client_secret=8d394188-a02f-4164-8e84-296a6060ece1&grant_type=client_credentials").
			BindBody(&text).Code(&statusCode).Do()
		fmt.Println(text)
		for _, toke := range gjson.Get(text, "custom.access_token").Array() {
			fmt.Println("token", toke)
			p.GetRequest().Postdata = "params=%7B'id'%3A+'"+ids+"'%7D&access_token="+toke.String()+""
			p.GetRequest().Header.Set("access_token", toke.String())
		}
		// endregion
	}

	fmt.Println("重新设置的postdata", p.GetRequest().Postdata)
	//请求失败
	if p.GetStatusCode() == 500 {
		p.SetSkip(true)
		fmt.Println("请求失败")
		return
	}
	if p.GetStatusCode() != 200 {
		p.SetStatus(true, fmt.Sprintf("status code：%d", p.GetStatusCode()))
		return
	}
	fmt.Println("shuju",p.GetBodyStr())
	//if !gjson.Get(p.GetBodyStr(), "success").Bool() {
	//	p.SetStatus(true, "数据获取失败")
	//	return
	//}
	meta := p.GetRequest().GetMeta().(*BuildingRecord)
	city := meta.City
	presell := meta.Presell
	ids := gjson.Get(presell.PresellInfo, "id").String()
	var buildingList []*model.Building
	for _, building := range gjson.Get(fmt.Sprintf("{data:%s}",p.GetBodyStr()), "data").Array() {
		buidlingName := building.Get("").String()
		//id:=building.Get("ysxmid").String()
		zrdid := building.Get("ksts").String()
		fmt.Println("ksts",zrdid)
		info, _ := json.Marshal(g.Map{"id": ids, "zrdid": zrdid})
		buildingHashId := util.Base64Encrypt(city.Host + "_" + zrdid)
		buildingList = append(buildingList, &model.Building{
			PresellName:    presell.PresellName,
			BuildingName:   buidlingName,
			Info:           string(info),
			HashId:         presell.HashId,
			PresellHashId:  presell.PresellHashId,
			BuildingHashId: buildingHashId,
		})
	}
	p.AddField("buildingList", buildingList)
}

//endregion
func (city *CityLiShui) Building() {
	list, err := service.PresellList(fmt.Sprintf("and m.city_id=%d", city.CityId))
	if err != nil {
		g.Log().Panic(err)
	}

	header1 := make(http.Header)
	header1.Set("Content-Type","application/x-www-form-urlencoded; charset=UTF-8")
	spider := spider.NewSpider(NewBuildingPageProcesser(), "").SetThreadnum(1).AddPipeline(pipline.NewBuildingPipeline())
	for _, presell := range list {
		urlPost := "http://183.246.198.92:8003/epoint-web-lishui/rest/propertiesrest/getpermitsale"
		//fmt.Println("presellid", presell.PresellInfo)
		//data := make(url.Values)
		id := gjson.Get(presell.PresellInfo, "id").String()
		ids = id
		tokens := "dc1697295ee9530716c9485a9f720d09"
		req := request.NewRequest(urlPost, "text", "", "POST", "params=%7B'id'%3A+'"+id+"'%7D&access_token="+tokens+"", header1, nil, nil, nil)
		spider.AddRequest(req)
	}
	spider.Run()
	g.Log().Info(fmt.Sprintf("丽水市楼幢抓取完成 共%d", len(list)))

}
