package excel

import "fmt"

var leftShore map[int64]ExcelInfo

var rightShore map[int64]ExcelInfo

type ExcelInfo struct {
	ShoreType  string
	ShoreIndex int64
	Longitude  string
	Latitude   string
}

func (e ExcelInfo) check() bool {
	return len(e.ShoreType) > 0 && len(e.Longitude) > 0 && len(e.Latitude) > 0
}

func (e ExcelInfo) print() {
	fmt.Println(fmt.Sprintf("%s;%s;%s", e.ShoreType, e.Longitude, e.Latitude))
}

func Print() {
	fmt.Println(leftShore)
	fmt.Println(rightShore)
}
