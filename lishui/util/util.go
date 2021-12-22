package util

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gurl"
	"github.com/gogf/gf/frame/g"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

//格式化日期
func FormatDate(time time.Time,dateModel int) string {
	switch dateModel {
	case 0:
		return time.Format("2006-01-02")
	case 1:
		return time.Format("2006-01-02 15:04:05")
	default:
		return time.Format("2006-01-02 15:04:05")
	}
}

//Base64Encrypt、gmd5.MustEncryptString
func Base64Encrypt(str string) string {
	return gbase64.EncodeString(gurl.Encode(str))
}

//Base64Decode
func Base64Decode(str string) string {
	b,_:=gbase64.DecodeString(str)
	s,_:=gurl.Decode(string(b))
	return s
}

//MD5
func Md5(data string) string {
	return gmd5.MustEncryptString(data)
}

func GetQueryString(href string,name string) string {
	url,err:=url.Parse(href)
	if err!=nil {
		return ""
	}
	return regexp.MustCompile("(^|&)"+ name +"=([^&]*)(&|$)").FindStringSubmatch(url.RawQuery)[2]
}

//打开excel
func OpenExcel(filename string,sheet string) ([][]string,error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil,err
	}
	return f.GetRows(sheet),nil
}
//写入excel
//header:=[]string{"姓名","年龄","性格","属性"}
//var data []g.Slice
//data = append(data,g.Slice{"张三",18,"暴躁","AA"})
//err:=util.WriteExcel("Sheet1",nil,header,data,"d:/123.xlsx")
//if err!=nil { panic(err) }
func WriteExcel(name string,width []float64,header []string,data []g.Slice,path string) error {
	axis:=[]string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
	f := excelize.NewFile()
	// 创建一个工作表 Sheet1
	index := f.NewSheet(name)
	if width!=nil {
		for i,w:=range width{
			f.SetColWidth(name,axis[i],axis[i],w)
		}
	}
	for i,h:=range header{
		f.SetCellValue(name,fmt.Sprintf("%s%d",axis[i],1),h)
	}
	for i,m:=range data{
		for j,n:=range m{
			f.SetCellValue(name,fmt.Sprintf("%s%d",axis[j],i+2),n)
		}
	}
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	return f.SaveAs(path)
}

var replacer = strings.NewReplacer(" ", "", "\t", "", "\n", "", "\v", "", "\r", "", "\f", "", string([]byte{194, 160}), "")
// TrimSpace 删除网页中的空字符串
func TrimSpace(str string) string {
	return replacer.Replace(str)
}

//正则
func RegexpIndex(data string,reg string,index int) string {
	return regexp.MustCompile(reg).FindStringSubmatch(data)[index]
}

func RegexpString(data string,reg string) string {
	return regexp.MustCompile(reg).FindString(data)
}

func BaseHeader(h map[string]string) http.Header {
	header := make(http.Header)
	header.Set("Accept-Language","zh-CN,zh;q=0.9")
	header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	if h!=nil {
		for k,v:=range h{
			header.Set(k,v)
		}
	}
	return header
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
//获取一房一价当前成交时间
func GetDealDate() string {
	if time.Now().Hour()<10 {
		return time.Now().Add(-time.Hour*24).Format("2006-01-02")
	}else {
		return time.Now().Format("2006-01-02")
	}
}

//func example()  {
//	arr:=[]string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N"}
//	size:=5
//	for i:=0;i<len(arr);i+=size{
//		var temp []string
//		for j:=0;j<size;j++{
//			if (i+j)>=len(arr) {
//				continue
//			}
//			temp= append(temp, arr[i+j])
//		}
//		fmt.Println(strings.Join(temp,","))
//	}
//}