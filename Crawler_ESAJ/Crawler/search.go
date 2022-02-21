package Crawler

import (
	"Crawler_ESAJ/Error"
	"github.com/tebeka/selenium"
)

func Search(driver selenium.WebDriver, initDate string, endDate string, searchLink string) error {
	err1 := driver.Get(searchLink)
	Error.CheckError(err1)

	date1Field, err2 := driver.FindElement(selenium.ByXPATH, "//*[@id=\"iddados.dtPublicacaoInicio\"]")
	Error.CheckError(err2)

	date2Field, err3 := driver.FindElement(selenium.ByXPATH, "//*[@id=\"iddados.dtPublicacaoFim\"]")
	Error.CheckError(err3)

	search1, err4 := driver.FindElement(selenium.ByID, "pbSubmit")
	Error.CheckError(err4)

	selecOrigin(driver)
	selecPublicationType(driver)

	err5 := date1Field.SendKeys(initDate)
	Error.CheckError(err5)

	err6 := date2Field.SendKeys(endDate)
	Error.CheckError(err6)

	err7 := search1.Click()
	Error.CheckError(err7)

	return nil
}

func selecOrigin(driver selenium.WebDriver) {
	elem, err0 := driver.FindElement(selenium.ByID, "origemRecursal")
	Error.CheckError(err0)
	err1 := elem.Click()
	Error.CheckError(err1)
}

func selecPublicationType(driver selenium.WebDriver) {
	elem, err0 := driver.FindElements(selenium.ByID, "Dcheckbox")
	Error.CheckError(err0)

	if len(elem) > 0 {
		err1 := elem[0].Click()
		Error.CheckError(err1)
	}

	elem1, err2 := driver.FindElements(selenium.ByID, "Hcheckbox")
	Error.CheckError(err2)

	if len(elem1) > 0 {
		err3 := elem1[0].Click()
		Error.CheckError(err3)
	}
}
