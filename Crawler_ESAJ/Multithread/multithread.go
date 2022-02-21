package Multithread

import (
	"Crawler_ESAJ/Crawler"
	"Crawler_ESAJ/Functions"
	"Crawler_ESAJ/WebDriver"
	"github.com/tebeka/selenium"
	"sync"
)

func CreateWorkerPool(driver selenium.WebDriver, noOfWorkers int, date string, searchLink string, courtOfJustice string) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		date1 := Functions.AddData(date, i)
		dateF := Functions.AddDataMonth(date, i+1)
		wg.Add(1)
		go worker(driver, &wg, date1, dateF , searchLink, courtOfJustice)
	}
	wg.Wait()
}

func worker(driver selenium.WebDriver, wg *sync.WaitGroup, date1 string, date2 string, searchLink string, courtOfJustice string)  {
	driver = WebDriver.SeleniumWebDriver()
	Crawler.DayCrawler(driver, date1, date2, searchLink, courtOfJustice)
	wg.Done()
}
