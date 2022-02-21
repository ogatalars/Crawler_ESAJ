package CSV

import (
	"Crawler_ESAJ/Error"
	Structs "Crawler_ESAJ/Struct"
	"encoding/csv"
	"os"
	"path/filepath"
)

func ExportCSV(nameCSV string, courtOfJustice string, year string, month string, dataArray []Structs.CrawlerData) {
	var empData [][]string

	for i := 0; i < len(dataArray); i++ {
		final := []string{
			dataArray[i].TJ,
			dataArray[i].Ano,
			dataArray[i].Mes,
			dataArray[i].Acordao,
			dataArray[i].Homologacao_Acordo,
			dataArray[i].Decisao_Mono,
			dataArray[i].Total,
			dataArray[i].DT1,
			dataArray[i].DT2,
		}
		empData = append(empData, final)
	}

	csvFile, _ := createFile("Files CSV/"+ courtOfJustice + "/" + year +  "/" + nameCSV  + "_" + month +".csv")
	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	err := csvFile.Close()
	Error.CheckError(err)
}

func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func ReadCsvFile(filePath string) []Structs.CrawlerData{
	var data2 []Structs.CrawlerData

	csvFile, err := os.Open(filePath)
	Error.CheckError(err)

	defer csvFile.Close()

	csvLines, err0 := csv.NewReader(csvFile).ReadAll()
	Error.CheckError(err0)

	for _, line := range csvLines {
		emp := Structs.CrawlerData{
			TJ: line[0],
			Ano: line[1],
			Mes: line[2],
			Acordao: line[3],
			Homologacao_Acordo: line[4],
			Decisao_Mono: line[5],
			Total: line[6],
			DT1: line[7],
			DT2: line[8],
		}
		data2 = append(data2, emp)
	}

	return  data2
}

func ExportFinalCSV(nameCSV string, dataArray []Structs.CrawlerData) {
	var empData [][]string

	head := []string{
		"Tribunal",
		"Ano",
		"Mês",
		"Acórdão",
		"Homologação de Acordo",
		"Decisaão monocrárica",
		"Total",
		"Data Publicação 1",
		"Data Publicação 2",
	}

	empData = append(empData, head)

	for i := 0; i < len(dataArray); i++ {
		final := []string{
			dataArray[i].TJ,
			dataArray[i].Ano,
			dataArray[i].Mes,
			dataArray[i].Acordao,
			dataArray[i].Homologacao_Acordo,
			dataArray[i].Decisao_Mono,
			dataArray[i].Total,
			dataArray[i].DT1,
			dataArray[i].DT2,
		}
		empData = append(empData, final)
	}

	csvFile, _ := createFile("Files CSV/" + nameCSV + ".csv")
	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	err := csvFile.Close()
	Error.CheckError(err)
}



