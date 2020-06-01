package Concurrent

import (
	"errors"
	"fmt"
	TestProject "go_study/datastructure"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type BankMarketing struct {
	age                         []byte
	jobAdmin                    []byte
	jobBlueCollar               []byte
	jobEntrepreneur             []byte
	jobHousemaid                []byte
	jobManagement               []byte
	jobRetired                  []byte
	jobSelfEmployed             []byte
	jobServices                 []byte
	jobStudent                  []byte
	jobTechnician               []byte
	jobUnemployed               []byte
	jobUnknown                  []byte
	maritalDivorced             []byte
	maritalMarried              []byte
	maritalSingle               []byte
	maritalUnknown              []byte
	educationBasic4y            []byte
	educationBasic6y            []byte
	educationBasic9y            []byte
	educationHighSchool         []byte
	educationIlliterate         []byte
	educationProfessionalCourse []byte
	educationUniversityDegree   []byte
	educationUnknown            []byte
	creditNo                    []byte
	creditYes                   []byte
	creditUnknown               []byte
	housingNo                   []byte
	housingYes                  []byte
	housingUnknown              []byte
	loanNo                      []byte
	loanYes                     []byte
	loanUnknown                 []byte
	contactCellular             []byte
	contactTelephone            []byte
	contactJan                  []byte
	contactFeb                  []byte
	contactMar                  []byte
	contactApr                  []byte
	contactMay                  []byte
	contactJun                  []byte
	contactJul                  []byte
	contactAug                  []byte
	contactSep                  []byte
	contactOct                  []byte
	contactNov                  []byte
	contactDec                  []byte
	contactMon                  []byte
	contactTue                  []byte
	contactWed                  []byte
	contactThu                  []byte
	contactFri                  []byte
	duration                    int
	campaign                    []byte
	pdays                       int
	pdaysNever                  []byte
	previous                    []byte
	poutcomeFailure             []byte
	poutcomeNonexistent         []byte
	poutcomeSuccess             []byte
	empVarRate                  float64
	consPriceIdx                float64
	consConfIdx                 float64
	euribor3m                   float64
	nrEmployed                  float64
	target                      string
}

type Distance struct {
	index    int
	distance float64
}

func NewBankMarketing() *BankMarketing {
	return new(BankMarketing)
}

func (b *BankMarketing) setData(data []string) {
	if len(data) != 67 {
		panic("Wrong data length: " + string(len(data)))
	}
	b.age = []byte(data[0])
	b.jobAdmin = []byte(data[1])
	b.jobBlueCollar = []byte(data[2])
	b.jobEntrepreneur = []byte(data[3])
	b.jobHousemaid = []byte(data[4])
	b.jobManagement = []byte(data[5])
	b.jobRetired = []byte(data[6])
	b.jobSelfEmployed = []byte(data[7])
	b.jobServices = []byte(data[8])
	b.jobStudent = []byte(data[9])
	b.jobTechnician = []byte(data[10])
	b.jobUnemployed = []byte(data[11])
	b.jobUnknown = []byte(data[12])
	b.maritalDivorced = []byte(data[13])
	b.maritalMarried = []byte(data[14])
	b.maritalSingle = []byte(data[15])
	b.maritalUnknown = []byte(data[16])
	b.educationBasic4y = []byte(data[17])
	b.educationBasic6y = []byte(data[18])
	b.educationBasic9y = []byte(data[19])
	b.educationHighSchool = []byte(data[20])
	b.educationIlliterate = []byte(data[21])
	b.educationProfessionalCourse = []byte(data[22])
	b.educationUniversityDegree = []byte(data[23])
	b.educationUnknown = []byte(data[24])
	b.creditNo = []byte(data[25])
	b.creditYes = []byte(data[26])
	b.creditUnknown = []byte(data[27])
	b.housingNo = []byte(data[28])
	b.housingYes = []byte(data[29])
	b.housingUnknown = []byte(data[30])
	b.loanNo = []byte(data[31])
	b.loanYes = []byte(data[32])
	b.loanUnknown = []byte(data[33])
	b.contactCellular = []byte(data[34])
	b.contactTelephone = []byte(data[35])
	b.contactJan = []byte(data[36])
	b.contactFeb = []byte(data[37])
	b.contactMar = []byte(data[38])
	b.contactApr = []byte(data[39])
	b.contactMay = []byte(data[40])
	b.contactJun = []byte(data[41])
	b.contactJul = []byte(data[42])
	b.contactAug = []byte(data[43])
	b.contactSep = []byte(data[44])
	b.contactOct = []byte(data[45])
	b.contactNov = []byte(data[46])
	b.contactDec = []byte(data[47])
	b.contactMon = []byte(data[48])
	b.contactTue = []byte(data[49])
	b.contactWed = []byte(data[50])
	b.contactThu = []byte(data[51])
	b.contactFri = []byte(data[52])
	b.duration, _ = strconv.Atoi(data[53])
	b.campaign = []byte(data[54])
	b.pdays, _ = strconv.Atoi(data[55])
	b.pdaysNever = []byte(data[56])
	b.previous = []byte(data[57])
	b.poutcomeFailure = []byte(data[58])
	b.poutcomeNonexistent = []byte(data[59])
	b.poutcomeSuccess = []byte(data[60])
	b.empVarRate, _ = strconv.ParseFloat(data[61], 64)
	b.consPriceIdx, _ = strconv.ParseFloat(data[62], 64)
	b.consConfIdx, _ = strconv.ParseFloat(data[63], 64)
	b.euribor3m, _ = strconv.ParseFloat(data[64], 64)
	b.nrEmployed, _ = strconv.ParseFloat(data[65], 64)
	b.target = data[66]
}

func bytes2Float64(bytes []byte) float64 {
	f, _ := strconv.ParseFloat(string(bytes), 64)
	return f
}

func (b *BankMarketing) getData() []float64 {
	return []float64{
		bytes2Float64(b.age),
		bytes2Float64(b.jobAdmin),
		bytes2Float64(b.jobBlueCollar),
		bytes2Float64(b.jobEntrepreneur),
		bytes2Float64(b.jobHousemaid),
		bytes2Float64(b.jobManagement),
		bytes2Float64(b.jobRetired),
		bytes2Float64(b.jobSelfEmployed),
		bytes2Float64(b.jobServices),
		bytes2Float64(b.jobStudent),
		bytes2Float64(b.jobTechnician),
		bytes2Float64(b.jobUnemployed),
		bytes2Float64(b.jobUnknown),
		bytes2Float64(b.maritalDivorced),
		bytes2Float64(b.maritalMarried),
		bytes2Float64(b.maritalSingle),
		bytes2Float64(b.maritalUnknown),
		bytes2Float64(b.educationBasic4y),
		bytes2Float64(b.educationBasic6y),
		bytes2Float64(b.educationBasic9y),
		bytes2Float64(b.educationHighSchool),
		bytes2Float64(b.educationIlliterate),
		bytes2Float64(b.educationProfessionalCourse),
		bytes2Float64(b.educationUniversityDegree),
		bytes2Float64(b.educationUnknown),
		bytes2Float64(b.creditNo),
		bytes2Float64(b.creditYes),
		bytes2Float64(b.creditUnknown),
		bytes2Float64(b.housingNo),
		bytes2Float64(b.housingYes),
		bytes2Float64(b.housingUnknown),
		bytes2Float64(b.loanNo),
		bytes2Float64(b.loanYes),
		bytes2Float64(b.loanUnknown),
		bytes2Float64(b.contactCellular),
		bytes2Float64(b.contactTelephone),
		bytes2Float64(b.contactJan),
		bytes2Float64(b.contactFeb),
		bytes2Float64(b.contactMar),
		bytes2Float64(b.contactApr),
		bytes2Float64(b.contactMay),
		bytes2Float64(b.contactJun),
		bytes2Float64(b.contactJul),
		bytes2Float64(b.contactAug),
		bytes2Float64(b.contactSep),
		bytes2Float64(b.contactOct),
		bytes2Float64(b.contactNov),
		bytes2Float64(b.contactDec),
		bytes2Float64(b.contactMon),
		bytes2Float64(b.contactTue),
		bytes2Float64(b.contactWed),
		bytes2Float64(b.contactThu),
		bytes2Float64(b.contactFri),
		float64(b.duration),
		bytes2Float64(b.campaign),
		float64(b.pdays),
		bytes2Float64(b.pdaysNever),
		bytes2Float64(b.previous),
		bytes2Float64(b.poutcomeFailure),
		bytes2Float64(b.poutcomeNonexistent),
		bytes2Float64(b.poutcomeSuccess),
		b.empVarRate,
		b.consPriceIdx,
		b.consConfIdx,
		b.euribor3m,
		b.nrEmployed,
	}
}

func Load(filePath string) (*TestProject.ArrayList, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("read file error")
	}
	datas := strings.Split(string(data), "\n")
	list := TestProject.NewArray(len(datas))
	for _, line := range datas {
		if line != "" {
			ds := strings.Split(line, ";")
			bankMarketing := NewBankMarketing()
			bankMarketing.setData(ds)
			list.Add(bankMarketing)
		}
	}
	return list, nil
}

