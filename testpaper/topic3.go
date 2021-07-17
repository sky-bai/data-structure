package main

import (
	"fmt"
	"time"
)

type geter struct {
}

//根据身份证号获取年龄
func (g *geter) GetAge(idCard string) int {
	dateStr := idCard[6:14] + "000001"
	idDate, _ := time.Parse("20060102150405", dateStr)

	idYear := idDate.Year()
	idMonth := idDate.Month()
	idDay := idDate.Day()

	nowYear := time.Now().Year()
	nowMonth := time.Now().Month()
	nowDay := time.Now().Day()

	yearDis := nowYear - idYear
	monthDis := nowMonth - idMonth
	dayDis := nowDay - idDay

	if monthDis < 0 || (monthDis == 0 && dayDis < 0) {
		yearDis--
	}
	return yearDis
}
func main() {
	g1 := geter{}
	age := g1.GetAge("510000199201011111")
	fmt.Println(age)
}
