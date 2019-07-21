package tools

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func GenerateUserList(filename string) (error, [][]string) {
	var arr [][]string

	f, err := excelize.OpenFile(filename)

	if err != nil {
		// 不存在excel 文件
		return err, arr
	}

	rows, err := f.GetRows("Sheet1")

	return nil, rows

}
