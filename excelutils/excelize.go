package excelutils

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/crazy-me/houstool/convert"
	"os"
	"strconv"
)

// create Excel.xlsx
// data 数据
// header 表头
// filePath 文件路径
// sheetName 文件名称
func CreateExcel(data [][]interface{}, header []string, filePath, sheetName string) (fileName string, err error) {
	if len(data) > MAX_LINE {
		err = errors.New("cannot be greater than 1000 lines")
		return
	}
	// 文件路径及名称
	fileName = filePath + string(os.PathSeparator) + sheetName + ".xlsx"
	excelFile := excelize.NewFile()
	err = excelFile.SetDefinedName(&excelize.DefinedName{Name: sheetName})
	if err != nil {
		return
	}

	// 设置Sheet
	sheetIndex := excelFile.NewSheet(sheetName)
	excelFile.SetActiveSheet(sheetIndex)

	// 创建样式
	tableHeaderStyle, _ := excelFile.NewStyle(TABLE_HEADER)
	tableDataStyle, _ := excelFile.NewStyle(TABLE_DATA)

	// 写表头
	for k, v := range header {
		ck := convert.Num2alpha(k + 1)
		// 从第二行开始写表头
		_ = excelFile.SetCellValue(sheetName, ck+strconv.Itoa(2), v)
		_ = excelFile.SetCellStyle(sheetName, ck+strconv.Itoa(2), ck+strconv.Itoa(2), tableHeaderStyle)
	}

	// 写数据
	for key, value := range data {
		for kk, vv := range value {
			ck := convert.Num2alpha(kk + 1)
			_ = excelFile.SetCellValue(sheetName, ck+strconv.Itoa(key+3), vv)
			_ = excelFile.SetCellStyle(sheetName, ck+strconv.Itoa(key+3), ck+strconv.Itoa(key+3), tableDataStyle)
		}
	}

	err = excelFile.SaveAs(fileName)
	if err != nil {
		return
	}

	return
}

// reader excel
// fileName 文件
func ReaderExcel(fileName string) (data [][]string, err error) {
	readerExcel, err := excelize.OpenFile(fileName)
	if err != nil {
		return
	}

	rows, err := readerExcel.GetRows("Sheet1")
	if err != nil {
		return
	}

	for _, row := range rows {
		data = append(data, row)
	}

	return
}
