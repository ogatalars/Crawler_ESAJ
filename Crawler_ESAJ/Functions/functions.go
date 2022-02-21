package Functions

import (
	"Crawler_ESAJ/Error"
	"strconv"
	"strings"
	"time"
)

func MonthAmount(initDate string, endDate string) int {
	var totalMonth int
	date1 := strings.Replace(initDate, "/", "-", -1)
	date2 := strings.Replace(endDate, "/", "-", -1)

	date3, err0 := time.Parse("02-01-2006", date1)
	Error.CheckError(err0)
	date4, err1 := time.Parse("02-01-2006", date2)
	Error.CheckError(err1)

	year := (date4.Year() - date3.Year()) * 12
	month := monthNumber(date2) - monthNumber(date1)

	date5 := year + month

	if date5 <= 12{
		totalMonth = 1
	} else {
		totalMonth = date5
	}

	return totalMonth
}

func LoopAmount(months int, magicNumber int) int {
	loop := months / magicNumber
	remainder := months % magicNumber

	if months > 12 {
		if remainder != 0 {
			loop = loop + 1
		}
	}else if months == 1 {
		loop = 1
	}

	return loop
}

func Threads(months int, maxMonths int) int {
	var threadNumber int
	if months < maxMonths {
		threadNumber = months
	} else {
		threadNumber = maxMonths
	}
	return threadNumber
}

func AddData(searchDate string, i int) string {
	date1 := strings.Replace(searchDate, "/", "-", -1)

	date2, err0 := time.Parse("02-01-2006", date1)
	Error.CheckError(err0)

	date3 := date2.AddDate(0, 1*i,0)
	date4 := date3.Format("02-01-2006")

	return date4
}

func AddDataMonth(searchDate string, i int) string {
	date1 := strings.Replace(searchDate, "/", "-", -1)

	date2, err0 := time.Parse("02-01-2006", date1)
	Error.CheckError(err0)

	date3 := date2.AddDate(0, 1*i,0)

	if date3.Day() == 1 {
		date3 = date3.AddDate(0,0,-1)
	}


	date4 := date3.Format("02-01-2006")

	return date4
}

func monthNumber(date string) int{
	monthA := date[len(date)-7:]
	monthB := monthA[0:2]

	monthFinal, err := strconv.Atoi(monthB)
	Error.CheckError(err)

	return monthFinal
}

func YearMonthString(initDate string) (string, string) {
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
