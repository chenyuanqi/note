package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ExportXlsx(name string, fields []string, header []string, data []map[string]string) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 创建一个工作表
	sheet1, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置头部
	header_char := 'A'
	for _, h := range header {
		f.SetCellValue("Sheet1", fmt.Sprintf("%c1", header_char), h)
		header_char++
	}

	// 设置单元格的值
	line := 2
	for _, d := range data {
		if len(d) == 0 {
			continue
		}

		value_char := 'A'
		for _, field := range fields {
			f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", value_char, line), d[field])
			value_char++
		}

		line++
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(sheet1)

	// 根据指定路径保存文件
	if err := f.SaveAs(name + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
