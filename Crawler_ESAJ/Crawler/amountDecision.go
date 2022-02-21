package Crawler

import (
	"Crawler_ESAJ/Error"
	"github.com/tebeka/selenium"
	"strconv"
	"strings"
)

func amountDecision(driver selenium.WebDriver) (int,int,int) {
	var finalTextA string
	var finalTextH string
	var finalTextD string

	textA := "0"
	textH := "0"
	textD := "0"

	totalLI, err := driver.FindElements(selenium.ByXPATH, "//*[@id=\"tabs\"]/ul/li")
	Error.CheckError(err)

	if len(totalLI) == 3 {
		amountReturnA, err0 := driver.FindElements(selenium.ByXPATH, "//*[@id=\"tabs\"]/ul/li[1]/a")
		Error.CheckError(err0)

		amountReturnH, err1 := driver.FindElements(selenium.ByXPATH, "//*[@id=\"tabs\"]/ul/li[2]/a")
		Error.CheckError(err1)

		amountReturnD, err2 := driver.FindElements(selenium.ByXPATH, "//*[@id=\"tabs\"]/ul/li[3]/a")
		Error.CheckError(err2)

		if len(amountReturnA) > 0 {
			textA, _ = amountReturnA[0].Text()
		}

		if len(amountReturnH) > 0 {
			textH, _ = amountReturnH[0].Text()
		}

		if len(amountReturnD) > 0 {
			textD, _ = amountReturnD[0].Text()
		}

		finalTextA = strings.Replace(strings.Replace(strings.Replace(textA, "Acórdãos", "", -1), "(", "", -1), ")", "", -1)
		finalTextH = strings.Replace(strings.Replace(strings.Replace(textH, "Homologações de Acordo", "", -1), "(", "", -1), ")", "", -1)
		finalTextD = strings.Replace(strings.Replace(strings.Replace(textD, "Decisões Monocráticas", "", -1), "(", "", -1), ")", "", -1)

	}else{
		amountReturnA, err0 := driver.FindElements(selenium.ByXPATH, "//*[@id=\"tabs\"]/ul/li[1]/a")
		Error.CheckError(err0)

		amountReturnD, err2 := driver.FindElements(selenium.ByXPATH, "//*[@id=\"tabs\"]/ul/li[2]/a")
		Error.CheckError(err2)

		if len(amountReturnA) > 0 {
			textA, _ = amountReturnA[0].Text()
		}

		if len(amountReturnD) > 0 {
			textD, _ = amountReturnD[0].Text()
		}

		finalTextA = strings.Replace(strings.Replace(strings.Replace(textA, "Acórdãos", "", -1), "(", "", -1), ")", "", -1)
		finalTextH = strings.Replace(strings.Replace(strings.Replace(textH, "Homologações de Acordo", "", -1), "(", "", -1), ")", "", -1)
		finalTextD = strings.Replace(strings.Replace(strings.Replace(textD, "Decisões Monocráticas", "", -1), "(", "", -1), ")", "", -1)
	}


	numberA, err6 := strconv.Atoi(finalTextA)
	Error.CheckError(err6)

	numberH, err7 := strconv.Atoi(finalTextH)
	Error.CheckError(err7)

	numberD, err8 := strconv.Atoi(finalTextD)
	Error.CheckError(err8)

	return numberA, numberH, numberD
}