func NewDistance() *Distance {
	return &Distance{
		index:    0,
		distance: 0,
	}
}

func KnnClassifier(dataset *TestProject.ArrayList, bankMarketing *BankMarketing, k int) string {
	len := dataset.Length()
	dList := TestProject.NewArray(len)
	for i := 0; i < len; i++ {
		distance := NewDistance()
		distance.index = i
		distance.distance = calculate(dataset.GetValue(i).(*BankMarketing), bankMarketing)
		dList.Add(distance)
	}
	result := dList.ToArray()
	sort.Slice(result, func(i, j int) bool {
		return result[i].(*Distance).distance < result[j].(*Distance).distance
	})
	m := make(map[string]int)
	for _, v := range result[:k] {
		v1 := dataset.GetValue(v.(*Distance).index).(*BankMarketing).target
		if _, ok := m[v1]; !ok {
			m[v1] = 1
		} else {
			m[v1]++
		}
	}
	var t string
	var mx int
	for tag, max := range m {
		if max > mx {
			t = tag
			mx = max
		}
	}
	return t
}

func calculate(data, test *BankMarketing) float64 {
	dataF64 := data.getData()
	testF64 := test.getData()
	var res float64
	len := len(dataF64)
	for i := 0; i < len; i++ {
		res += math.Pow(testF64[i]-dataF64[i], 2)
	}
	return math.Sqrt(res)
}

func Knn() {
	train, _ := Load("data/bank.data")
	test, _ := Load("data/bank.test")
	var success, distaked int
	len := test.Length()
	for i := 0; i < len; i++ {
		t := test.GetValue(i).(*BankMarketing)
		tag := KnnClassifier(train, t, 10)
		if tag == t.target {
			success++
		} else {
			distaked++
		}
	}
	fmt.Println("success:", success)
	fmt.Println("distaked:", distaked)
}
