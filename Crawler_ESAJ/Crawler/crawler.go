package Crawler

import (
	"Crawler_ESAJ/CSV"
	"Crawler_ESAJ/Error"
	Structs "Crawler_ESAJ/Struct"
	"github.com/tebeka/selenium"
	"strconv"
)

func DayCrawler(driver selenium.WebDriver, initDate string, endDate string, searchLink string, courtOfJustice string) {
	var dataArray []Structs.CrawlerData

	err1 := Search(driver, initDate, endDate, searchLink)
	Error.CheckError(err1)

	totalA, totalH, totalD := amountDecision(driver)
	totalDecision := totalA+totalH+totalD

	year, month := yearMonthString(initDate)

	numbers := Structs.CrawlerData{
		TJ: 				courtOfJustice,
		DT1: 				initDate,
		DT2: 				endDate,
		Ano: 				year,
		Mes: 				month,
		Acordao:            strconv.Itoa(totalA),
		Homologacao_Acordo: strconv.Itoa(totalH),
		Decisao_Mono:       strconv.Itoa(totalD),
		Total:				strconv.Itoa(totalDecision),

	}

	err5 := driver.Close()
	Error.CheckError(err5)

	dataArray = append(dataArray, numbers)

	CSV.ExportCSV("Revocação_ESAJ", courtOfJustice, year, month, dataArray)

}

func yearMonthString(initDate string) (string, string) {
	var year string
	var month string

	year = initDate[len(initDate)-4:]

	monthA := initDate[len(initDate)-7:]
	monthB := monthA[0:2]
	monthC := "-"

	if monthB == "01" {
		monthC = "Janeiro"
	} else if monthB == "02" {
		monthC = "Fevereiro"
	} else if monthB == "03" {
		monthC = "Março"
	} else if monthB == "04" {
		monthC = "Abril"
	} else if monthB == "05" {
		monthC = "Maio"
	} else if monthB == "06" {
		monthC = "Junho"
	} else if monthB == "07" {
		monthC = "Julho"
	} else if monthB == "08" {
		monthC = "Agosto"
	} else if monthB == "09" {
		monthC = "Setembro"
	} else if monthB == "10" {
		monthC = "Outubro"
	} else if monthB == "11" {
		monthC = "Novembro"
	} else if monthB == "12" {
		monthC = "Dezembro"
	}

	month = monthC

	return year, month
}