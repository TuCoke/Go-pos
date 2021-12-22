package page_processor

import "net/http"

type CityLiShui struct {
	CityId int
	Host string
}
//丽水市
func NewLiShui() *CityLiShui {
	return &CityLiShui{CityId:33,Host:"jsjlishui"}
}

func (city *CityLiShui) Header(h map[string]string) http.Header {
	header := make(http.Header)
	header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	header.Set("Accept-Language","zh-CN,zh;q=0.9")
	header.Set("Host","jsj.lishui.gov.cn")
	header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	if h!=nil {
		for k,v:=range h{
			header.Add(k,v)
		}
	}
	return header
}
