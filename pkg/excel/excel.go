package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
	"sync"
)

var once = &sync.Once{}

func ReadExcel(filePath string) error {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(fmt.Sprintf("open file %s fail", filePath))
		return err
	}
	rows, err := file.GetRows(file.GetSheetName(0))
	if err != nil {
		return err
	}
	for index, row := range rows {
		if index < 2 {
			continue
		}
		buildInfo(row)
	}
	return nil
}

func buildInfo(row []string) {
	once.Do(func() {
		leftShore = make(map[int64]ExcelInfo)
		rightShore = make(map[int64]ExcelInfo)
	})
	if len(row[4]) <= 0 {
		return
	}
	i, _ := strconv.ParseInt(row[4], 10, 16)
	excelInfo := ExcelInfo{
		ShoreType:  row[3],
		ShoreIndex: i,
		Longitude:  row[6],
		Latitude:   row[5],
	}
	switch excelInfo.ShoreType {
	case "Y":
		rightShore[excelInfo.ShoreIndex] = excelInfo
		break
	case "Z":
		leftShore[excelInfo.ShoreIndex] = excelInfo
	}
}
