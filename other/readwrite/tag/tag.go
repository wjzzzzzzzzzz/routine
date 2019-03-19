package main

import (
	"reflect"
	"log"
)

type User struct {
	name string `json:"ggg",mm:"fff"`
	id int  `json:"qqqq"`
}
func main(){
	var u User
	of := reflect.TypeOf(u)
	for i:=0;i<of.NumField() ; i++ {
		log.Println(of.Field(i).Tag.Get("json"))
	}
}