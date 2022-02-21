package main

import (
	"Crawler_ESAJ/CSV"
	"Crawler_ESAJ/Error"
	"Crawler_ESAJ/Functions"
	"Crawler_ESAJ/Multithread"
	Structs "Crawler_ESAJ/Struct"
	"fmt"
	"github.com/tebeka/selenium"
	"os"
	"time"
)

const maxMonths = 1

var initDate string
var endDate string
var data []Structs.CrawlerData
var driver selenium.WebDriver

func main() {
	fmt.Print("-> Data Inicial: ")
	_, err0 := fmt.Scanln(&initDate)
	Error.CheckError(err0)
	fmt.Print("-> Data Final: ")
	_, err01 := fmt.Scanln(&endDate)
	Error.CheckError(err01)

	fmt.Println("Start:", time.Now().Format("3:4:5"))

	courtOfJustice := []string{/*"tjsp","tjac","tjal","tjam",*/"tjce", "tjms"}

	for i := 0; i < len(courtOfJustice); i++ {
		var searchLink string

		if courtOfJustice[i] == "tjam"{
			searchLink = "https://consultasaj.tjam.jus.br/cjsg/consultaCompleta.do"
		} else if courtOfJustice[i] == "tjal"{
			searchLink = "https://www2.tjal.jus.br/cjsg/consultaCompleta.do"
		} else{
			searchLink = "https://esaj." + courtOfJustice[i] + ".jus.br/cjsg/consultaCompleta.do"
		}

		months := Functions.MonthAmount(initDate, endDate)
		threads := Functions.Threads(months, maxMonths)
		loop := Functions.LoopAmount(months,threads)

		for j := 0; j < loop; j++ {
			var date string
			if j == 0 {
				date = initDate
			} else {
				date = Functions.AddData(initDate, maxMonths*j)
			}
			Multithread.CreateWorkerPool(driver, threads , date, searchLink, courtOfJustice[i])
		}

		finalExport(courtOfJustice[i], loop)

		fmt.Println(courtOfJustice[i], "OK",  time.Now().Format("3:4:5"))
	}
}

func finalExport(court string, loop int){
	data = nil
	var years []string
	for k := 0; k < loop; k++ {
		var date0 string

		if k== 0 {
			date0 = initDate
		} else {
			date0 = Functions.AddData(initDate, maxMonths*k)
		}

		year, month := Functions.YearMonthString(date0)
		years = append(years, year)

		path := "Files CSV/"+ court + "/" + year +  "/" + "Revocação_ESAJ"  + "_" + month +".csv"

		csvD := CSV.ReadCsvFile(path)
		for i := 0; i < len(csvD); i++ {
			data = append(data, csvD[i])
		}

		os.Remove(path)
		for i := 0; i < len(years); i++ {
			os.Remove("Files CSV/"+ court + "/" + years[i])
		}
		os.Remove("Files CSV/"+ court)


	}


	CSV.ExportFinalCSV("Revocação_ESAJ_" + court +".csv", data)
}