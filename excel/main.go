package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/quexer/utee"
	"github.com/tealeg/xlsx"
	"io"
	"os"
	"strings"
)

func main() {
	path := "/Users/zhaijian/js/taobao-area-php/tmp/area.csv"
	areaMap, err := ParseArea(path)
	utee.Chk(err)

	//"441900"
	a, err := areaMap.GetCityOrTown("441900")
	utee.Chk(err)

	for _, a1 := range a {
		fmt.Println(a1.Value, a1.Name)
	}
}

func ParseArea(file string) (*AreaMap, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	areaMap := make(map[string][]*Area)

	defer f.Close()

	buf := bufio.NewReader(f)
	_, err = buf.ReadString('\n')
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, errors.WithStack(err)
			}
		}

		arr := strings.Split(line, "\t")
		if len(arr) != 9 {
			return nil, errors.New("文件列长度错误")
		}

		areas, ok := areaMap[arr[4]]
		if !ok {
			areas = make([]*Area, 0)
		}

		area := &Area{
			Value: arr[0],
			Name:  arr[1],
		}

		areas = append(areas, area)
		areaMap[arr[4]] = areas
	}
	return &AreaMap{
		M: areaMap,
	}, nil
}

type AreaMap struct {
	M map[string][]*Area
}

func (am *AreaMap) GetProvince() ([]*Area, error) {
	areas, ok := am.M["1"]
	if !ok {
		return nil, errors.New("获取失败")
	}
	return areas, nil
}

func (am *AreaMap) GetCityOrTown(parentId string) ([]*Area, error) {
	areas, ok := am.M[parentId]
	if !ok {
		return nil, errors.New("获取失败")
	}
	return areas, nil
}

type Area struct {
	Value string
	Name  string
}

func genFile(fileName string) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	writer := NewExcelWriter(buffer)

	writer.Write([]string{"订单编号", "订单项ID", "物流公司", "物流单号"})

	for i := 0; i < 2100; i++ {
		record := []string{}

		record = append(record, "2018030186375203")
		record = append(record, "1365")
		record = append(record, "天天快递")
		record = append(record, "1234567890123456")

		writer.Write(record)
	}

	err := writer.file.Save(fileName)
	utee.Chk(err)
}

type ExcelWriter struct {
	file  *xlsx.File
	sheet *xlsx.Sheet
	w     io.Writer
}

func NewExcelWriter(w io.Writer) *ExcelWriter {
	xw := &ExcelWriter{
		file: xlsx.NewFile(),
		w:    w,
	}
	sheet, err := xw.file.AddSheet("Sheet1")
	utee.Chk(err)
	xw.sheet = sheet

	return xw
}

func (p *ExcelWriter) Write(record []string) {
	row := p.sheet.AddRow()
	for _, s := range record {
		cell := row.AddCell()
		cell.Value = s
	}
}

func (p *ExcelWriter) Flush() error {
	return p.file.Write(p.w)
}
