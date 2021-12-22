package main

import (
  "fmt"
  "mian/page_processor"
)


func main(){
  //region 获取token
  //header:=gout.H{
  //  "Content-Type":"application/x-www-form-urlencoded; charset=UTF-8",
  //  "User-Agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
  //}
  //text,statusCode:="",0
  //gout.POST("http://183.246.198.92:8003/epoint-web-lishui/rest/oauth2/token").
  //  SetHeader(header).SetBody("client_id=28a990ac-49c4-425a-a2c5-a896a82fe9ad&client_secret=8d394188-a02f-4164-8e84-296a6060ece1&grant_type=client_credentials").
  //  BindBody(&text).Code(&statusCode).Do()
  //fmt.Println(text)
  //for _,toke :=range gjson.Get(text,"custom.access_token").Array(){
  //  fmt.Println("token", toke)
  //}

  //return
  fmt.Println("aa")
  test:= page_processor.NewLiShui()
  test.FetchData()
}
