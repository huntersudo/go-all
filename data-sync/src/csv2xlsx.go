package main

import (
	"encoding/csv"
	"os"

	"github.com/astaxie/beego/logs"
	"github.com/tealeg/xlsx"
)

// GenXLSXFromCSV generate excel xlsx format file from csv file
func GenXLSXFromCSV(csvPath, xlsxPath, delimiter string) error {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	if len(delimiter) > 0 {
		reader.Comma = rune(delimiter[0])
	} else {
		reader.Comma = rune(',')
	}
	xlsxFile := xlsx.NewFile()
	sheet, err := xlsxFile.AddSheet("data")
	if err != nil {
		return err
	}
	fields, err := reader.Read()
	for err == nil {
		row := sheet.AddRow()
		for _, field := range fields {
			cell := row.AddCell()
			cell.Value = field
		}
		fields, err = reader.Read()
	}
	logs.Info(err)
	return xlsxFile.Save(xlsxPath)
}
