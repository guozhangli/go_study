package net

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type WDI struct {
	countryName   string
	countryCode   string
	indicatorName string
	indicatorCode string
	values        [55]float64
}

func (w *WDI) setData(data [59]string) error {
	if len(data) != 59 {
		return errors.New("Data length is not correct: " + string(len(data)))
	}
	values := [55]float64{}
	w.countryName = getString(data[0])
	w.countryCode = getString(data[1])
	w.indicatorName = getString(data[2])
	w.indicatorCode = getString(data[3])
	values[0] = getValue(data[4])
	values[1] = getValue(data[5])
	values[2] = getValue(data[6])
	values[3] = getValue(data[7])
	values[4] = getValue(data[8])
	values[5] = getValue(data[9])
	values[6] = getValue(data[10])
	values[7] = getValue(data[11])
	values[8] = getValue(data[12])
	values[9] = getValue(data[13])
	values[10] = getValue(data[14])
	values[11] = getValue(data[15])
	values[12] = getValue(data[16])
	values[13] = getValue(data[17])
	values[14] = getValue(data[18])
	values[15] = getValue(data[19])
	values[16] = getValue(data[20])
	values[17] = getValue(data[21])
	values[18] = getValue(data[22])
	values[19] = getValue(data[23])
	values[20] = getValue(data[24])
	values[21] = getValue(data[25])
	values[22] = getValue(data[26])
	values[23] = getValue(data[27])
	values[24] = getValue(data[28])
	values[25] = getValue(data[29])
	values[26] = getValue(data[30])
	values[27] = getValue(data[31])
	values[28] = getValue(data[32])
	values[29] = getValue(data[33])
	values[30] = getValue(data[34])
	values[31] = getValue(data[35])
	values[32] = getValue(data[36])
	values[33] = getValue(data[37])
	values[34] = getValue(data[38])
	values[35] = getValue(data[39])
	values[36] = getValue(data[40])
	values[37] = getValue(data[41])
	values[38] = getValue(data[42])
	values[39] = getValue(data[43])
	values[40] = getValue(data[44])
	values[41] = getValue(data[45])
	values[42] = getValue(data[46])
	values[43] = getValue(data[47])
	values[44] = getValue(data[48])
	values[45] = getValue(data[49])
	values[46] = getValue(data[50])
	values[47] = getValue(data[51])
	values[48] = getValue(data[52])
	values[49] = getValue(data[53])
	values[50] = getValue(data[54])
	values[51] = getValue(data[55])
	values[52] = getValue(data[56])
	values[53] = getValue(data[57])
	values[54] = getValue(data[58])
	w.values = values
	return nil
}

func getString(s string) string {
	if strings.HasSuffix(s, `"`) {
		return s[1 : len(s)-1]
	}
	return s
}

func getValue(s string) float64 {
	if len(strings.TrimSpace(s)) == 0 {
		return 0.0
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println("string convert float64 error", err)
	}
	return f
}

func parse(str string) [59]string {
	res := [59]string{}
	var index int
	var buffer []string
	var flag = false
	for _, c := range str {
		c := string(c)
		if c == `"` {
			flag = !flag
		} else if c == `,` && !flag {
			var s string
			for _, v := range buffer {
				s += v
			}
			res[index] = s
			index++
			buffer = nil
		} else {
			buffer = append(buffer, c)
		}
	}
	var s string
	for _, v := range buffer {
		s += v
	}
	res[index] = s
	index++
	return res
}

func (w *WDI) getValue(year int) float64 {
	if year < FIRSTYEAR || year >= FIRSTYEAR+len(w.values) {
		log.Println("No data for that year")
	}

	index := int(year - FIRSTYEAR)
	return w.values[index]
}

type WDIDAO struct{}

var list []*WDI
var dao *WDIDAO

func NewWDIDAO(route string) {
	dao = new(WDIDAO)
	list = load(route)
}

func GetDAO() *WDIDAO {
	if dao == nil {
		NewWDIDAO(FILEPATH)
	}
	return dao
}
func (dao *WDIDAO) getData() []*WDI {
	return list
}

const FILEPATH = "../data/WDI_Data.csv"
const FIRSTYEAR = 1960

func load(route string) []*WDI {
	file, err := os.Open(route)
	if err != nil {
		log.Fatal("read file error")
	}
	defer file.Close()
	var list []*WDI
	rd := bufio.NewReader(file)
	rd.ReadString('\n')
	for {
		line, _, err := rd.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		res := parse(string(line))
		w := new(WDI)
		w.setData(res)
		list = append(list, w)
	}
	return list
}

func (w *WDIDAO) query2(codCountry, codIndicator string) string {
	var wdi *WDI
	for _, v := range w.getData() {
		wdi = v
		if wdi.countryCode == codCountry && wdi.indicatorCode == codIndicator {
			break
		}
	}

	var writer string
	writer = codCountry
	writer += ";"
	writer += codIndicator
	writer += ";"
	years := wdi.values
	for i, v := range years {
		writer += strconv.FormatFloat(v, 'E', -1, 64)
		if i < len(years)-1 {
			writer += ";"
		}
	}
	return writer
}

func (w *WDIDAO) query3(codCountry, codIndicator string, year int) string {
	var wdi *WDI
	for _, v := range w.getData() {
		wdi = v
		if wdi.countryCode == codCountry && wdi.indicatorCode == codIndicator {
			break
		}
	}
	var writer string
	writer = codCountry
	writer += ";"
	writer += codIndicator
	writer += ";"
	years := wdi.getValue(year)
	writer += strconv.FormatFloat(years, 'E', -1, 64)
	return writer
}

func (w *WDIDAO) report(codIndicator string) string {
	var wdi *WDI
	var writer string
	writer += codIndicator
	writer += ";"
	for _, v := range w.getData() {
		wdi = v
		if wdi.indicatorCode == codIndicator {
			years := wdi.values
			var mean float64
			for _, v := range years {
				mean += v
			}
			l := len(years)
			mean /= float64(l)
			writer += wdi.countryCode
			writer += ";"
			writer += "" + strconv.FormatFloat(mean, 'E', -1, 64)
			writer += ";"
		}
	}
	return writer
}
